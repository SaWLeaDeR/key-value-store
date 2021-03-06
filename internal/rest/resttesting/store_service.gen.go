// Code generated by counterfeiter. DO NOT EDIT.
package resttesting

import (
	"context"
	"sync"

	"github.com/SaWLeaDeR/key-value-store/internal"
	"github.com/SaWLeaDeR/key-value-store/internal/rest"
)

type FakeStoreService struct {
	FlushStub        func(context.Context) error
	flushMutex       sync.RWMutex
	flushArgsForCall []struct {
		arg1 context.Context
	}
	flushReturns struct {
		result1 error
	}
	flushReturnsOnCall map[int]struct {
		result1 error
	}
	GetStoreDataStub        func(context.Context, string) (internal.Data, error)
	getStoreDataMutex       sync.RWMutex
	getStoreDataArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getStoreDataReturns struct {
		result1 internal.Data
		result2 error
	}
	getStoreDataReturnsOnCall map[int]struct {
		result1 internal.Data
		result2 error
	}
	StoreDataStub        func(context.Context, string, string) internal.Data
	storeDataMutex       sync.RWMutex
	storeDataArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	storeDataReturns struct {
		result1 internal.Data
	}
	storeDataReturnsOnCall map[int]struct {
		result1 internal.Data
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStoreService) Flush(arg1 context.Context) error {
	fake.flushMutex.Lock()
	ret, specificReturn := fake.flushReturnsOnCall[len(fake.flushArgsForCall)]
	fake.flushArgsForCall = append(fake.flushArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.FlushStub
	fakeReturns := fake.flushReturns
	fake.recordInvocation("Flush", []interface{}{arg1})
	fake.flushMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStoreService) FlushCallCount() int {
	fake.flushMutex.RLock()
	defer fake.flushMutex.RUnlock()
	return len(fake.flushArgsForCall)
}

func (fake *FakeStoreService) FlushCalls(stub func(context.Context) error) {
	fake.flushMutex.Lock()
	defer fake.flushMutex.Unlock()
	fake.FlushStub = stub
}

func (fake *FakeStoreService) FlushArgsForCall(i int) context.Context {
	fake.flushMutex.RLock()
	defer fake.flushMutex.RUnlock()
	argsForCall := fake.flushArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStoreService) FlushReturns(result1 error) {
	fake.flushMutex.Lock()
	defer fake.flushMutex.Unlock()
	fake.FlushStub = nil
	fake.flushReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreService) FlushReturnsOnCall(i int, result1 error) {
	fake.flushMutex.Lock()
	defer fake.flushMutex.Unlock()
	fake.FlushStub = nil
	if fake.flushReturnsOnCall == nil {
		fake.flushReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.flushReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreService) GetStoreData(arg1 context.Context, arg2 string) (internal.Data, error) {
	fake.getStoreDataMutex.Lock()
	ret, specificReturn := fake.getStoreDataReturnsOnCall[len(fake.getStoreDataArgsForCall)]
	fake.getStoreDataArgsForCall = append(fake.getStoreDataArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetStoreDataStub
	fakeReturns := fake.getStoreDataReturns
	fake.recordInvocation("GetStoreData", []interface{}{arg1, arg2})
	fake.getStoreDataMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStoreService) GetStoreDataCallCount() int {
	fake.getStoreDataMutex.RLock()
	defer fake.getStoreDataMutex.RUnlock()
	return len(fake.getStoreDataArgsForCall)
}

func (fake *FakeStoreService) GetStoreDataCalls(stub func(context.Context, string) (internal.Data, error)) {
	fake.getStoreDataMutex.Lock()
	defer fake.getStoreDataMutex.Unlock()
	fake.GetStoreDataStub = stub
}

func (fake *FakeStoreService) GetStoreDataArgsForCall(i int) (context.Context, string) {
	fake.getStoreDataMutex.RLock()
	defer fake.getStoreDataMutex.RUnlock()
	argsForCall := fake.getStoreDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStoreService) GetStoreDataReturns(result1 internal.Data, result2 error) {
	fake.getStoreDataMutex.Lock()
	defer fake.getStoreDataMutex.Unlock()
	fake.GetStoreDataStub = nil
	fake.getStoreDataReturns = struct {
		result1 internal.Data
		result2 error
	}{result1, result2}
}

func (fake *FakeStoreService) GetStoreDataReturnsOnCall(i int, result1 internal.Data, result2 error) {
	fake.getStoreDataMutex.Lock()
	defer fake.getStoreDataMutex.Unlock()
	fake.GetStoreDataStub = nil
	if fake.getStoreDataReturnsOnCall == nil {
		fake.getStoreDataReturnsOnCall = make(map[int]struct {
			result1 internal.Data
			result2 error
		})
	}
	fake.getStoreDataReturnsOnCall[i] = struct {
		result1 internal.Data
		result2 error
	}{result1, result2}
}

func (fake *FakeStoreService) StoreData(arg1 context.Context, arg2 string, arg3 string) internal.Data {
	fake.storeDataMutex.Lock()
	ret, specificReturn := fake.storeDataReturnsOnCall[len(fake.storeDataArgsForCall)]
	fake.storeDataArgsForCall = append(fake.storeDataArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.StoreDataStub
	fakeReturns := fake.storeDataReturns
	fake.recordInvocation("StoreData", []interface{}{arg1, arg2, arg3})
	fake.storeDataMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStoreService) StoreDataCallCount() int {
	fake.storeDataMutex.RLock()
	defer fake.storeDataMutex.RUnlock()
	return len(fake.storeDataArgsForCall)
}

func (fake *FakeStoreService) StoreDataCalls(stub func(context.Context, string, string) internal.Data) {
	fake.storeDataMutex.Lock()
	defer fake.storeDataMutex.Unlock()
	fake.StoreDataStub = stub
}

func (fake *FakeStoreService) StoreDataArgsForCall(i int) (context.Context, string, string) {
	fake.storeDataMutex.RLock()
	defer fake.storeDataMutex.RUnlock()
	argsForCall := fake.storeDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStoreService) StoreDataReturns(result1 internal.Data) {
	fake.storeDataMutex.Lock()
	defer fake.storeDataMutex.Unlock()
	fake.StoreDataStub = nil
	fake.storeDataReturns = struct {
		result1 internal.Data
	}{result1}
}

func (fake *FakeStoreService) StoreDataReturnsOnCall(i int, result1 internal.Data) {
	fake.storeDataMutex.Lock()
	defer fake.storeDataMutex.Unlock()
	fake.StoreDataStub = nil
	if fake.storeDataReturnsOnCall == nil {
		fake.storeDataReturnsOnCall = make(map[int]struct {
			result1 internal.Data
		})
	}
	fake.storeDataReturnsOnCall[i] = struct {
		result1 internal.Data
	}{result1}
}

func (fake *FakeStoreService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.flushMutex.RLock()
	defer fake.flushMutex.RUnlock()
	fake.getStoreDataMutex.RLock()
	defer fake.getStoreDataMutex.RUnlock()
	fake.storeDataMutex.RLock()
	defer fake.storeDataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStoreService) recordInvocation(key string, args []interface{}) {
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

var _ rest.StoreService = new(FakeStoreService)
