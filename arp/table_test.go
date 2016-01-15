package arp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/arp-agent/arp"
)

var _ = Describe("Table", func() {
	var (
		ip          = "10.61.8.107"
		expectedMac = "aa:aa:aa:aa:aa:bb"
		table       = arp.NewTable()
	)

	Describe(".Insert", func() {
		It("does not return an error", func() {
			err := table.Insert(ip, expectedMac)
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe(".Read", func() {
		BeforeEach(func() {
			err := table.Insert(ip, expectedMac)
			Expect(err).ToNot(HaveOccurred())
		})

		It("does not return an error", func() {
			_, err := table.Read(ip)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns a mac address", func() {
			mac, _ := table.Read(ip)
			Expect(mac).To(Equal(expectedMac))
		})

		Context("for an non-existing IP", func() {
			It("returns an empty string", func() {
				mac, _ := table.Read("1.2.3.4")
				Expect(mac).To(BeEmpty())
			})

			It("does not return an error", func() {
				_, err := table.Read("1.2.3.4")
				Expect(err).ToNot(HaveOccurred())
			})

		})
	})

})
