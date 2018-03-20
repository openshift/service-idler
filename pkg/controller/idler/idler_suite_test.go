package idler_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestIdlerController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Idler Controller Suite")
}
