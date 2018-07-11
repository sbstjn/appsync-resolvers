package resolvers

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAppsyncResolvers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "appsync-resolvers")
}
