// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package upgrader

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pivotal-cf/on-demand-service-broker/broker/services"
	"github.com/pivotal-cf/on-demand-service-broker/config"
	"github.com/pivotal-cf/on-demand-service-broker/service"
)

type LoggingListener struct {
	logger *log.Logger
}

func NewLoggingListener(logger *log.Logger) Listener {
	return LoggingListener{logger: logger}
}

func (ll LoggingListener) Starting(maxInFlight int) {
	ll.logger.Printf("STARTING UPGRADES with %d concurrent workers\n", maxInFlight)
}

func (ll LoggingListener) RetryAttempt(num, limit int) {
	var remaining string
	typeOfUpgrade := "instances"
	if num > 1 {
		remaining = "remaining "
	}
	ll.logger.Printf("Upgrading all %s%s. Attempt %d/%d\n", remaining, typeOfUpgrade, num, limit)
}

func (ll LoggingListener) RetryCanariesAttempt(attempt, limit, remainingCanaries int) {
	remaining := "all"
	if attempt > 1 {
		remaining = fmt.Sprintf("%d remaining", remainingCanaries)
	}
	ll.logger.Printf("Upgrading %s canaries. Attempt %d/%d\n", remaining, attempt, limit)
}

func (ll LoggingListener) InstancesToUpgrade(instances []service.Instance) {
	msg := "Service Instances:"
	for _, instance := range instances {
		msg = fmt.Sprintf("%s %s", msg, instance.GUID)
	}
	ll.logger.Println(msg)
	ll.logger.Printf("Total Service Instances found in Cloud Foundry: %d\n", len(instances))
}

func (ll LoggingListener) InstanceUpgradeStarting(instance string, index, totalInstances int, isCanary bool) {
	var instanceCount string
	if !isCanary {
		instanceCount = fmt.Sprintf(" %d of %d", index, totalInstances)
	}
	ll.logger.Printf("[%s] Starting to upgrade service instance%s", instance, instanceCount)
}

func (ll LoggingListener) InstanceUpgradeStartResult(instance string, resultType services.UpgradeOperationType) {
	var message string

	switch resultType {
	case services.UpgradeAccepted:
		message = "accepted upgrade"
	case services.InstanceNotFound:
		message = "already deleted from platform"
	case services.OrphanDeployment:
		message = "orphan service instance detected - no corresponding bosh deployment"
	case services.OperationInProgress:
		message = "operation in progress"
	default:
		message = "unexpected result"
	}

	ll.logger.Printf("[%s] Result: %s", instance, message)
}

func (ll LoggingListener) InstanceUpgraded(instance string, result string) {
	ll.logger.Printf("[%s] Result: Service Instance upgrade %s\n", instance, result)
}

func (ll LoggingListener) WaitingFor(instance string, boshTaskId int) {
	ll.logger.Printf("[%s] Waiting for upgrade to complete: bosh task id %d", instance, boshTaskId)
}

func (ll LoggingListener) Progress(pollingInterval time.Duration, orphanCount, upgradedCount, toRetryCount, deletedCount int) {
	ll.logger.Printf("Upgrade progress summary: "+
		"Sleep interval until next attempt: %s; "+
		"Number of successful upgrades so far: %d; "+
		"Number of service instance orphans detected so far: %d; "+
		"Number of deleted instances before upgrade could occur: %d; "+
		"Number of operations in progress (to retry) so far: %d",
		pollingInterval,
		upgradedCount,
		orphanCount,
		deletedCount,
		toRetryCount,
	)
}

func (ll LoggingListener) Finished(orphanCount, upgradedCount, deletedCount int, busyInstances, failedInstances []string) {
	var failedList string
	var busyList string
	if len(failedInstances) > 0 {
		failedList = fmt.Sprintf(" [%s]", strings.Join(failedInstances, ", "))
	}
	if len(busyInstances) > 0 {
		busyList = fmt.Sprintf(" [%s]", strings.Join(busyInstances, ", "))
	}

	status := "SUCCESS"
	if len(failedInstances) > 0 || len(busyInstances) > 0 {
		status = "FAILED"
	}

	ll.logger.Printf("FINISHED UPGRADES Status: %s; Summary: "+
		"Number of successful upgrades: %d; "+
		"Number of service instance orphans detected: %d; "+
		"Number of deleted instances before upgrade could occur: %d; "+
		"Number of busy instances which could not be upgraded: %d%s; "+
		"Number of service instances that failed to upgrade: %d%s",
		status,
		upgradedCount,
		orphanCount,
		deletedCount,
		len(busyInstances),
		busyList,
		len(failedInstances),
		failedList,
	)
}

func (ll LoggingListener) CanariesStarting(canaries int, filter config.CanarySelectionParams) {
	msg := fmt.Sprintf("STARTING CANARY UPGRADES: %d canaries", canaries)
	if len(filter) > 0 {
		msg = fmt.Sprintf("%s with selection criteria: %s", msg, filter)
	}
	ll.logger.Println(msg)
}

func (ll LoggingListener) CanariesFinished() {
	ll.logger.Printf("FINISHED CANARY UPGRADES")
}

func (ll LoggingListener) FailedToRefreshInstanceInfo(instance string) {
	ll.logger.Printf("[%s] Failed to get refreshed list of instances. Continuing with previously fetched info.\n", instance)
}
