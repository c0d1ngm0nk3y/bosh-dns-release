// Code generated by counterfeiter. DO NOT EDIT.
package addressesfakes

import (
	"bosh-dns/dns/config/addresses"
	"sync"
)

type FakeConfigGlobber struct {
	GlobStub        func(string) ([]string, error)
	globMutex       sync.RWMutex
	globArgsForCall []struct {
		arg1 string
	}
	globReturns struct {
		result1 []string
		result2 error
	}
	globReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfigGlobber) Glob(arg1 string) ([]string, error) {
	fake.globMutex.Lock()
	ret, specificReturn := fake.globReturnsOnCall[len(fake.globArgsForCall)]
	fake.globArgsForCall = append(fake.globArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Glob", []interface{}{arg1})
	fake.globMutex.Unlock()
	if fake.GlobStub != nil {
		return fake.GlobStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.globReturns.result1, fake.globReturns.result2
}

func (fake *FakeConfigGlobber) GlobCallCount() int {
	fake.globMutex.RLock()
	defer fake.globMutex.RUnlock()
	return len(fake.globArgsForCall)
}

func (fake *FakeConfigGlobber) GlobArgsForCall(i int) string {
	fake.globMutex.RLock()
	defer fake.globMutex.RUnlock()
	return fake.globArgsForCall[i].arg1
}

func (fake *FakeConfigGlobber) GlobReturns(result1 []string, result2 error) {
	fake.GlobStub = nil
	fake.globReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigGlobber) GlobReturnsOnCall(i int, result1 []string, result2 error) {
	fake.GlobStub = nil
	if fake.globReturnsOnCall == nil {
		fake.globReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.globReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeConfigGlobber) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.globMutex.RLock()
	defer fake.globMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfigGlobber) recordInvocation(key string, args []interface{}) {
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

var _ addresses.ConfigGlobber = new(FakeConfigGlobber)
