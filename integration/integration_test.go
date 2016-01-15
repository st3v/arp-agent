package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("arp-agent", func() {
	It("says hello world", func() {
		session := execBin("--listen", "localhost")
		Eventually(session).Should(gbytes.Say("Hello World"))
	})

	It("exits with exit code 0", func() {
		session := execBin("--listen", "localhost")
		Eventually(session).Should(gexec.Exit(0))
	})

	Context("when --listen is not specified", func() {
		It("fails with exit code 1", func() {
			session := execBin()
			Eventually(session).Should(gexec.Exit(1))
		})

		It("returns a meaningful error message", func() {
			session := execBin()
			Eventually(session.Err).Should(gbytes.Say("Must specify --listen"))
		})
	})
})
