package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"os/exec"
)

var _ = Describe("Integration", func() {
	var binaryPath string

	BeforeSuite(func() {
		var err error
		binaryPath, err = gexec.Build("github.com/mmb/gocliboil")
		Expect(err).ToNot(HaveOccurred())
	})

	AfterSuite(func() {
		gexec.CleanupBuildArtifacts()
	})

	Context("when the command is successful", func() {
		It("exits with 0", func() {
			command := exec.Command(binaryPath, "-exitCode", "0")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))
		})

		It("writes exiting 0 to stdout", func() {
			command := exec.Command(binaryPath, "-exitCode", "0")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))
			Expect(session.Out.Contents()).To(ContainSubstring("exiting with 0"))
		})

		It("does not write to stderr", func() {
			command := exec.Command(binaryPath, "-exitCode", "0")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))
			Expect(session.Err.Contents()).To(BeEmpty())
		})
	})

	Context("when the command is not successful", func() {
		It("exits with 1", func() {
			command := exec.Command(binaryPath, "-exitCode", "78")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(78))
		})

		It("does not write to stdout", func() {
			command := exec.Command(binaryPath, "-exitCode", "78")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(78))
			Expect(session.Out.Contents()).To(BeEmpty())
		})

		It("writes exiting 1 to stderr", func() {
			command := exec.Command(binaryPath, "-exitCode", "78")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).ToNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(78))
			Expect(session.Err.Contents()).To(ContainSubstring("exiting with 78"))
		})
	})
})
