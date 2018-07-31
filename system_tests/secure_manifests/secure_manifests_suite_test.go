package secure_manifests_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	bosh "github.com/pivotal-cf/on-demand-service-broker/system_tests/bosh_helpers"
)

var (
	brokerName           string
	brokerDeploymentName string
	brokerURL            string
	brokerUsername       string
	brokerPassword       string
	serviceOffering      string
	serviceID            string
	boshClient           *bosh.BoshHelperClient
)

func TestSecureManifests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SecureManifests Suite")
}

var _ = BeforeSuite(func() {
	parseEnv()
	// Eventually(cf.Cf("create-service-broker", brokerName, brokerUsername, brokerPassword, brokerURL), cf.CfTimeout).Should(gexec.Exit(0))
	// Eventually(cf.Cf("enable-service-access", serviceOffering), cf.CfTimeout).Should(gexec.Exit(0))
	boshURL := envMustHave("BOSH_URL")
	boshUsername := envMustHave("BOSH_USERNAME")
	boshPassword := envMustHave("BOSH_PASSWORD")
	uaaURL := os.Getenv("UAA_URL")
	boshCACert := os.Getenv("BOSH_CA_CERT_FILE")
	disableTLSVerification := boshCACert == ""

	if uaaURL == "" {
		boshClient = bosh.NewBasicAuth(boshURL, boshUsername, boshPassword, boshCACert, disableTLSVerification)
	} else {
		boshClient = bosh.New(boshURL, uaaURL, boshUsername, boshPassword, boshCACert)
	}
})

func parseEnv() {
	brokerName = envMustHave("BROKER_NAME")
	brokerDeploymentName = envMustHave("BROKER_DEPLOYMENT_NAME")
	brokerUsername = envMustHave("BROKER_USERNAME")
	brokerPassword = envMustHave("BROKER_PASSWORD")
	brokerURL = envMustHave("BROKER_URL")
	serviceOffering = envMustHave("SERVICE_OFFERING_NAME")
	serviceID = envMustHave("SERVICE_OFFERING_ID")
}

func envMustHave(key string) string {
	value := os.Getenv(key)
	Expect(value).ToNot(BeEmpty(), fmt.Sprintf("must set %s", key))
	return value
}
