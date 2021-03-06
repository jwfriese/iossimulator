// This file was generated by counterfeiter
package iossimulatorfakes

import (
	"sync"

	"github.com/jwfriese/iossimulator"
)

type FakeEnvironmentParser struct {
	ParseEnvironmentStub        func() *iossimulator.SimulatorEnvironment
	parseEnvironmentMutex       sync.RWMutex
	parseEnvironmentArgsForCall []struct{}
	parseEnvironmentReturns     struct {
		result1 *iossimulator.SimulatorEnvironment
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnvironmentParser) ParseEnvironment() *iossimulator.SimulatorEnvironment {
	fake.parseEnvironmentMutex.Lock()
	fake.parseEnvironmentArgsForCall = append(fake.parseEnvironmentArgsForCall, struct{}{})
	fake.recordInvocation("ParseEnvironment", []interface{}{})
	fake.parseEnvironmentMutex.Unlock()
	if fake.ParseEnvironmentStub != nil {
		return fake.ParseEnvironmentStub()
	} else {
		return fake.parseEnvironmentReturns.result1
	}
}

func (fake *FakeEnvironmentParser) ParseEnvironmentCallCount() int {
	fake.parseEnvironmentMutex.RLock()
	defer fake.parseEnvironmentMutex.RUnlock()
	return len(fake.parseEnvironmentArgsForCall)
}

func (fake *FakeEnvironmentParser) ParseEnvironmentReturns(result1 *iossimulator.SimulatorEnvironment) {
	fake.ParseEnvironmentStub = nil
	fake.parseEnvironmentReturns = struct {
		result1 *iossimulator.SimulatorEnvironment
	}{result1}
}

func (fake *FakeEnvironmentParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.parseEnvironmentMutex.RLock()
	defer fake.parseEnvironmentMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeEnvironmentParser) recordInvocation(key string, args []interface{}) {
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

var _ iossimulator.EnvironmentParser = new(FakeEnvironmentParser)
