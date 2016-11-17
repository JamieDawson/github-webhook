// This file was generated by counterfeiter
package retryhttpfakes

import (
	"sync"
	"time"

	"github.com/concourse/retryhttp"
)

type FakeSleeper struct {
	SleepStub        func(time.Duration)
	sleepMutex       sync.RWMutex
	sleepArgsForCall []struct {
		arg1 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSleeper) Sleep(arg1 time.Duration) {
	fake.sleepMutex.Lock()
	fake.sleepArgsForCall = append(fake.sleepArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	fake.recordInvocation("Sleep", []interface{}{arg1})
	fake.sleepMutex.Unlock()
	if fake.SleepStub != nil {
		fake.SleepStub(arg1)
	}
}

func (fake *FakeSleeper) SleepCallCount() int {
	fake.sleepMutex.RLock()
	defer fake.sleepMutex.RUnlock()
	return len(fake.sleepArgsForCall)
}

func (fake *FakeSleeper) SleepArgsForCall(i int) time.Duration {
	fake.sleepMutex.RLock()
	defer fake.sleepMutex.RUnlock()
	return fake.sleepArgsForCall[i].arg1
}

func (fake *FakeSleeper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sleepMutex.RLock()
	defer fake.sleepMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSleeper) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ retryhttp.Sleeper = new(FakeSleeper)