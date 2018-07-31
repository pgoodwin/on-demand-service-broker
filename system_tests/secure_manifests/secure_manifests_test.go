package secure_manifests_test

import (
	"encoding/json"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/pborman/uuid"
	cf "github.com/pivotal-cf/on-demand-service-broker/system_tests/cf_helpers"
)

var _ = Describe("SecureManifest", func() {
	When("it is disabled", func() {
		var (
			plan           string
			serviceName    string
			serviceKeyName string
			skipDelete     bool
		)

		BeforeEach(func() {
			manifest := boshClient.GetManifest(brokerDeploymentName)
			for i := range manifest.InstanceGroups {
				if manifest.InstanceGroups[i].Name == "broker" {
					if manifest.InstanceGroups[i].Jobs[0].Properties["enable_secure_manifests"].(bool) {
						manifest.InstanceGroups[i].Jobs[0].Properties["enable_secure_manifests"] = false
						boshClient.DeployODB(*manifest)
						break
					}
				}
			}

			serviceName = os.Getenv("SERVICE_INSTANCE_NAME")
			if serviceName == "" {
				serviceName = fmt.Sprintf("instance-%s", uuid.New()[:7])
				plan = "dedicated-vm"
				cf.CreateService(serviceOffering, plan, serviceName, "")
			} else {
				skipDelete = true
			}

			serviceKeyName = uuid.New()[:7]
		})

		AfterEach(func() {
			Eventually(
				cf.Cf("delete-service-key", "-f", serviceName, serviceKeyName),
				cf.CfTimeout,
			).Should(gexec.Exit(0))

			if !skipDelete {
				cf.DeleteService(serviceName)
			}
		})

		It("does not tries to use ODB managed secrets", func() {
			var serviceKey map[string]interface{}

			By("creating a service key", func() {
				cf.CreateServiceKey(serviceName, serviceKeyName)
				serviceKeyRaw := cf.GetServiceKey(serviceName, serviceKeyName)
				Expect(serviceKeyRaw).NotTo(BeEmpty())

				err := json.Unmarshal([]byte(serviceKeyRaw), &serviceKey)
				Expect(err).NotTo(HaveOccurred())
			})

			managedSecret, found := serviceKey["odb_managed_secret"]
			Expect(found).To(BeTrue())
			Expect(managedSecret.(string)).NotTo(BeEmpty())
		})
	})
})
