package crypta_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

func initDaemon(setCmdEnv func(*exec.Cmd)) {
	BeforeAll(func() {
		command := exec.Command(pathToCrypta, "daemon")
		setCmdEnv(command)
		daemon, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(daemon).ShouldNot(gexec.Exit(0))
	})

	AfterAll(func() {
		gexec.KillAndWait()
	})
}

func defaultEnv(*exec.Cmd) {}

func setTestDaemonTimeout(c *exec.Cmd) {
	c.Env = os.Environ()
	c.Env = append(c.Env, "CRYPTA_TIMEOUT=0.1")
}

func setValue(key, val string) {
	cmd := exec.Command(pathToCrypta, "set", key, val)
	setTestDaemonTimeout(cmd)
	crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Ω(err).ShouldNot(HaveOccurred())

	Eventually(crypta).Should(gexec.Exit(0))
}

func getValue(key string) string {
	cmd := exec.Command(pathToCrypta, "get", key)
	setTestDaemonTimeout(cmd)
	crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Ω(err).ShouldNot(HaveOccurred())

	Eventually(crypta).Should(gexec.Exit(0))

	return string(crypta.Out.Contents())
}

var _ = Describe("Crypta", func() {
	Describe("Version can be read if", Ordered, func() {
		It("is called with the version flag", func() {
			cmd := exec.Command(pathToCrypta, "--version")
			crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(crypta.Out).Should(Say("0.0.0"))
			Eventually(crypta).Should(gexec.Exit(0))
		})
	})

	Describe("A set value can be retrieved", Ordered, func() {
		key := "abcd"
		val := "xyz"
		otherVal := "zyx"

		initDaemon(defaultEnv)

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

		initDaemon(defaultEnv)

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

		initDaemon(defaultEnv)

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
		checkErrOutputForConnectionTimeout := func(out string) {
			Expect(out).To(ContainSubstring("Using set maximum connection timeout: 100ms"))
		}

		checkErrOutputForRetries := func(out string) {
			for i := 1; i <= 4; i++ {
				Expect(out).To(ContainSubstring(fmt.Sprintf("Daemon currently not reachable. Retry %d of 4...", i)))
			}
		}

		startCrypta := func(args ...string) string {
			cmd := exec.Command(pathToCrypta, args...)
			setTestDaemonTimeout(cmd)
			crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(crypta).WithTimeout(3 * time.Second).Should(gexec.Exit(1))

			return string(crypta.Err.Contents())
		}

		By("trying to retrieve a value", func() {
			errOutput := startCrypta("get", "abcd")
			Expect(errOutput).To(ContainSubstring(`Error: the daemon does not seem to be running: unavailable: POST http://127.0.0.1:35997/secret.v1.SecretService/GetSecret giving up after 5 attempt(s): Post "http://127.0.0.1:35997/secret.v1.SecretService/GetSecret": dial tcp 127.0.0.1:35997: connect: connection refused`))
		})
		By("trying to set a value", func() {
			errOutput := startCrypta("set", "abcd", "xyz")
			Expect(errOutput).To(ContainSubstring(`Error: unavailable: POST http://127.0.0.1:35997/secret.v1.SecretService/SetSecret giving up after 5 attempt(s): Post "http://127.0.0.1:35997/secret.v1.SecretService/SetSecret": dial tcp 127.0.0.1:35997: connect: connection refused`))
		})
		By("trying to retrieve a value with verbose mode enabled", func() {
			errOutput := startCrypta("get", "abcd", "-v")
			checkErrOutputForConnectionTimeout(errOutput)
			checkErrOutputForRetries(errOutput)
		})
		By("trying to set a value with verbose mode enabled", func() {
			errOutput := startCrypta("set", "abcd", "xyz", "-v")
			checkErrOutputForConnectionTimeout(errOutput)
			checkErrOutputForRetries(errOutput)
		})
		By("trying to retrieve a value and not print the usage", func() {
			errOutput := startCrypta("get", "abcd")
			Expect(errOutput).To(Not(BeEmpty()))
			Expect(errOutput).To(Not(ContainSubstring("Usage:")))
		})
		By("trying to set a value and not print the usage", func() {
			errOutput := startCrypta("set", "abcd", "xyz")
			Expect(errOutput).To(Not(BeEmpty()))
			Expect(errOutput).To(Not(ContainSubstring("Usage:")))
		})
	})

	Describe("An initial value set via environment variables can be retrieved", Ordered, func() {
		key := "abcd"
		val := "xyz"

		initDaemon(func(c *exec.Cmd) {
			c.Env = os.Environ()

			// set secret as environment variable
			c.Env = append(c.Env, fmt.Sprintf("CRYPTA_SECRET_%s=%s", key, val))
		})

		It("retrieves the initial value set via the environment variable", func() {
			Ω(getValue(key)).Should(Equal(val + "\n"))
		})
	})

	Describe("A usage message shall be printed if", func() {
		It("receives a wrong flag", func() {
			cmd := exec.Command(pathToCrypta, "get", "--unknown")
			crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(crypta).Should(gexec.Exit(1))

			errStr := string(crypta.Err.Contents())

			Expect(crypta.Out.Contents()).To(BeEmpty())
			Expect(errStr).To(Not(BeEmpty()))
			Expect(errStr).To(ContainSubstring("Usage:"))
		})

		It("is missing a required argument", func() {
			cmd := exec.Command(pathToCrypta, "get")
			crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(crypta).Should(gexec.Exit(1))

			errStr := string(crypta.Err.Contents())

			Expect(crypta.Out.Contents()).To(BeEmpty())
			Expect(errStr).To(Not(BeEmpty()))
			Expect(errStr).To(ContainSubstring("Usage:"))
		})

		It("is called with help", func() {
			cmd := exec.Command(pathToCrypta, "help")
			crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(crypta).Should(gexec.Exit(0))

			outStr := string(crypta.Out.Contents())

			Expect(outStr).To(Not(BeEmpty()))
			Expect(outStr).To(ContainSubstring("Usage:"))
			Expect(crypta.Err.Contents()).To(BeEmpty())
		})
	})

})
