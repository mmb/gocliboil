package gocliboiltest_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGocliboiltest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gocliboiltest Suite")
}
