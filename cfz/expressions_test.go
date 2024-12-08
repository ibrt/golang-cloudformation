package cfz_test

import (
	"encoding/json"
	"testing"

	"github.com/ibrt/golang-utils/fixturez"
	"github.com/ibrt/golang-utils/memz"
	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-cloudformation/cfz"
	. "github.com/ibrt/golang-cloudformation/cfz/gen/aws_accessanalyzer"
)

type ExpressionsSuite struct {
	// intentionally empty
}

func TestExpressionsSuite(t *testing.T) {
	fixturez.RunSuite(t, &ExpressionsSuite{})
}

type testOmitEmptyHolder[T any] struct {
	V T `json:"v,omitempty"`
}

type testHolder[T any] struct {
	V T `json:"v"`
}

func (*ExpressionsSuite) TestSomething(g *WithT) {
	New__AWS_AccessAnalyzer_Analyzer("ASd").
		SetV__AnalyzerConfiguration(
			New__AWS_AccessAnalyzer_Analyzer_AnalyzerConfiguration().
				Set__UnusedAccessConfiguration(nil),
		).
		SetV__AnalyzerConfiguration(AWS_AccessAnalyzer_Analyzer_AnalyzerConfiguration{
			UnusedAccessConfiguration: nil,
		})

	New__AWS_AccessAnalyzer_Analyzer_AnalysisRule().SetS__Exclusions()
}

func (*ExpressionsSuite) TestValueExpression(g *WithT) {
	g.Expect(func() { cfz.V("s").Expression("") }).ToNot(Panic())

	{
		v := cfz.V("s")
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`"s"`)))
	}
	{
		v := cfz.V(1)
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`1`)))
	}
	{
		v := cfz.V(1.1)
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`1.1`)))
	}
	{
		v := cfz.V(true)
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`true`)))
	}
	{
		v := cfz.V[*string](nil)
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`null`)))
	}
	{
		v := cfz.V(memz.Ptr(""))
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`""`)))
	}
	{
		v := cfz.V(memz.Ptr("s"))
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`"s"`)))
	}
	{
		v := cfz.V(json.RawMessage("{}"))
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{}`)))
	}
	{
		v := cfz.V(json.RawMessage("null"))
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`null`)))
	}
	{
		v := cfz.V[json.RawMessage](nil)
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`null`)))
	}
	{
		v := cfz.V(testHolder[string]{V: "s"})
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":"s"}`)))
	}
	{
		v := cfz.V(testHolder[string]{V: ""})
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":""}`)))
	}
	{
		v := cfz.V(testOmitEmptyHolder[string]{V: ""})
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{}`)))
	}
	{
		v := testHolder[cfz.Expression[string]]{V: nil}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":null}`)))
	}
	{
		v := testHolder[cfz.Expression[string]]{V: cfz.V("")}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":""}`)))
	}
	{
		v := testHolder[cfz.Expression[string]]{V: cfz.V("s")}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":"s"}`)))
	}
	{
		v := testOmitEmptyHolder[cfz.Expression[string]]{V: nil}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{}`)))
	}
	{
		v := testOmitEmptyHolder[cfz.Expression[string]]{V: cfz.V("")}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":""}`)))
	}
	{
		v := testOmitEmptyHolder[cfz.Expression[string]]{V: cfz.V("s")}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":"s"}`)))
	}
}

func (*ExpressionsSuite) TestNullExpression(g *WithT) {
	g.Expect(func() { cfz.N[string]().Expression("") }).ToNot(Panic())

	{
		v := cfz.N[string]()
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`null`)))
	}
	{
		v := testHolder[cfz.Expression[string]]{V: cfz.N[string]()}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":null}`)))
	}
	{
		v := testOmitEmptyHolder[cfz.Expression[string]]{V: cfz.N[string]()}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":null}`)))
	}
}

func (*ExpressionsSuite) TestGetAttExpression(g *WithT) {
	g.Expect(func() { cfz.GetAtt[string](cfz.V(""), cfz.V("")).Expression("") }).ToNot(Panic())

	{
		v := cfz.GetAtt[string](cfz.V("A"), cfz.V("B"))
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"Fn::GetAtt":["A","B"]}`)))
	}
	{
		v := cfz.GetAtt[string](cfz.V("A"), cfz.N[string]())
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"Fn::GetAtt":["A",null]}`)))
	}
	{
		v := testHolder[cfz.Expression[string]]{V: cfz.GetAtt[string](cfz.V("A"), cfz.V("B"))}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":{"Fn::GetAtt":["A","B"]}}`)))
	}
}

func (*ExpressionsSuite) TestJoinExpression(g *WithT) {
	g.Expect(func() { cfz.Join(":", cfz.SV("A")).Expression("") }).ToNot(Panic())

	{
		v := cfz.Join(":", cfz.S[string]())
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"Fn::Join":[":",[]]}`)))
	}
	{
		v := cfz.Join(":", cfz.SV("A", "B"))
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"Fn::Join":[":",["A","B"]]}`)))
	}
	{
		v := cfz.Join(":", cfz.S(cfz.V("A"), cfz.N[string]()))
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"Fn::Join":[":",["A",null]]}`)))
	}
	{
		v := testHolder[cfz.Expression[string]]{V: cfz.Join(":", cfz.S(cfz.V("A"), cfz.N[string]()))}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":{"Fn::Join":[":",["A",null]]}}`)))
	}
}

func (*ExpressionsSuite) TestSubExpression(g *WithT) {
	g.Expect(func() { cfz.Sub("s").Expression("") }).ToNot(Panic())

	{
		v := cfz.Sub("")
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"Fn::Sub":""}`)))
	}
	{
		v := cfz.Sub("s")
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"Fn::Sub":"s"}`)))
	}
	{
		v := testHolder[cfz.Expression[string]]{V: cfz.Sub("s")}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":{"Fn::Sub":"s"}}`)))
	}
}

func (*ExpressionsSuite) TestRefExpression(g *WithT) {
	g.Expect(func() { cfz.Ref(cfz.V("s")).Expression("") }).ToNot(Panic())

	{
		v := cfz.Ref(cfz.V(""))
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"Ref":""}`)))
	}
	{
		v := cfz.Ref(cfz.V("s"))
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"Ref":"s"}`)))
	}
	{
		v := testHolder[cfz.Expression[string]]{V: cfz.Ref(cfz.V("s"))}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":{"Ref":"s"}}`)))
	}
}

func (*ExpressionsSuite) TestSlice(g *WithT) {
	g.Expect(func() { cfz.S[string]().ExpressionSlice("") }).ToNot(Panic())

	{
		v := cfz.Slice[string](nil)
		g.Expect(v.GetShallowCopy()).To(Equal(v))
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`null`)))
	}
	{
		v := cfz.S[string]()
		g.Expect(v.GetShallowCopy()).To(Equal(v))
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`[]`)))
	}
	{
		v := cfz.S(cfz.V("1"), cfz.Ref(cfz.V("2")))
		g.Expect(v.GetShallowCopy()).To(Equal(v))
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`["1",{"Ref":"2"}]`)))
	}
	{
		v := testHolder[cfz.ExpressionSlice[string]]{V: cfz.S(cfz.V("1"), cfz.V("2"))}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":["1","2"]}`)))
	}

	{
		v := cfz.SV("1", "2")
		g.Expect(v.GetShallowCopy()).To(Equal(v))
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`["1","2"]`)))
	}
}

func (*ExpressionsSuite) TestGetAttExpressionSlice(g *WithT) {
	g.Expect(func() { cfz.GetAttSlice[string](cfz.V(""), cfz.V("")).ExpressionSlice("") }).ToNot(Panic())

	{
		v := cfz.GetAttSlice[string](cfz.V("A"), cfz.V("B"))
		g.Expect(v.GetShallowCopy()).To(Equal(v))
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"Fn::GetAtt":["A","B"]}`)))
	}
	{
		v := cfz.GetAttSlice[string](cfz.V("A"), cfz.N[string]())
		g.Expect(v.GetShallowCopy()).To(Equal(v))
		buf, err := v.MarshalJSON()
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"Fn::GetAtt":["A",null]}`)))
	}
	{
		v := testHolder[cfz.ExpressionSlice[string]]{V: cfz.GetAttSlice[string](cfz.V("A"), cfz.V("B"))}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":{"Fn::GetAtt":["A","B"]}}`)))
	}
}

func (*ExpressionsSuite) TestMap(g *WithT) {
	g.Expect(func() { cfz.M[string]().ExpressionMap("") }).ToNot(Panic())

	{
		v := cfz.Map[string](nil)
		g.Expect(v.GetShallowCopy()).To(Equal(v))
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`null`)))
	}
	{
		v := cfz.M[string]()
		g.Expect(v.GetShallowCopy()).To(Equal(v))
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{}`)))
	}
	{
		v := cfz.M(map[string]cfz.Expression[string]{
			"k1": cfz.V("v1"),
			"k2": cfz.Ref(cfz.V("v2")),
		})
		g.Expect(v.GetShallowCopy()).To(Equal(v))
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"k1":"v1","k2":{"Ref":"v2"}}`)))
	}
	{
		v := testHolder[cfz.ExpressionMap[string]]{
			V: cfz.M(map[string]cfz.Expression[string]{
				"k1": cfz.V("v1"),
				"k2": cfz.V("v2"),
			})}
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"v":{"k1":"v1","k2":"v2"}}`)))
	}

	{
		v := cfz.MV(map[string]string{"k1": "v1", "k2": "v2"})
		g.Expect(v.GetShallowCopy()).To(Equal(v))
		buf, err := json.Marshal(v)
		g.Expect(err).To(Succeed())
		g.Expect(buf).To(Equal([]byte(`{"k1":"v1","k2":"v2"}`)))
	}
}
