package arp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestArp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Arp Suite")
}
