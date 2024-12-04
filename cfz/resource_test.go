package cfz_test

import (
	"testing"

	"github.com/ibrt/golang-utils/fixturez"
	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-cloudformation/cfz"
)

type ResourceSuite struct {
	// intentionally empty
}

func TestResourceSuite(t *testing.T) {
	fixturez.RunSuite(t, &ResourceSuite{})
}

func (*ResourceSuite) TestResourceLogicalName(g *WithT) {
	g.Expect(cfz.ResourceLogicalName("MyResource").GetResourceLogicalName()).
		To(Equal("MyResource"))
}
