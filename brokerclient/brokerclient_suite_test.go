package brokerclient_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBrokerclient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Broker Client Suite")
}
