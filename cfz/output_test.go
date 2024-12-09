package cfz_test

import (
	"testing"

	"github.com/ibrt/golang-utils/fixturez"
	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-cloudformation/cfz"
)

type OutputSuite struct {
	// intentionalyl empty
}

func TestOutputSuite(t *testing.T) {
	fixturez.RunSuite(t, &OutputSuite{})
}

func (*OutputSuite) TestSomething(g *WithT) {
	buf, err := cfz.NewOutput("ASD", cfz.Ref(cfz.V("Banana"))).MarshalJSON()
	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(string(buf)).To(Equal(""))
}
