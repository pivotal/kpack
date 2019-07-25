// Code generated by counterfeiter. DO NOT EDIT.
package sourceresolverfakes

import (
	"sync"

	"github.com/pivotal/build-service-beam/pkg/apis/build/v1alpha1"
	"github.com/pivotal/build-service-beam/pkg/git"
	"github.com/pivotal/build-service-beam/pkg/reconciler/v1alpha1/sourceresolver"
)

type FakeGitResolver struct {
	ResolveStub        func(git.Auth, v1alpha1.Git) (v1alpha1.ResolvedGitSource, error)
	resolveMutex       sync.RWMutex
	resolveArgsForCall []struct {
		arg1 git.Auth
		arg2 v1alpha1.Git
	}
	resolveReturns struct {
		result1 v1alpha1.ResolvedGitSource
		result2 error
	}
	resolveReturnsOnCall map[int]struct {
		result1 v1alpha1.ResolvedGitSource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGitResolver) Resolve(arg1 git.Auth, arg2 v1alpha1.Git) (v1alpha1.ResolvedGitSource, error) {
	fake.resolveMutex.Lock()
	ret, specificReturn := fake.resolveReturnsOnCall[len(fake.resolveArgsForCall)]
	fake.resolveArgsForCall = append(fake.resolveArgsForCall, struct {
		arg1 git.Auth
		arg2 v1alpha1.Git
	}{arg1, arg2})
	fake.recordInvocation("Resolve", []interface{}{arg1, arg2})
	fake.resolveMutex.Unlock()
	if fake.ResolveStub != nil {
		return fake.ResolveStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resolveReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitResolver) ResolveCallCount() int {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return len(fake.resolveArgsForCall)
}

func (fake *FakeGitResolver) ResolveCalls(stub func(git.Auth, v1alpha1.Git) (v1alpha1.ResolvedGitSource, error)) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = stub
}

func (fake *FakeGitResolver) ResolveArgsForCall(i int) (git.Auth, v1alpha1.Git) {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	argsForCall := fake.resolveArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeGitResolver) ResolveReturns(result1 v1alpha1.ResolvedGitSource, result2 error) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	fake.resolveReturns = struct {
		result1 v1alpha1.ResolvedGitSource
		result2 error
	}{result1, result2}
}

func (fake *FakeGitResolver) ResolveReturnsOnCall(i int, result1 v1alpha1.ResolvedGitSource, result2 error) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	if fake.resolveReturnsOnCall == nil {
		fake.resolveReturnsOnCall = make(map[int]struct {
			result1 v1alpha1.ResolvedGitSource
			result2 error
		})
	}
	fake.resolveReturnsOnCall[i] = struct {
		result1 v1alpha1.ResolvedGitSource
		result2 error
	}{result1, result2}
}

func (fake *FakeGitResolver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGitResolver) recordInvocation(key string, args []interface{}) {
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

var _ sourceresolver.GitResolver = new(FakeGitResolver)
