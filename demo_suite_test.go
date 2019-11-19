package learning_badges

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"os"
	"testing"
)

func TestAwesome(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter(os.Getenv("CI_REPORT"))
	RunSpecsWithDefaultAndCustomReporters(t, "Awesome Suite", []Reporter{junitReporter})
}
