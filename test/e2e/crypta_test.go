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
		setTimeout := func(c *exec.Cmd) {
			c.Env = os.Environ()
			c.Env = append(c.Env, "CRYPTA_TIMEOUT=0.1")
		}

		checkErrOutputForConnectionTimeout := func(out string) {
			Expect(out).To(ContainSubstring("Using set maximum connection timeout: 100ms"))
		}

		checkErrOutputForRetries := func(out string) {
			for i := 1; i <= 5; i++ {
				Expect(out).To(ContainSubstring(fmt.Sprintf("Daemon currently not reachable. Retry %d of 5...", i)))
			}
		}

		By("trying to retrieve a value", func() {
			cmd := exec.Command(pathToCrypta, "get", "abcd")
			setTimeout(cmd)
			crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(crypta).WithTimeout(3 * time.Second).Should(gexec.Exit(1))

			errOutput := string(crypta.Err.Contents())
			checkErrOutputForConnectionTimeout(errOutput)
			checkErrOutputForRetries(errOutput)
		})
		By("trying to set a value", func() {
			cmd := exec.Command(pathToCrypta, "set", "abcd", "xyz")
			setTimeout(cmd)
			crypta, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(crypta).WithTimeout(3 * time.Second).Should(gexec.Exit(1))
			errOutput := string(crypta.Err.Contents())
			checkErrOutputForConnectionTimeout(errOutput)
			checkErrOutputForRetries(errOutput)
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
})
