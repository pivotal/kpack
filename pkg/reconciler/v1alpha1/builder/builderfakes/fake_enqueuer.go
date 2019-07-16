// Code generated by counterfeiter. DO NOT EDIT.
package builderfakes

import (
	"sync"

	"github.com/pivotal/build-service-system/pkg/apis/build/v1alpha1"
	"github.com/pivotal/build-service-system/pkg/reconciler/v1alpha1/builder"
)

type FakeEnqueuer struct {
	EnqueueStub        func(*v1alpha1.Builder) error
	enqueueMutex       sync.RWMutex
	enqueueArgsForCall []struct {
		arg1 *v1alpha1.Builder
	}
	enqueueReturns struct {
		result1 error
	}
	enqueueReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnqueuer) Enqueue(arg1 *v1alpha1.Builder) error {
	fake.enqueueMutex.Lock()
	ret, specificReturn := fake.enqueueReturnsOnCall[len(fake.enqueueArgsForCall)]
	fake.enqueueArgsForCall = append(fake.enqueueArgsForCall, struct {
		arg1 *v1alpha1.Builder
	}{arg1})
	fake.recordInvocation("Enqueue", []interface{}{arg1})
	fake.enqueueMutex.Unlock()
	if fake.EnqueueStub != nil {
		return fake.EnqueueStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.enqueueReturns
	return fakeReturns.result1
}

func (fake *FakeEnqueuer) EnqueueCallCount() int {
	fake.enqueueMutex.RLock()
	defer fake.enqueueMutex.RUnlock()
	return len(fake.enqueueArgsForCall)
}

func (fake *FakeEnqueuer) EnqueueCalls(stub func(*v1alpha1.Builder) error) {
	fake.enqueueMutex.Lock()
	defer fake.enqueueMutex.Unlock()
	fake.EnqueueStub = stub
}

func (fake *FakeEnqueuer) EnqueueArgsForCall(i int) *v1alpha1.Builder {
	fake.enqueueMutex.RLock()
	defer fake.enqueueMutex.RUnlock()
	argsForCall := fake.enqueueArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEnqueuer) EnqueueReturns(result1 error) {
	fake.enqueueMutex.Lock()
	defer fake.enqueueMutex.Unlock()
	fake.EnqueueStub = nil
	fake.enqueueReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEnqueuer) EnqueueReturnsOnCall(i int, result1 error) {
	fake.enqueueMutex.Lock()
	defer fake.enqueueMutex.Unlock()
	fake.EnqueueStub = nil
	if fake.enqueueReturnsOnCall == nil {
		fake.enqueueReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.enqueueReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEnqueuer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.enqueueMutex.RLock()
	defer fake.enqueueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEnqueuer) recordInvocation(key string, args []interface{}) {
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

var _ builder.Enqueuer = new(FakeEnqueuer)
