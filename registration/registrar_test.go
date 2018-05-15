// Copyright (C) 2015-Present Pivotal Software, Inc. All rights reserved.

// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registration_test

import (
	"fmt"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/pivotal-cf/on-demand-service-broker/config"
	"github.com/pivotal-cf/on-demand-service-broker/registration"
	"github.com/pivotal-cf/on-demand-service-broker/registration/fakes"

	"errors"

	"os"

	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Registrar Config", func() {
	It("loads valid config", func() {
		cwd, err := os.Getwd()
		Expect(err).NotTo(HaveOccurred())

		configFilePath := filepath.Join(cwd, "fixtures", "deregister_config.yml")

		configFileBytes, err := ioutil.ReadFile(configFilePath)
		Expect(err).NotTo(HaveOccurred())

		var deregistrarConfig registration.Config
		err = yaml.Unmarshal(configFileBytes, &deregistrarConfig)
		Expect(err).NotTo(HaveOccurred())

		expected := registration.Config{
			DisableSSLCertVerification: true,
			CF: config.CF{
				URL:         "some-cf-url",
				TrustedCert: "some-cf-cert",
				Authentication: config.Authentication{
					UAA: config.UAAAuthentication{
						URL: "a-uaa-url",
						UserCredentials: config.UserCredentials{
							Username: "some-cf-username",
							Password: "some-cf-password",
						},
					},
				},
			},
		}

		Expect(expected).To(Equal(deregistrarConfig))
	})
})

var _ = Describe("Broker registration", func() {
	const (
		brokerGUID     = "broker-guid"
		brokerName     = "broker-name"
		brokerUsername = "foo"
		brokerPassword = "foo"
		brokerURI      = "foo"
		serviceName    = "redis"
	)

	var fakeCFClient *fakes.FakeCloudFoundryClient

	BeforeEach(func() {
		fakeCFClient = new(fakes.FakeCloudFoundryClient)
	})

	Describe("Register", func() {
		It("can register a new broker with one plan", func() {
			registrar := registration.New(fakeCFClient, nil)
			Expect(registrar.Register(brokerName, brokerUsername, brokerPassword, brokerURI, serviceName)).NotTo(HaveOccurred())

			Expect(fakeCFClient.CreateServiceBrokerCallCount()).To(Equal(1))
			expectedName, expectedUsername, expectedPassword, expectedURI, expectedServiceName := fakeCFClient.CreateServiceBrokerArgsForCall(0)
			Expect(expectedName).To(Equal(brokerName))
			Expect(expectedUsername).To(Equal(brokerUsername))
			Expect(expectedPassword).To(Equal(brokerPassword))
			Expect(expectedURI).To(Equal(brokerURI))
			Expect(expectedServiceName).To(Equal(serviceName))

			Expect(fakeCFClient.EnableServiceAccessCallCount()).To(Equal(1))
		})

		It("returns an error when registering broker fails", func() {
			fakeCFClient.CreateServiceBrokerReturns(fmt.Errorf("CF is borked"))

			registrar := registration.New(fakeCFClient, nil)
			expectedErr := registrar.Register(brokerName, brokerUsername, brokerPassword, brokerURI, serviceName)

			Expect(expectedErr).To(HaveOccurred())
			Expect(expectedErr.Error()).To(ContainSubstring("CF is borked"))
		})

		It("returns an error when enabling service access fails", func() {
			fakeCFClient.EnableServiceAccessReturns(fmt.Errorf("CF is borked"))

			registrar := registration.New(fakeCFClient, nil)
			expectedErr := registrar.Register(brokerName, brokerUsername, brokerPassword, brokerURI, serviceName)

			Expect(expectedErr).To(HaveOccurred())
			Expect(expectedErr.Error()).To(ContainSubstring("CF is borked"))
		})
	})

	Describe("Deregister", func() {
		It("does not return an error when deregistering", func() {
			fakeCFClient.GetServiceOfferingGUIDReturns(brokerGUID, nil)

			registrar := registration.New(fakeCFClient, nil)

			Expect(registrar.Deregister(brokerName)).NotTo(HaveOccurred())
			Expect(fakeCFClient.GetServiceOfferingGUIDCallCount()).To(Equal(1))
			Expect(fakeCFClient.DeregisterBrokerCallCount()).To(Equal(1))
			Expect(fakeCFClient.DeregisterBrokerArgsForCall(0)).To(Equal(brokerGUID))
		})

		It("returns an error when cf client fails to ge the service offering guid", func() {
			fakeCFClient.GetServiceOfferingGUIDReturns("", errors.New("list service broker failed"))

			registrar := registration.New(fakeCFClient, nil)

			Expect(registrar.Deregister(brokerName)).To(MatchError("list service broker failed"))
		})

		It("returns an error when cf client fails to deregister", func() {
			fakeCFClient.GetServiceOfferingGUIDReturns(brokerGUID, nil)
			fakeCFClient.DeregisterBrokerReturns(errors.New("failed"))

			registrar := registration.New(fakeCFClient, nil)

			errMsg := fmt.Sprintf("Failed to deregister broker with %s with guid %s, err: failed", brokerName, brokerGUID)
			Expect(registrar.Deregister(brokerName)).To(MatchError(errMsg))
		})
	})

})
