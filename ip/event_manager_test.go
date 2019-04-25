package ip

import (
	"context"
	"testing"
	"time"

	"github.com/Scalingo/link/config"
	"github.com/Scalingo/link/healthcheck/healthcheckmock"
	"github.com/Scalingo/link/locker/lockermock"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestManager_SingleEtcdRun(t *testing.T) {
	examples := []struct {
		Name           string
		Locker         func(*lockermock.MockLocker)
		ExpectedEvents []string
	}{
		{
			Name: "When refresh fails",
			Locker: func(mock *lockermock.MockLocker) {
				mock.EXPECT().Refresh(gomock.Any()).Return(errors.New("NOP"))
			},
			ExpectedEvents: []string{FaultEvent},
		}, {
			Name: "When IsMaster fails just one time",
			Locker: func(mock *lockermock.MockLocker) {
				mock.EXPECT().Refresh(gomock.Any()).Return(nil)
				mock.EXPECT().IsMaster(gomock.Any()).Return(false, errors.New("NOP"))
			},
			ExpectedEvents: []string{FaultEvent},
		}, {
			Name: "When we are not master",
			Locker: func(mock *lockermock.MockLocker) {
				mock.EXPECT().Refresh(gomock.Any()).Return(nil)
				mock.EXPECT().IsMaster(gomock.Any()).Return(false, nil)
			},
			ExpectedEvents: []string{DemotedEvent},
		}, {
			Name: "When we are master",
			Locker: func(mock *lockermock.MockLocker) {
				mock.EXPECT().Refresh(gomock.Any()).Return(nil)
				mock.EXPECT().IsMaster(gomock.Any()).Return(true, nil)
			},
			ExpectedEvents: []string{ElectedEvent},
		},
	}

	for _, example := range examples {
		t.Run(example.Name, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			locker := lockermock.NewMockLocker(ctrl)
			example.Locker(locker)

			manager := &manager{
				locker: locker,
			}

			eventChan := make(chan string, 10)
			doneChan := make(chan bool)
			manager.eventChan = eventChan
			go func() {
				manager.singleEtcdRun(ctx)
				doneChan <- true
			}()
			timer := time.NewTimer(500 * time.Millisecond)
			var events []string

			cont := true
			for cont {
				select {
				case <-timer.C:
					t.Fatal("Method did not return")
					break
				case event := <-eventChan:
					events = append(events, event)
				case <-doneChan:
					cont = false
				}
			}

			require.Equal(t, len(example.ExpectedEvents), len(events))

			for i := 0; i < len(example.ExpectedEvents); i++ {
				assert.Equal(t, example.ExpectedEvents[i], events[i])
			}
		})
	}
}

func TestManager_HealthChecker(t *testing.T) {
	examples := []struct {
		Name           string
		Checker        func(*healthcheckmock.MockChecker)
		ExpectedEvents []string
	}{
		{
			Name: "With not enough failing events",
			Checker: func(mock *healthcheckmock.MockChecker) {
				mock.EXPECT().IsHealthy(gomock.Any()).Return(false).MaxTimes(2)
				mock.EXPECT().IsHealthy(gomock.Any()).Return(true).AnyTimes()
			},
			ExpectedEvents: []string{HealthCheckSuccessEvent},
		}, {
			Name: "With enough failing events",
			Checker: func(mock *healthcheckmock.MockChecker) {
				mock.EXPECT().IsHealthy(gomock.Any()).Return(false).MaxTimes(3)
				mock.EXPECT().IsHealthy(gomock.Any()).Return(true).AnyTimes()
			},
			ExpectedEvents: []string{HealthCheckFailEvent, HealthCheckSuccessEvent},
		}, {
			Name: "With a success event and a stop",
			Checker: func(mock *healthcheckmock.MockChecker) {
				mock.EXPECT().IsHealthy(gomock.Any()).Return(true).MaxTimes(2)
			},
			ExpectedEvents: []string{HealthCheckSuccessEvent},
		},
	}

	for _, example := range examples {
		t.Run(example.Name, func(t *testing.T) {
			ctx := context.Background()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			checker := healthcheckmock.NewMockChecker(ctrl)
			example.Checker(checker)

			manager := &manager{
				checker:      checker,
				stateMachine: NewStateMachine(ctx, NewStateMachineOpts{}),
				config: config.Config{
					HealthcheckInterval:     10 * time.Millisecond,
					FailCountBeforeFailover: 3,
				},
			}

			eventChan := make(chan string)
			doneChan := make(chan bool)
			manager.eventChan = eventChan
			go func() {
				manager.healthChecker(ctx)
				doneChan <- true
			}()
			timer := time.NewTimer(500 * time.Millisecond)
			var events []string

			cont := true
			i := 0
			for cont {
				select {
				case <-timer.C:
					t.Fatal("Method did not return")
					break
				case event := <-eventChan:
					events = append(events, event)
					i++
				case <-doneChan:
					cont = false
				}
				if i >= len(example.ExpectedEvents) {
					cont = false
				}

			}
			manager.Stop(ctx, func(context.Context) error { return nil })

			for i := 0; i < len(example.ExpectedEvents); i++ {
				assert.Equal(t, example.ExpectedEvents[i], events[i])
			}
		})
	}
}

func TestWaitTwiceLeaseTimeOrReallocation(t *testing.T) {
	c := config.Config{
		KeepAliveInterval: 5 * time.Second,
	}
	t.Run("when someone else took the lock", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		lockerMock := lockermock.NewMockLocker(ctrl)
		lockerMock.EXPECT().IsMaster(gomock.Any()).Return(false, nil).Times(1)

		m := &manager{
			locker: lockerMock,
			config: c,
			stopper: func(ctx context.Context) error {
				t.Log("Stopper should not be called")
				t.Fail()
				return nil
			},
		}

		start := time.Now()
		m.waitTwiceLeaseTimeOrReallocation(context.Background())
		assert.WithinDuration(t, start, time.Now(), 600*time.Millisecond)
	})

	t.Run("if we're not stopping", func(t *testing.T) {
		m := &manager{
			stopper: nil,
			config:  c,
		}
		start := time.Now()
		m.waitTwiceLeaseTimeOrReallocation(context.Background())
		assert.WithinDuration(t, start, time.Now(), 600*time.Millisecond)
	})
}
