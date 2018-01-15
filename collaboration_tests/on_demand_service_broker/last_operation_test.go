package on_demand_service_broker_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-cf/on-demand-service-broker/boshdirector"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
	brokerConfig "github.com/pivotal-cf/on-demand-service-broker/config"
	sdk "github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
	"github.com/pkg/errors"
)

var _ = Describe("Last Operation", func() {

	const (
		instanceID           = "some-instance-id"
		postDeployErrandName = "post-deploy-errand"
		boshTaskID           = 6782
	)

	var (
		operationData        broker.OperationData
		processingTask       = boshdirector.BoshTask{ID: boshTaskID, State: boshdirector.TaskProcessing}
		doneTask             = boshdirector.BoshTask{ID: boshTaskID, State: boshdirector.TaskDone, Description: "running"}
		failedTask           = boshdirector.BoshTask{ID: boshTaskID, State: boshdirector.TaskError, Description: "failed"}
		failedErrandTask     = boshdirector.BoshTask{ID: boshTaskID + 1, State: boshdirector.TaskError, Description: "failed"}
		doneErrandTask       = boshdirector.BoshTask{ID: boshTaskID + 1, State: boshdirector.TaskDone, Description: "errand completed"}
		processingErrandTask = boshdirector.BoshTask{ID: boshTaskID + 1, State: boshdirector.TaskProcessing, Description: "errand running"}
	)

	Context("when no lifecycle errands are configured", func() {
		BeforeEach(func() {
			conf := brokerConfig.Config{
				Broker: brokerConfig.Broker{
					Port: serverPort, Username: brokerUsername, Password: brokerPassword,
				},
				ServiceCatalog: brokerConfig.ServiceOffering{
					Name: serviceName,
				},
			}

			StartServer(conf)
		})

		DescribeTable("depending on the operation type, responds with 200 when", func(operationType broker.OperationType, action, responseDescription string) {
			fakeBoshClient.GetTaskReturns(boshdirector.BoshTask{ID: boshTaskID, State: boshdirector.TaskProcessing}, nil)

			operationData := broker.OperationData{
				BoshTaskID:    boshTaskID,
				OperationType: operationType,
				BoshContextID: "",
				PlanID:        dedicatedPlanID,
				PostDeployErrand: broker.PostDeployErrand{
					Name: postDeployErrandName,
				},
			}

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("returning the correct error description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(responseBody).To(MatchJSON(fmt.Sprintf(`
				{
					"state":       "in progress",
					"description": "%s"
				}`,
				responseDescription,
			)))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(
				fmt.Sprintf(`BOSH task ID %d status: processing %s deployment for instance %s`,
					boshTaskID,
					action,
					instanceID,
				),
			))
		},
			Entry("create is in progress", broker.OperationTypeCreate, "create", "Instance provisioning in progress"),
			Entry("delete is in progress", broker.OperationTypeDelete, "delete", "Instance deletion in progress"),
			Entry("update is in progress", broker.OperationTypeUpdate, "update", "Instance update in progress"),
			Entry("upgrade is in progress", broker.OperationTypeUpgrade, "upgrade", "Instance upgrade in progress"),
		)

		DescribeTable("depending on the task state, responds with 200 when", func(taskState, responseState, description string) {
			fakeBoshClient.GetTaskReturns(boshdirector.BoshTask{ID: boshTaskID, State: taskState}, nil)

			operationData := broker.OperationData{
				BoshTaskID:    boshTaskID,
				OperationType: broker.OperationTypeCreate,
				BoshContextID: "",
				PlanID:        dedicatedPlanID,
				PostDeployErrand: broker.PostDeployErrand{
					Name: postDeployErrandName,
				},
			}

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("returning the correct error description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())

			var unmarshalled map[string]interface{}
			json.Unmarshal(responseBody, &unmarshalled)

			Expect(unmarshalled["state"]).To(Equal(responseState))

			if responseState != string(brokerapi.Failed) {
				Expect(unmarshalled["description"]).To(Equal(description))
			} else {
				Expect(unmarshalled["description"]).To(SatisfyAll(
					ContainSubstring("Instance provisioning failed: There was a problem completing your request. Please contact your operations team providing the following information:"),
					MatchRegexp(`broker-request-id: [0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`),
					ContainSubstring(fmt.Sprintf("service: %s", serviceName)),
					ContainSubstring(fmt.Sprintf("service-instance-guid: %s", instanceID)),
					ContainSubstring("operation: create"),
					ContainSubstring(fmt.Sprintf("task-id: %d", boshTaskID)),
				))
			}

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(
				fmt.Sprintf(`BOSH task ID %d status: %s create deployment for instance %s`,
					boshTaskID,
					taskState,
					instanceID,
				),
			))
		},
			Entry("a task is processing", boshdirector.TaskProcessing, string(brokerapi.InProgress), "Instance provisioning in progress"),
			Entry("a task is done", boshdirector.TaskDone, string(brokerapi.Succeeded), "Instance provisioning completed"),
			Entry("a task is cancelling", boshdirector.TaskCancelling, string(brokerapi.InProgress), "Instance provisioning in progress"),
			Entry("a task has timed out", boshdirector.TaskTimeout, string(brokerapi.Failed), ""),
			Entry("a task is cancelled", boshdirector.TaskCancelled, string(brokerapi.Failed), ""),
			Entry("a task has errored", boshdirector.TaskError, string(brokerapi.Failed), ""),
			Entry("a task has an unrecognised state", "other-state", string(brokerapi.Failed), ""),
		)

		It("responds with 500 if BOSH fails to get the task", func() {
			fakeBoshClient.GetTaskReturns(boshdirector.BoshTask{}, errors.New("oops"))
			operationData := broker.OperationData{
				BoshTaskID:    boshTaskID,
				OperationType: broker.OperationTypeCreate,
				BoshContextID: "",
				PlanID:        dedicatedPlanID,
				PostDeployErrand: broker.PostDeployErrand{
					Name: postDeployErrandName,
				},
			}

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusInternalServerError))

			By("returning the correct error description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())

			var unmarshalled map[string]interface{}
			json.Unmarshal(responseBody, &unmarshalled)

			Expect(unmarshalled["description"]).To(SatisfyAll(
				ContainSubstring("There was a problem completing your request. Please contact your operations team providing the following information:"),
				MatchRegexp(`broker-request-id: [0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`),
				ContainSubstring(fmt.Sprintf("service: %s", serviceName)),
				ContainSubstring(fmt.Sprintf("service-instance-guid: %s", instanceID)),
				ContainSubstring("operation: create"),
				ContainSubstring(fmt.Sprintf("task-id: %d", boshTaskID)),
			))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(fmt.Sprintf(`error: error retrieving tasks from bosh, for deployment 'service-instance_%s'`, instanceID)))
		})

		It("responds with 500 if Cloud Controller does not send operation data", func() {
			fakeBoshClient.GetTaskReturns(boshdirector.BoshTask{ID: boshTaskID, State: boshdirector.TaskProcessing}, nil)
			response := doLastOperationRequest(instanceID, broker.OperationData{})

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusInternalServerError))

			By("returning the correct error description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())

			var unmarshalled map[string]interface{}
			json.Unmarshal(responseBody, &unmarshalled)

			Expect(unmarshalled["description"]).To(SatisfyAll(
				ContainSubstring("There was a problem completing your request. Please contact your operations team providing the following information:"),
				MatchRegexp(`broker-request-id: [0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`),
				ContainSubstring(fmt.Sprintf("service: %s", serviceName)),
				ContainSubstring(fmt.Sprintf("service-instance-guid: %s", instanceID)),
				Not(ContainSubstring("operation:")),
				Not(ContainSubstring("task-id")),
			))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say("Request missing operation data, please check your Cloud Foundry version is v238+"))
		})
	})

	Context("when post-deploy errand is configured", func() {
		const (
			planID        = "post-deploy-plan-id"
			operationType = broker.OperationTypeCreate
			contextID     = "some-context-id"
			errandName    = "post-deploy-errand"
		)

		BeforeEach(func() {
			planWithPostDeploy := brokerConfig.Plan{
				ID:   planID,
				Name: "post-deploy-plan",
				LifecycleErrands: &sdk.LifecycleErrands{
					PostDeploy: sdk.Errand{
						Name:      errandName,
						Instances: []string{"instance-group-name/0"},
					},
				},
			}
			conf := brokerConfig.Config{
				Broker: brokerConfig.Broker{
					Port: serverPort, Username: brokerUsername, Password: brokerPassword,
				},
				ServiceCatalog: brokerConfig.ServiceOffering{
					Name:  serviceName,
					Plans: []brokerConfig.Plan{planWithPostDeploy},
				},
			}

			operationData = broker.OperationData{
				BoshTaskID:    doneTask.ID,
				OperationType: operationType,
				BoshContextID: contextID,
				PlanID:        planID,
				PostDeployErrand: broker.PostDeployErrand{
					Name: postDeployErrandName,
				},
			}
			StartServer(conf)
		})

		It("returns 200 when there is a single incomplete task", func() {
			fakeBoshClient.GetNormalisedTasksByContextReturns(boshdirector.BoshTasks{processingTask}, nil)

			operationData.BoshTaskID = processingTask.ID

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("returning the correct response description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(responseBody).To(MatchJSON(`
				{
					"state":       "in progress",
					"description": "Instance provisioning in progress"
				}`,
			))

			By("not running the post deploy errand")
			Expect(fakeBoshClient.RunErrandCallCount()).To(Equal(0))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(
				fmt.Sprintf(`BOSH task ID %d status: processing create deployment for instance %s`,
					boshTaskID,
					instanceID,
				),
			))
		})

		It("returns 200 when there is a single complete task", func() {
			fakeBoshClient.GetNormalisedTasksByContextReturns(boshdirector.BoshTasks{doneTask}, nil)
			fakeBoshClient.RunErrandReturns(processingTask.ID, nil)
			fakeBoshClient.GetTaskReturns(processingTask, nil)

			operationData.BoshTaskID = doneTask.ID

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("returning the correct response description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(responseBody).To(MatchJSON(`
				{
					"state":       "in progress",
					"description": "Instance provisioning in progress"
				}`,
			))

			By("running the post deploy errand")
			Expect(fakeBoshClient.RunErrandCallCount()).To(Equal(1))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(
				fmt.Sprintf(`BOSH task ID %d status: processing create deployment for instance %s: Description: %s`,
					boshTaskID,
					instanceID,
					processingTask.Description,
				),
			))
		})

		It("returns 200 when there are two tasks and the errand task is still running", func() {
			fakeBoshClient.GetNormalisedTasksByContextReturns(boshdirector.BoshTasks{processingErrandTask, doneTask}, nil)

			operationData.BoshTaskID = doneTask.ID

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("returning the correct response description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(responseBody).To(MatchJSON(`
				{
					"state":       "in progress",
					"description": "Instance provisioning in progress"
				}`,
			))

			By("not running the post deploy errand")
			Expect(fakeBoshClient.RunErrandCallCount()).To(Equal(0))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(
				fmt.Sprintf(`BOSH task ID %d status: processing create deployment for instance %s: Description: %s`,
					processingErrandTask.ID,
					instanceID,
					processingErrandTask.Description,
				),
			))
		})

		It("returns 200 when there are two tasks and the errand task has failed", func() {
			fakeBoshClient.GetNormalisedTasksByContextReturns(boshdirector.BoshTasks{failedErrandTask, doneTask}, nil)

			operationData.BoshTaskID = doneTask.ID

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("returning the correct response description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())

			var parsedResponse map[string]interface{}
			json.Unmarshal(responseBody, &parsedResponse)
			Expect(parsedResponse["state"]).To(Equal(string(brokerapi.Failed)))

			Expect(parsedResponse["description"]).To(SatisfyAll(
				ContainSubstring("Instance provisioning failed: There was a problem completing your request. Please contact your operations team providing the following information:"),
				MatchRegexp(`broker-request-id: [0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`),
				ContainSubstring(fmt.Sprintf("service: %s", serviceName)),
				ContainSubstring(fmt.Sprintf("service-instance-guid: %s", instanceID)),
				ContainSubstring("operation: create"),
				ContainSubstring(fmt.Sprintf("task-id: %d", failedErrandTask.ID)),
			))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(
				fmt.Sprintf(`BOSH task ID %d status: error create deployment for instance %s: Description: %s`,
					failedErrandTask.ID,
					instanceID,
					failedErrandTask.Description),
			))
		})

		It("returns 200 when there are two tasks and the errand task has succeeded", func() {
			fakeBoshClient.GetNormalisedTasksByContextReturns(boshdirector.BoshTasks{doneErrandTask, doneTask}, nil)

			operationData.BoshTaskID = doneTask.ID

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("returning the correct response description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(responseBody).To(MatchJSON(`
				{
					"state":       "succeeded",
					"description": "Instance provisioning completed"
				}`,
			))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(
				fmt.Sprintf(`BOSH task ID %d status: done create deployment for instance %s: Description: %s`,
					doneErrandTask.ID,
					instanceID,
					doneErrandTask.Description),
			))
		})

	})

	Context("when pre-delete errand is configured", func() {
		const (
			planID        = "pre-delete-plan-id"
			operationType = broker.OperationTypeDelete
			contextID     = "some-context-id"
			errandName    = "pre-delete-errand"
		)

		BeforeEach(func() {
			planWithPreDelete := brokerConfig.Plan{
				ID:   planID,
				Name: "pre-delete-plan",
				LifecycleErrands: &sdk.LifecycleErrands{
					PreDelete: sdk.Errand{
						Name:      errandName,
						Instances: []string{"instance-group-name/0"},
					},
				},
			}
			conf := brokerConfig.Config{
				Broker: brokerConfig.Broker{
					Port: serverPort, Username: brokerUsername, Password: brokerPassword,
				},
				ServiceCatalog: brokerConfig.ServiceOffering{
					Name:  serviceName,
					Plans: []brokerConfig.Plan{planWithPreDelete},
				},
			}

			operationData = broker.OperationData{
				BoshTaskID:    doneTask.ID,
				OperationType: operationType,
				BoshContextID: contextID,
				PlanID:        planID,
				PostDeployErrand: broker.PostDeployErrand{
					Name: errandName,
				},
			}
			StartServer(conf)
		})

		It("returns 200 when there is a single incomplete task", func() {
			fakeBoshClient.GetNormalisedTasksByContextReturns(boshdirector.BoshTasks{processingErrandTask}, nil)

			operationData.BoshTaskID = processingErrandTask.ID

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("returning the correct response description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(responseBody).To(MatchJSON(`
				{
					"state":       "in progress",
					"description": "Instance deletion in progress"
				}`,
			))

			By("not running the delete deployment")
			Expect(fakeBoshClient.DeleteDeploymentCallCount()).To(Equal(0))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(
				fmt.Sprintf(`BOSH task ID %d status: processing delete deployment for instance %s`,
					processingErrandTask.ID,
					instanceID,
				),
			))
		})

		It("returns 200 when there is a single complete task", func() {
			fakeBoshClient.GetNormalisedTasksByContextReturns(boshdirector.BoshTasks{doneErrandTask}, nil)
			fakeBoshClient.DeleteDeploymentReturns(processingTask.ID, nil)
			fakeBoshClient.GetTaskReturns(processingTask, nil)

			operationData.BoshTaskID = doneErrandTask.ID

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("returning the correct response description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(responseBody).To(MatchJSON(`
				{
					"state":       "in progress",
					"description": "Instance deletion in progress"
				}`,
			))

			By("running the delete deployment")
			Expect(fakeBoshClient.DeleteDeploymentCallCount()).To(Equal(1))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(
				fmt.Sprintf(`BOSH task ID %d status: processing delete deployment for instance %s: Description: %s`,
					processingTask.ID,
					instanceID,
					processingTask.Description,
				),
			))
		})

		It("returns 200 when there are two tasks and the delete task is still running", func() {
			fakeBoshClient.GetNormalisedTasksByContextReturns(boshdirector.BoshTasks{processingTask, doneErrandTask}, nil)

			operationData.BoshTaskID = doneErrandTask.ID

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("returning the correct response description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(responseBody).To(MatchJSON(`
				{
					"state":       "in progress",
					"description": "Instance deletion in progress"
				}`,
			))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(
				fmt.Sprintf(`BOSH task ID %d status: processing delete deployment for instance %s: Description: %s`,
					processingTask.ID,
					instanceID,
					processingTask.Description,
				),
			))
		})

		It("returns 200 when there are two tasks and the delete task has failed", func() {
			fakeBoshClient.GetNormalisedTasksByContextReturns(boshdirector.BoshTasks{failedTask, doneErrandTask}, nil)

			operationData.BoshTaskID = doneErrandTask.ID

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("returning the correct response description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())

			var parsedResponse map[string]interface{}
			json.Unmarshal(responseBody, &parsedResponse)
			Expect(parsedResponse["state"]).To(Equal(string(brokerapi.Failed)))

			Expect(parsedResponse["description"]).To(SatisfyAll(
				ContainSubstring("Instance deletion failed: There was a problem completing your request. Please contact your operations team providing the following information:"),
				MatchRegexp(`broker-request-id: [0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`),
				ContainSubstring(fmt.Sprintf("service: %s", serviceName)),
				ContainSubstring(fmt.Sprintf("service-instance-guid: %s", instanceID)),
				ContainSubstring("operation: delete"),
				ContainSubstring(fmt.Sprintf("task-id: %d", failedTask.ID)),
			))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(
				fmt.Sprintf(`BOSH task ID %d status: error delete deployment for instance %s: Description: %s`,
					failedTask.ID,
					instanceID,
					failedTask.Description),
			))
		})

		It("returns 200 when there are two tasks and the delete task has succeeded", func() {
			fakeBoshClient.GetNormalisedTasksByContextReturns(boshdirector.BoshTasks{doneTask, doneErrandTask}, nil)

			operationData.BoshTaskID = doneTask.ID

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("returning the correct response description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(responseBody).To(MatchJSON(`
				{
					"state":       "succeeded",
					"description": "Instance deletion completed"
				}`,
			))

			By("logging the appropriate message")
			Eventually(loggerBuffer).Should(gbytes.Say(
				fmt.Sprintf(`BOSH task ID %d status: done delete deployment for instance %s: Description: %s`,
					doneTask.ID,
					instanceID,
					doneTask.Description),
			))
		})

	})

	Context("depending on the setting of the context id on the request body", func() {
		const (
			planID        = "post-deploy-plan-id"
			operationType = broker.OperationTypeCreate
			errandName    = "post-deploy-errand"
		)
		BeforeEach(func() {
			planWithPostDeploy := brokerConfig.Plan{
				ID:   planID,
				Name: "post-deploy-plan",
				LifecycleErrands: &sdk.LifecycleErrands{
					PostDeploy: sdk.Errand{
						Name:      errandName,
						Instances: []string{"instance-group-name/0"},
					},
				},
			}

			conf := brokerConfig.Config{
				Broker: brokerConfig.Broker{
					Port: serverPort, Username: brokerUsername, Password: brokerPassword,
				},
				ServiceCatalog: brokerConfig.ServiceOffering{
					Name:  serviceName,
					Plans: []brokerConfig.Plan{planWithPostDeploy},
				},
			}
			StartServer(conf)

		})

		It("doesn't run the lifecycle errands when it's empty", func() {
			operationData = broker.OperationData{
				BoshTaskID:    boshTaskID,
				OperationType: operationType,
				PlanID:        planID,
			}
			fakeBoshClient.GetTaskReturns(doneTask, nil)

			operationData.BoshTaskID = doneTask.ID

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("not running the post deploy errand")
			Expect(fakeBoshClient.RunErrandCallCount()).To(Equal(0))

			By("returning the correct response description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(responseBody).To(MatchJSON(`
				{
					"state":       "succeeded",
					"description": "Instance provisioning completed"
				}`,
			))
		})

		It("runs the lifecycle errands when it's set", func() {
			operationData = broker.OperationData{
				BoshTaskID:    boshTaskID,
				OperationType: operationType,
				BoshContextID: "some-context-id",
				PlanID:        planID,
			}

			fakeBoshClient.GetNormalisedTasksByContextReturns(boshdirector.BoshTasks{doneTask}, nil)
			fakeBoshClient.RunErrandReturns(processingTask.ID, nil)
			fakeBoshClient.GetTaskReturns(processingTask, nil)

			operationData.BoshTaskID = doneTask.ID

			response := doLastOperationRequest(instanceID, operationData)

			By("returning the correct HTTP status code")
			Expect(response.StatusCode).To(Equal(http.StatusOK))

			By("running the post deploy errand")
			Expect(fakeBoshClient.RunErrandCallCount()).To(Equal(1))

			By("returning the correct response description")
			var responseBody []byte
			responseBody, err := ioutil.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(responseBody).To(MatchJSON(`
				{
					"state":       "in progress",
					"description": "Instance provisioning in progress"
				}`,
			))
		})
	})

})

func doLastOperationRequest(instanceID string, operationData broker.OperationData) *http.Response {
	lastOperationURL := fmt.Sprintf("http://%s/v2/service_instances/%s/last_operation", serverURL, instanceID)

	if operationData.PlanID != "" {
		operationDataBytes, err := json.Marshal(operationData)
		Expect(err).NotTo(HaveOccurred())
		lastOperationURL = fmt.Sprintf("%s?operation=%s", lastOperationURL, url.QueryEscape(string(operationDataBytes)))
	}

	lastOperationRequest, err := http.NewRequest(http.MethodGet, lastOperationURL, nil)
	Expect(err).NotTo(HaveOccurred())

	lastOperationRequest.SetBasicAuth(brokerUsername, brokerPassword)

	lastOperationResponse, err := http.DefaultClient.Do(lastOperationRequest)
	Expect(err).NotTo(HaveOccurred())

	return lastOperationResponse
}
