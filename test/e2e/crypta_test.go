package crypta_test

import (
	"fmt"
	"io"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

func initDaemon() {
	BeforeAll(func() {
		command := exec.Command(pathToCrypta, "daemon", "start")
		daemon, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(daemon).Should(gexec.Exit(0))
	})

	AfterAll(func() {
		command := exec.Command(pathToCrypta, "daemon", "stop")
		daemon, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Eventually(daemon).Should(gexec.Exit(0))

		// gexec.KillAndWait()
	})
}

func setValue(key, val string) {
	cmd := exec.Command(pathToCrypta, "set", key, val)
	crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Ω(err).ShouldNot(HaveOccurred())

	Eventually(crypta).Should(gexec.Exit(0))
}

func getValue(key string) string {
	cmd := exec.Command(pathToCrypta, "get", key)
	crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Ω(err).ShouldNot(HaveOccurred())

	Eventually(crypta).Should(gexec.Exit(0))

	return string(crypta.Out.Contents())
}

var _ = Describe("Crypta", func() {
	Describe("A set value can be retrieved", Ordered, func() {
		key := "abcd"
		val := "xyz"
		otherVal := "zyx"

		initDaemon()

		It("sets a new value", func() {
			setValue(key, val)
		})

		It("retrieves the previously written value", func() {
			Ω(getValue(key)).Should(Equal(val + "\n"))
		})

		It("sets a new value in interactive mode", func() {
			cmd := exec.Command(pathToCrypta, "set", key)
			stdin, err := cmd.StdinPipe()
			Ω(err).ShouldNot(HaveOccurred())
			crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(crypta.Err).Should(Say(fmt.Sprintf("Value for %s:", key)))

			_, err = stdin.Write([]byte(otherVal))
			Ω(err).ShouldNot(HaveOccurred())
			err = stdin.Close()
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(crypta).Should(gexec.Exit(0))
		})

		It("retrieves the previously written value", func() {
			Ω(getValue(key)).Should(Equal(otherVal + "\n"))
		})
	})

	Describe("Overwriting an already set value", Ordered, func() {
		key := "abcd"
		val := "xyz"
		newVal := "xxx"

		initDaemon()

		It("sets a value", func() {
			setValue(key, val)
		})

		It("overwrites the already set value", func() {
			setValue(key, newVal)
		})

		It("verifies that the value is overwritten", func() {
			Ω(getValue(key)).Should(Equal(newVal + "\n"))
		})
	})

	Describe("Setting a new value if it is not available yet", Ordered, func() {
		var crypta *gexec.Session
		var stdin io.WriteCloser

		val := "xyz"

		initDaemon()

		BeforeEach(func() {
			cmd := exec.Command(pathToCrypta, "get", "abcd")
			var err error
			stdin, err = cmd.StdinPipe()
			Ω(err).ShouldNot(HaveOccurred())
			crypta, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			Eventually(crypta).Should(gexec.Exit(0))
		})

		Context("When no value is set yet", func() {
			It("asks for a new value to be set", func() {
				Eventually(crypta.Err).Should(Say("Value for abcd:"))

				_, err := stdin.Write([]byte(val))
				Ω(err).ShouldNot(HaveOccurred())
				err = stdin.Close()
				Ω(err).ShouldNot(HaveOccurred())

				Eventually(crypta).Should(gexec.Exit(0))
			})
		})

		Context("When the value is then retrieved again", func() {
			It("returns the newly written value", func() {
				Eventually(crypta).Should(Say(val))
			})
		})
	})

	It("should return an error if the daemon is not running when", func() {
		By("trying to retrieve a value", func() {
			cmd := exec.Command(pathToCrypta, "get", "abcd")
			crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(crypta).Should(gexec.Exit(1))
		})
		By("trying to set a value", func() {
			cmd := exec.Command(pathToCrypta, "set", "abcd", "xyz")
			crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(crypta).Should(gexec.Exit(1))
		})
	})
})
