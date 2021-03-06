// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"log"
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker/boshdirector"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
	"github.com/pivotal-cf/on-demand-service-broker/config"
	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
)

type FakeBoshClient struct {
	GetTaskStub        func(taskID int, logger *log.Logger) (boshdirector.BoshTask, error)
	getTaskMutex       sync.RWMutex
	getTaskArgsForCall []struct {
		taskID int
		logger *log.Logger
	}
	getTaskReturns struct {
		result1 boshdirector.BoshTask
		result2 error
	}
	getTaskReturnsOnCall map[int]struct {
		result1 boshdirector.BoshTask
		result2 error
	}
	GetTasksStub        func(deploymentName string, logger *log.Logger) (boshdirector.BoshTasks, error)
	getTasksMutex       sync.RWMutex
	getTasksArgsForCall []struct {
		deploymentName string
		logger         *log.Logger
	}
	getTasksReturns struct {
		result1 boshdirector.BoshTasks
		result2 error
	}
	getTasksReturnsOnCall map[int]struct {
		result1 boshdirector.BoshTasks
		result2 error
	}
	GetNormalisedTasksByContextStub        func(deploymentName, contextID string, logger *log.Logger) (boshdirector.BoshTasks, error)
	getNormalisedTasksByContextMutex       sync.RWMutex
	getNormalisedTasksByContextArgsForCall []struct {
		deploymentName string
		contextID      string
		logger         *log.Logger
	}
	getNormalisedTasksByContextReturns struct {
		result1 boshdirector.BoshTasks
		result2 error
	}
	getNormalisedTasksByContextReturnsOnCall map[int]struct {
		result1 boshdirector.BoshTasks
		result2 error
	}
	VMsStub        func(deploymentName string, logger *log.Logger) (bosh.BoshVMs, error)
	vMsMutex       sync.RWMutex
	vMsArgsForCall []struct {
		deploymentName string
		logger         *log.Logger
	}
	vMsReturns struct {
		result1 bosh.BoshVMs
		result2 error
	}
	vMsReturnsOnCall map[int]struct {
		result1 bosh.BoshVMs
		result2 error
	}
	GetDeploymentStub        func(name string, logger *log.Logger) ([]byte, bool, error)
	getDeploymentMutex       sync.RWMutex
	getDeploymentArgsForCall []struct {
		name   string
		logger *log.Logger
	}
	getDeploymentReturns struct {
		result1 []byte
		result2 bool
		result3 error
	}
	getDeploymentReturnsOnCall map[int]struct {
		result1 []byte
		result2 bool
		result3 error
	}
	GetDeploymentsStub        func(logger *log.Logger) ([]boshdirector.Deployment, error)
	getDeploymentsMutex       sync.RWMutex
	getDeploymentsArgsForCall []struct {
		logger *log.Logger
	}
	getDeploymentsReturns struct {
		result1 []boshdirector.Deployment
		result2 error
	}
	getDeploymentsReturnsOnCall map[int]struct {
		result1 []boshdirector.Deployment
		result2 error
	}
	DeleteDeploymentStub        func(name, contextID string, logger *log.Logger, taskReporter *boshdirector.AsyncTaskReporter) (int, error)
	deleteDeploymentMutex       sync.RWMutex
	deleteDeploymentArgsForCall []struct {
		name         string
		contextID    string
		logger       *log.Logger
		taskReporter *boshdirector.AsyncTaskReporter
	}
	deleteDeploymentReturns struct {
		result1 int
		result2 error
	}
	deleteDeploymentReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	GetInfoStub        func(logger *log.Logger) (boshdirector.Info, error)
	getInfoMutex       sync.RWMutex
	getInfoArgsForCall []struct {
		logger *log.Logger
	}
	getInfoReturns struct {
		result1 boshdirector.Info
		result2 error
	}
	getInfoReturnsOnCall map[int]struct {
		result1 boshdirector.Info
		result2 error
	}
	RunErrandStub        func(deploymentName, errandName string, errandInstances []string, contextID string, logger *log.Logger, taskReporter *boshdirector.AsyncTaskReporter) (int, error)
	runErrandMutex       sync.RWMutex
	runErrandArgsForCall []struct {
		deploymentName  string
		errandName      string
		errandInstances []string
		contextID       string
		logger          *log.Logger
		taskReporter    *boshdirector.AsyncTaskReporter
	}
	runErrandReturns struct {
		result1 int
		result2 error
	}
	runErrandReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	VariablesStub        func(deploymentName string, logger *log.Logger) ([]boshdirector.Variable, error)
	variablesMutex       sync.RWMutex
	variablesArgsForCall []struct {
		deploymentName string
		logger         *log.Logger
	}
	variablesReturns struct {
		result1 []boshdirector.Variable
		result2 error
	}
	variablesReturnsOnCall map[int]struct {
		result1 []boshdirector.Variable
		result2 error
	}
	VerifyAuthStub        func(logger *log.Logger) error
	verifyAuthMutex       sync.RWMutex
	verifyAuthArgsForCall []struct {
		logger *log.Logger
	}
	verifyAuthReturns struct {
		result1 error
	}
	verifyAuthReturnsOnCall map[int]struct {
		result1 error
	}
	GetDNSAddressesStub        func(deploymentName string, requestedDNS []config.BindingDNS) (map[string]string, error)
	getDNSAddressesMutex       sync.RWMutex
	getDNSAddressesArgsForCall []struct {
		deploymentName string
		requestedDNS   []config.BindingDNS
	}
	getDNSAddressesReturns struct {
		result1 map[string]string
		result2 error
	}
	getDNSAddressesReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	DeployStub        func(manifest []byte, contextID string, logger *log.Logger, reporter *boshdirector.AsyncTaskReporter) (int, error)
	deployMutex       sync.RWMutex
	deployArgsForCall []struct {
		manifest  []byte
		contextID string
		logger    *log.Logger
		reporter  *boshdirector.AsyncTaskReporter
	}
	deployReturns struct {
		result1 int
		result2 error
	}
	deployReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBoshClient) GetTask(taskID int, logger *log.Logger) (boshdirector.BoshTask, error) {
	fake.getTaskMutex.Lock()
	ret, specificReturn := fake.getTaskReturnsOnCall[len(fake.getTaskArgsForCall)]
	fake.getTaskArgsForCall = append(fake.getTaskArgsForCall, struct {
		taskID int
		logger *log.Logger
	}{taskID, logger})
	fake.recordInvocation("GetTask", []interface{}{taskID, logger})
	fake.getTaskMutex.Unlock()
	if fake.GetTaskStub != nil {
		return fake.GetTaskStub(taskID, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getTaskReturns.result1, fake.getTaskReturns.result2
}

func (fake *FakeBoshClient) GetTaskCallCount() int {
	fake.getTaskMutex.RLock()
	defer fake.getTaskMutex.RUnlock()
	return len(fake.getTaskArgsForCall)
}

func (fake *FakeBoshClient) GetTaskArgsForCall(i int) (int, *log.Logger) {
	fake.getTaskMutex.RLock()
	defer fake.getTaskMutex.RUnlock()
	return fake.getTaskArgsForCall[i].taskID, fake.getTaskArgsForCall[i].logger
}

func (fake *FakeBoshClient) GetTaskReturns(result1 boshdirector.BoshTask, result2 error) {
	fake.GetTaskStub = nil
	fake.getTaskReturns = struct {
		result1 boshdirector.BoshTask
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetTaskReturnsOnCall(i int, result1 boshdirector.BoshTask, result2 error) {
	fake.GetTaskStub = nil
	if fake.getTaskReturnsOnCall == nil {
		fake.getTaskReturnsOnCall = make(map[int]struct {
			result1 boshdirector.BoshTask
			result2 error
		})
	}
	fake.getTaskReturnsOnCall[i] = struct {
		result1 boshdirector.BoshTask
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetTasks(deploymentName string, logger *log.Logger) (boshdirector.BoshTasks, error) {
	fake.getTasksMutex.Lock()
	ret, specificReturn := fake.getTasksReturnsOnCall[len(fake.getTasksArgsForCall)]
	fake.getTasksArgsForCall = append(fake.getTasksArgsForCall, struct {
		deploymentName string
		logger         *log.Logger
	}{deploymentName, logger})
	fake.recordInvocation("GetTasks", []interface{}{deploymentName, logger})
	fake.getTasksMutex.Unlock()
	if fake.GetTasksStub != nil {
		return fake.GetTasksStub(deploymentName, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getTasksReturns.result1, fake.getTasksReturns.result2
}

func (fake *FakeBoshClient) GetTasksCallCount() int {
	fake.getTasksMutex.RLock()
	defer fake.getTasksMutex.RUnlock()
	return len(fake.getTasksArgsForCall)
}

func (fake *FakeBoshClient) GetTasksArgsForCall(i int) (string, *log.Logger) {
	fake.getTasksMutex.RLock()
	defer fake.getTasksMutex.RUnlock()
	return fake.getTasksArgsForCall[i].deploymentName, fake.getTasksArgsForCall[i].logger
}

func (fake *FakeBoshClient) GetTasksReturns(result1 boshdirector.BoshTasks, result2 error) {
	fake.GetTasksStub = nil
	fake.getTasksReturns = struct {
		result1 boshdirector.BoshTasks
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetTasksReturnsOnCall(i int, result1 boshdirector.BoshTasks, result2 error) {
	fake.GetTasksStub = nil
	if fake.getTasksReturnsOnCall == nil {
		fake.getTasksReturnsOnCall = make(map[int]struct {
			result1 boshdirector.BoshTasks
			result2 error
		})
	}
	fake.getTasksReturnsOnCall[i] = struct {
		result1 boshdirector.BoshTasks
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetNormalisedTasksByContext(deploymentName string, contextID string, logger *log.Logger) (boshdirector.BoshTasks, error) {
	fake.getNormalisedTasksByContextMutex.Lock()
	ret, specificReturn := fake.getNormalisedTasksByContextReturnsOnCall[len(fake.getNormalisedTasksByContextArgsForCall)]
	fake.getNormalisedTasksByContextArgsForCall = append(fake.getNormalisedTasksByContextArgsForCall, struct {
		deploymentName string
		contextID      string
		logger         *log.Logger
	}{deploymentName, contextID, logger})
	fake.recordInvocation("GetNormalisedTasksByContext", []interface{}{deploymentName, contextID, logger})
	fake.getNormalisedTasksByContextMutex.Unlock()
	if fake.GetNormalisedTasksByContextStub != nil {
		return fake.GetNormalisedTasksByContextStub(deploymentName, contextID, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getNormalisedTasksByContextReturns.result1, fake.getNormalisedTasksByContextReturns.result2
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextCallCount() int {
	fake.getNormalisedTasksByContextMutex.RLock()
	defer fake.getNormalisedTasksByContextMutex.RUnlock()
	return len(fake.getNormalisedTasksByContextArgsForCall)
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextArgsForCall(i int) (string, string, *log.Logger) {
	fake.getNormalisedTasksByContextMutex.RLock()
	defer fake.getNormalisedTasksByContextMutex.RUnlock()
	return fake.getNormalisedTasksByContextArgsForCall[i].deploymentName, fake.getNormalisedTasksByContextArgsForCall[i].contextID, fake.getNormalisedTasksByContextArgsForCall[i].logger
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextReturns(result1 boshdirector.BoshTasks, result2 error) {
	fake.GetNormalisedTasksByContextStub = nil
	fake.getNormalisedTasksByContextReturns = struct {
		result1 boshdirector.BoshTasks
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetNormalisedTasksByContextReturnsOnCall(i int, result1 boshdirector.BoshTasks, result2 error) {
	fake.GetNormalisedTasksByContextStub = nil
	if fake.getNormalisedTasksByContextReturnsOnCall == nil {
		fake.getNormalisedTasksByContextReturnsOnCall = make(map[int]struct {
			result1 boshdirector.BoshTasks
			result2 error
		})
	}
	fake.getNormalisedTasksByContextReturnsOnCall[i] = struct {
		result1 boshdirector.BoshTasks
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) VMs(deploymentName string, logger *log.Logger) (bosh.BoshVMs, error) {
	fake.vMsMutex.Lock()
	ret, specificReturn := fake.vMsReturnsOnCall[len(fake.vMsArgsForCall)]
	fake.vMsArgsForCall = append(fake.vMsArgsForCall, struct {
		deploymentName string
		logger         *log.Logger
	}{deploymentName, logger})
	fake.recordInvocation("VMs", []interface{}{deploymentName, logger})
	fake.vMsMutex.Unlock()
	if fake.VMsStub != nil {
		return fake.VMsStub(deploymentName, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.vMsReturns.result1, fake.vMsReturns.result2
}

func (fake *FakeBoshClient) VMsCallCount() int {
	fake.vMsMutex.RLock()
	defer fake.vMsMutex.RUnlock()
	return len(fake.vMsArgsForCall)
}

func (fake *FakeBoshClient) VMsArgsForCall(i int) (string, *log.Logger) {
	fake.vMsMutex.RLock()
	defer fake.vMsMutex.RUnlock()
	return fake.vMsArgsForCall[i].deploymentName, fake.vMsArgsForCall[i].logger
}

func (fake *FakeBoshClient) VMsReturns(result1 bosh.BoshVMs, result2 error) {
	fake.VMsStub = nil
	fake.vMsReturns = struct {
		result1 bosh.BoshVMs
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) VMsReturnsOnCall(i int, result1 bosh.BoshVMs, result2 error) {
	fake.VMsStub = nil
	if fake.vMsReturnsOnCall == nil {
		fake.vMsReturnsOnCall = make(map[int]struct {
			result1 bosh.BoshVMs
			result2 error
		})
	}
	fake.vMsReturnsOnCall[i] = struct {
		result1 bosh.BoshVMs
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetDeployment(name string, logger *log.Logger) ([]byte, bool, error) {
	fake.getDeploymentMutex.Lock()
	ret, specificReturn := fake.getDeploymentReturnsOnCall[len(fake.getDeploymentArgsForCall)]
	fake.getDeploymentArgsForCall = append(fake.getDeploymentArgsForCall, struct {
		name   string
		logger *log.Logger
	}{name, logger})
	fake.recordInvocation("GetDeployment", []interface{}{name, logger})
	fake.getDeploymentMutex.Unlock()
	if fake.GetDeploymentStub != nil {
		return fake.GetDeploymentStub(name, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getDeploymentReturns.result1, fake.getDeploymentReturns.result2, fake.getDeploymentReturns.result3
}

func (fake *FakeBoshClient) GetDeploymentCallCount() int {
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	return len(fake.getDeploymentArgsForCall)
}

func (fake *FakeBoshClient) GetDeploymentArgsForCall(i int) (string, *log.Logger) {
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	return fake.getDeploymentArgsForCall[i].name, fake.getDeploymentArgsForCall[i].logger
}

func (fake *FakeBoshClient) GetDeploymentReturns(result1 []byte, result2 bool, result3 error) {
	fake.GetDeploymentStub = nil
	fake.getDeploymentReturns = struct {
		result1 []byte
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBoshClient) GetDeploymentReturnsOnCall(i int, result1 []byte, result2 bool, result3 error) {
	fake.GetDeploymentStub = nil
	if fake.getDeploymentReturnsOnCall == nil {
		fake.getDeploymentReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 bool
			result3 error
		})
	}
	fake.getDeploymentReturnsOnCall[i] = struct {
		result1 []byte
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBoshClient) GetDeployments(logger *log.Logger) ([]boshdirector.Deployment, error) {
	fake.getDeploymentsMutex.Lock()
	ret, specificReturn := fake.getDeploymentsReturnsOnCall[len(fake.getDeploymentsArgsForCall)]
	fake.getDeploymentsArgsForCall = append(fake.getDeploymentsArgsForCall, struct {
		logger *log.Logger
	}{logger})
	fake.recordInvocation("GetDeployments", []interface{}{logger})
	fake.getDeploymentsMutex.Unlock()
	if fake.GetDeploymentsStub != nil {
		return fake.GetDeploymentsStub(logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDeploymentsReturns.result1, fake.getDeploymentsReturns.result2
}

func (fake *FakeBoshClient) GetDeploymentsCallCount() int {
	fake.getDeploymentsMutex.RLock()
	defer fake.getDeploymentsMutex.RUnlock()
	return len(fake.getDeploymentsArgsForCall)
}

func (fake *FakeBoshClient) GetDeploymentsArgsForCall(i int) *log.Logger {
	fake.getDeploymentsMutex.RLock()
	defer fake.getDeploymentsMutex.RUnlock()
	return fake.getDeploymentsArgsForCall[i].logger
}

func (fake *FakeBoshClient) GetDeploymentsReturns(result1 []boshdirector.Deployment, result2 error) {
	fake.GetDeploymentsStub = nil
	fake.getDeploymentsReturns = struct {
		result1 []boshdirector.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetDeploymentsReturnsOnCall(i int, result1 []boshdirector.Deployment, result2 error) {
	fake.GetDeploymentsStub = nil
	if fake.getDeploymentsReturnsOnCall == nil {
		fake.getDeploymentsReturnsOnCall = make(map[int]struct {
			result1 []boshdirector.Deployment
			result2 error
		})
	}
	fake.getDeploymentsReturnsOnCall[i] = struct {
		result1 []boshdirector.Deployment
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) DeleteDeployment(name string, contextID string, logger *log.Logger, taskReporter *boshdirector.AsyncTaskReporter) (int, error) {
	fake.deleteDeploymentMutex.Lock()
	ret, specificReturn := fake.deleteDeploymentReturnsOnCall[len(fake.deleteDeploymentArgsForCall)]
	fake.deleteDeploymentArgsForCall = append(fake.deleteDeploymentArgsForCall, struct {
		name         string
		contextID    string
		logger       *log.Logger
		taskReporter *boshdirector.AsyncTaskReporter
	}{name, contextID, logger, taskReporter})
	fake.recordInvocation("DeleteDeployment", []interface{}{name, contextID, logger, taskReporter})
	fake.deleteDeploymentMutex.Unlock()
	if fake.DeleteDeploymentStub != nil {
		return fake.DeleteDeploymentStub(name, contextID, logger, taskReporter)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteDeploymentReturns.result1, fake.deleteDeploymentReturns.result2
}

func (fake *FakeBoshClient) DeleteDeploymentCallCount() int {
	fake.deleteDeploymentMutex.RLock()
	defer fake.deleteDeploymentMutex.RUnlock()
	return len(fake.deleteDeploymentArgsForCall)
}

func (fake *FakeBoshClient) DeleteDeploymentArgsForCall(i int) (string, string, *log.Logger, *boshdirector.AsyncTaskReporter) {
	fake.deleteDeploymentMutex.RLock()
	defer fake.deleteDeploymentMutex.RUnlock()
	return fake.deleteDeploymentArgsForCall[i].name, fake.deleteDeploymentArgsForCall[i].contextID, fake.deleteDeploymentArgsForCall[i].logger, fake.deleteDeploymentArgsForCall[i].taskReporter
}

func (fake *FakeBoshClient) DeleteDeploymentReturns(result1 int, result2 error) {
	fake.DeleteDeploymentStub = nil
	fake.deleteDeploymentReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) DeleteDeploymentReturnsOnCall(i int, result1 int, result2 error) {
	fake.DeleteDeploymentStub = nil
	if fake.deleteDeploymentReturnsOnCall == nil {
		fake.deleteDeploymentReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.deleteDeploymentReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetInfo(logger *log.Logger) (boshdirector.Info, error) {
	fake.getInfoMutex.Lock()
	ret, specificReturn := fake.getInfoReturnsOnCall[len(fake.getInfoArgsForCall)]
	fake.getInfoArgsForCall = append(fake.getInfoArgsForCall, struct {
		logger *log.Logger
	}{logger})
	fake.recordInvocation("GetInfo", []interface{}{logger})
	fake.getInfoMutex.Unlock()
	if fake.GetInfoStub != nil {
		return fake.GetInfoStub(logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getInfoReturns.result1, fake.getInfoReturns.result2
}

func (fake *FakeBoshClient) GetInfoCallCount() int {
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	return len(fake.getInfoArgsForCall)
}

func (fake *FakeBoshClient) GetInfoArgsForCall(i int) *log.Logger {
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	return fake.getInfoArgsForCall[i].logger
}

func (fake *FakeBoshClient) GetInfoReturns(result1 boshdirector.Info, result2 error) {
	fake.GetInfoStub = nil
	fake.getInfoReturns = struct {
		result1 boshdirector.Info
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetInfoReturnsOnCall(i int, result1 boshdirector.Info, result2 error) {
	fake.GetInfoStub = nil
	if fake.getInfoReturnsOnCall == nil {
		fake.getInfoReturnsOnCall = make(map[int]struct {
			result1 boshdirector.Info
			result2 error
		})
	}
	fake.getInfoReturnsOnCall[i] = struct {
		result1 boshdirector.Info
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) RunErrand(deploymentName string, errandName string, errandInstances []string, contextID string, logger *log.Logger, taskReporter *boshdirector.AsyncTaskReporter) (int, error) {
	var errandInstancesCopy []string
	if errandInstances != nil {
		errandInstancesCopy = make([]string, len(errandInstances))
		copy(errandInstancesCopy, errandInstances)
	}
	fake.runErrandMutex.Lock()
	ret, specificReturn := fake.runErrandReturnsOnCall[len(fake.runErrandArgsForCall)]
	fake.runErrandArgsForCall = append(fake.runErrandArgsForCall, struct {
		deploymentName  string
		errandName      string
		errandInstances []string
		contextID       string
		logger          *log.Logger
		taskReporter    *boshdirector.AsyncTaskReporter
	}{deploymentName, errandName, errandInstancesCopy, contextID, logger, taskReporter})
	fake.recordInvocation("RunErrand", []interface{}{deploymentName, errandName, errandInstancesCopy, contextID, logger, taskReporter})
	fake.runErrandMutex.Unlock()
	if fake.RunErrandStub != nil {
		return fake.RunErrandStub(deploymentName, errandName, errandInstances, contextID, logger, taskReporter)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.runErrandReturns.result1, fake.runErrandReturns.result2
}

func (fake *FakeBoshClient) RunErrandCallCount() int {
	fake.runErrandMutex.RLock()
	defer fake.runErrandMutex.RUnlock()
	return len(fake.runErrandArgsForCall)
}

func (fake *FakeBoshClient) RunErrandArgsForCall(i int) (string, string, []string, string, *log.Logger, *boshdirector.AsyncTaskReporter) {
	fake.runErrandMutex.RLock()
	defer fake.runErrandMutex.RUnlock()
	return fake.runErrandArgsForCall[i].deploymentName, fake.runErrandArgsForCall[i].errandName, fake.runErrandArgsForCall[i].errandInstances, fake.runErrandArgsForCall[i].contextID, fake.runErrandArgsForCall[i].logger, fake.runErrandArgsForCall[i].taskReporter
}

func (fake *FakeBoshClient) RunErrandReturns(result1 int, result2 error) {
	fake.RunErrandStub = nil
	fake.runErrandReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) RunErrandReturnsOnCall(i int, result1 int, result2 error) {
	fake.RunErrandStub = nil
	if fake.runErrandReturnsOnCall == nil {
		fake.runErrandReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.runErrandReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) Variables(deploymentName string, logger *log.Logger) ([]boshdirector.Variable, error) {
	fake.variablesMutex.Lock()
	ret, specificReturn := fake.variablesReturnsOnCall[len(fake.variablesArgsForCall)]
	fake.variablesArgsForCall = append(fake.variablesArgsForCall, struct {
		deploymentName string
		logger         *log.Logger
	}{deploymentName, logger})
	fake.recordInvocation("Variables", []interface{}{deploymentName, logger})
	fake.variablesMutex.Unlock()
	if fake.VariablesStub != nil {
		return fake.VariablesStub(deploymentName, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.variablesReturns.result1, fake.variablesReturns.result2
}

func (fake *FakeBoshClient) VariablesCallCount() int {
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	return len(fake.variablesArgsForCall)
}

func (fake *FakeBoshClient) VariablesArgsForCall(i int) (string, *log.Logger) {
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	return fake.variablesArgsForCall[i].deploymentName, fake.variablesArgsForCall[i].logger
}

func (fake *FakeBoshClient) VariablesReturns(result1 []boshdirector.Variable, result2 error) {
	fake.VariablesStub = nil
	fake.variablesReturns = struct {
		result1 []boshdirector.Variable
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) VariablesReturnsOnCall(i int, result1 []boshdirector.Variable, result2 error) {
	fake.VariablesStub = nil
	if fake.variablesReturnsOnCall == nil {
		fake.variablesReturnsOnCall = make(map[int]struct {
			result1 []boshdirector.Variable
			result2 error
		})
	}
	fake.variablesReturnsOnCall[i] = struct {
		result1 []boshdirector.Variable
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) VerifyAuth(logger *log.Logger) error {
	fake.verifyAuthMutex.Lock()
	ret, specificReturn := fake.verifyAuthReturnsOnCall[len(fake.verifyAuthArgsForCall)]
	fake.verifyAuthArgsForCall = append(fake.verifyAuthArgsForCall, struct {
		logger *log.Logger
	}{logger})
	fake.recordInvocation("VerifyAuth", []interface{}{logger})
	fake.verifyAuthMutex.Unlock()
	if fake.VerifyAuthStub != nil {
		return fake.VerifyAuthStub(logger)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.verifyAuthReturns.result1
}

func (fake *FakeBoshClient) VerifyAuthCallCount() int {
	fake.verifyAuthMutex.RLock()
	defer fake.verifyAuthMutex.RUnlock()
	return len(fake.verifyAuthArgsForCall)
}

func (fake *FakeBoshClient) VerifyAuthArgsForCall(i int) *log.Logger {
	fake.verifyAuthMutex.RLock()
	defer fake.verifyAuthMutex.RUnlock()
	return fake.verifyAuthArgsForCall[i].logger
}

func (fake *FakeBoshClient) VerifyAuthReturns(result1 error) {
	fake.VerifyAuthStub = nil
	fake.verifyAuthReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBoshClient) VerifyAuthReturnsOnCall(i int, result1 error) {
	fake.VerifyAuthStub = nil
	if fake.verifyAuthReturnsOnCall == nil {
		fake.verifyAuthReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyAuthReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBoshClient) GetDNSAddresses(deploymentName string, requestedDNS []config.BindingDNS) (map[string]string, error) {
	var requestedDNSCopy []config.BindingDNS
	if requestedDNS != nil {
		requestedDNSCopy = make([]config.BindingDNS, len(requestedDNS))
		copy(requestedDNSCopy, requestedDNS)
	}
	fake.getDNSAddressesMutex.Lock()
	ret, specificReturn := fake.getDNSAddressesReturnsOnCall[len(fake.getDNSAddressesArgsForCall)]
	fake.getDNSAddressesArgsForCall = append(fake.getDNSAddressesArgsForCall, struct {
		deploymentName string
		requestedDNS   []config.BindingDNS
	}{deploymentName, requestedDNSCopy})
	fake.recordInvocation("GetDNSAddresses", []interface{}{deploymentName, requestedDNSCopy})
	fake.getDNSAddressesMutex.Unlock()
	if fake.GetDNSAddressesStub != nil {
		return fake.GetDNSAddressesStub(deploymentName, requestedDNS)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDNSAddressesReturns.result1, fake.getDNSAddressesReturns.result2
}

func (fake *FakeBoshClient) GetDNSAddressesCallCount() int {
	fake.getDNSAddressesMutex.RLock()
	defer fake.getDNSAddressesMutex.RUnlock()
	return len(fake.getDNSAddressesArgsForCall)
}

func (fake *FakeBoshClient) GetDNSAddressesArgsForCall(i int) (string, []config.BindingDNS) {
	fake.getDNSAddressesMutex.RLock()
	defer fake.getDNSAddressesMutex.RUnlock()
	return fake.getDNSAddressesArgsForCall[i].deploymentName, fake.getDNSAddressesArgsForCall[i].requestedDNS
}

func (fake *FakeBoshClient) GetDNSAddressesReturns(result1 map[string]string, result2 error) {
	fake.GetDNSAddressesStub = nil
	fake.getDNSAddressesReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetDNSAddressesReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.GetDNSAddressesStub = nil
	if fake.getDNSAddressesReturnsOnCall == nil {
		fake.getDNSAddressesReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.getDNSAddressesReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) Deploy(manifest []byte, contextID string, logger *log.Logger, reporter *boshdirector.AsyncTaskReporter) (int, error) {
	var manifestCopy []byte
	if manifest != nil {
		manifestCopy = make([]byte, len(manifest))
		copy(manifestCopy, manifest)
	}
	fake.deployMutex.Lock()
	ret, specificReturn := fake.deployReturnsOnCall[len(fake.deployArgsForCall)]
	fake.deployArgsForCall = append(fake.deployArgsForCall, struct {
		manifest  []byte
		contextID string
		logger    *log.Logger
		reporter  *boshdirector.AsyncTaskReporter
	}{manifestCopy, contextID, logger, reporter})
	fake.recordInvocation("Deploy", []interface{}{manifestCopy, contextID, logger, reporter})
	fake.deployMutex.Unlock()
	if fake.DeployStub != nil {
		return fake.DeployStub(manifest, contextID, logger, reporter)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deployReturns.result1, fake.deployReturns.result2
}

func (fake *FakeBoshClient) DeployCallCount() int {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	return len(fake.deployArgsForCall)
}

func (fake *FakeBoshClient) DeployArgsForCall(i int) ([]byte, string, *log.Logger, *boshdirector.AsyncTaskReporter) {
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	return fake.deployArgsForCall[i].manifest, fake.deployArgsForCall[i].contextID, fake.deployArgsForCall[i].logger, fake.deployArgsForCall[i].reporter
}

func (fake *FakeBoshClient) DeployReturns(result1 int, result2 error) {
	fake.DeployStub = nil
	fake.deployReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) DeployReturnsOnCall(i int, result1 int, result2 error) {
	fake.DeployStub = nil
	if fake.deployReturnsOnCall == nil {
		fake.deployReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.deployReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTaskMutex.RLock()
	defer fake.getTaskMutex.RUnlock()
	fake.getTasksMutex.RLock()
	defer fake.getTasksMutex.RUnlock()
	fake.getNormalisedTasksByContextMutex.RLock()
	defer fake.getNormalisedTasksByContextMutex.RUnlock()
	fake.vMsMutex.RLock()
	defer fake.vMsMutex.RUnlock()
	fake.getDeploymentMutex.RLock()
	defer fake.getDeploymentMutex.RUnlock()
	fake.getDeploymentsMutex.RLock()
	defer fake.getDeploymentsMutex.RUnlock()
	fake.deleteDeploymentMutex.RLock()
	defer fake.deleteDeploymentMutex.RUnlock()
	fake.getInfoMutex.RLock()
	defer fake.getInfoMutex.RUnlock()
	fake.runErrandMutex.RLock()
	defer fake.runErrandMutex.RUnlock()
	fake.variablesMutex.RLock()
	defer fake.variablesMutex.RUnlock()
	fake.verifyAuthMutex.RLock()
	defer fake.verifyAuthMutex.RUnlock()
	fake.getDNSAddressesMutex.RLock()
	defer fake.getDNSAddressesMutex.RUnlock()
	fake.deployMutex.RLock()
	defer fake.deployMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBoshClient) recordInvocation(key string, args []interface{}) {
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

var _ broker.BoshClient = new(FakeBoshClient)
