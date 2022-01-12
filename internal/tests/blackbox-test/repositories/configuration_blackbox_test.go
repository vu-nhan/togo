package blackbox_test

import (
	"context"
	"github.com/manabie-com/togo/internal/repositories"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestConfigurationBlackbox(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Starting Configuration Repository Blackbox TestSuite")
}

var _ = Describe("Configuration Repository Blackbox Test", func() {
	var repository repositories.ConfigurationRepository

	BeforeSuite(func() {
		gdb, err := gorm.Open(sqlite.Open("../../../../data.db"), &gorm.Config{})
		Expect(err).ShouldNot(HaveOccurred())
		repository = repositories.NewConfigurationRepository(gdb)
	})

	BeforeEach(func() {
		println("Starting test")
	})

	AfterEach(func() {
		println("Ending test")
	})

	Describe("Test Configuration Capacity", func() {
		Context("Get Capacity by User ID and Date", func() {
			It("Capacity has found", func() {
				capacity, err := repository.GetCapacity(context.TODO(), "firstUser", "2020-08-22")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(capacity).Should(BeEquivalentTo(2))
			})

			It("Capacity has not found", func() {
				capacity, err := repository.GetCapacity(context.TODO(), "123", "2020-01-01")
				Expect(err).Should(HaveOccurred())
				Expect(capacity).Should(BeZero())
			})
		})
	})
})
