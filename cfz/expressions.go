package cfz

import (
	"encoding/json"

	"github.com/ibrt/golang-utils/errorz"
	"github.com/ibrt/golang-utils/memz"
)

// Expression describes a template expression.
type Expression[T any] interface {
	Expression(T)
	json.Marshaler
}

// ExpressionSlice describes a slice of Expression[T].
type ExpressionSlice[T any] interface {
	ExpressionSlice(T)
	GetShallowCopy() ExpressionSlice[T]
	json.Marshaler
}

// ExpressionMap describes a map of (string, Expression[T]).
type ExpressionMap[T any] interface {
	ExpressionMap(T)
	GetShallowCopy() ExpressionMap[T]
	json.Marshaler
}

// -------

type valueExpression[T any] struct {
	value T
}

// V returns an Expression[T] that resolves to a constant value.
func V[T any](value T) Expression[T] {
	return &valueExpression[T]{
		value: value,
	}
}

// Expression implements the Expression[T] interface.
func (*valueExpression[T]) Expression(_ T) {
	// intentionally empty
}

// MarshalJSON implements the Expression[T] interface.
func (e *valueExpression[T]) MarshalJSON() ([]byte, error) {
	errorz.Assertf(e != nil, "MarshalJSON is not called on nil pointers.")
	return json.Marshal(e.value)
}

// -------

type nullExpression[T any] struct {
	// intentionally empty
}

// N returns an Expression[T] that explicitly resolves to "null" instead of omitting the field when marshaled to JSON.
func N[T any]() Expression[T] {
	return &nullExpression[T]{}
}

// Expression implements the Expression[T] interface.
func (*nullExpression[T]) Expression(t T) {
	// intentionally empty
}

// MarshalJSON implements the Expression[T] interface.
func (e *nullExpression[T]) MarshalJSON() ([]byte, error) {
	errorz.Assertf(e != nil, "MarshalJSON is not called on nil pointers.")
	return []byte("null"), nil
}

// -------

type getAttExpression[T any] struct {
	logicalNameOfResource Expression[string]
	attributeName         Expression[string]
}

// GetAtt returns an Expression[T] tha resolves to the Fn::GetAtt intrinsic function.
// It works for attributes that return a single value.
// Note that if the "AWS::LanguageExtensions" transform is not enabled, the parameters can only be a constant value.
func GetAtt[T any](logicalNameOfResource, attributeName Expression[string]) Expression[T] {
	return &getAttExpression[T]{
		logicalNameOfResource: logicalNameOfResource,
		attributeName:         attributeName,
	}
}

// Expression implements the Expression[T] interface.
func (*getAttExpression[T]) Expression(_ T) {
	// intentionally empty
}

// MarshalJSON implements the Expression[T] interface.
func (e *getAttExpression[T]) MarshalJSON() ([]byte, error) {
	errorz.Assertf(e != nil, "MarshalJSON is not called on nil pointers.")

	return json.Marshal(map[string]any{
		"Fn::GetAtt": []any{
			e.logicalNameOfResource, e.attributeName,
		},
	})
}

// -------

type joinExpression struct {
	delimiter string
	values    ExpressionSlice[string]
}

// Join returns an Expression[string] that resolves to the Fn::Join intrinsic function.
func Join(delimiter string, values ExpressionSlice[string]) Expression[string] {
	return &joinExpression{
		delimiter: delimiter,
		values:    values.GetShallowCopy(),
	}
}

// Expression implements the Expression[string] interface.
func (*joinExpression) Expression(_ string) {
	// intentionally empty
}

// MarshalJSON implements the Expression[string] interface.
func (e *joinExpression) MarshalJSON() ([]byte, error) {
	errorz.Assertf(e != nil, "MarshalJSON is not called on nil pointers.")

	return json.Marshal(map[string]any{
		"Fn::Join": []any{
			e.delimiter,
			e.values,
		},
	})
}

// -------

type subExpression struct {
	str string
}

// Sub returns an Expression[string] that resolves to the Fn::Sub intrinsic function (without a key-value map).
func Sub(str string) Expression[string] {
	return &subExpression{
		str: str,
	}
}

// Expression implements the Expression[string] interface.
func (*subExpression) Expression(_ string) {
	// intentionally empty
}

// MarshalJSON implements the Expression[string] interface.
func (e *subExpression) MarshalJSON() ([]byte, error) {
	errorz.Assertf(e != nil, "MarshalJSON is not called on nil pointers.")

	return json.Marshal(map[string]any{
		"Fn::Sub": e.str,
	})
}

// -------

type refExpression struct {
	logicalName Expression[string]
}

// Ref is an Expression[string] that resolves to the Ref intrinsic function.
// Note that if the "AWS::LanguageExtensions" transform is not enabled, "logicalName" can only be a constant value.
func Ref(logicalName Expression[string]) Expression[string] {
	return &refExpression{
		logicalName: logicalName,
	}
}

// Expression implements the Expression[string] interface.
func (*refExpression) Expression(_ string) {
	// intentionally empty
}

// MarshalJSON implements the Expression[string] interface.
func (e *refExpression) MarshalJSON() ([]byte, error) {
	errorz.Assertf(e != nil, "MarshalJSON is not called on nil pointers.")

	return json.Marshal(map[string]any{
		"Ref": e.logicalName,
	})
}

// -------

var (
	_ ExpressionSlice[any] = (Slice[any])(nil)
)

// Slice is a slice type that implements ExpressionSlice[T].
type Slice[T any] []Expression[T]

// S returns a Slice[T] with the given elements.
func S[T any](items ...Expression[T]) Slice[T] {
	return memz.ConcatSlices(items)
}

// SV returns a Slice[T] with the given constant values as elements.
func SV[T any](items ...T) Slice[T] {
	return S(memz.TransformSlice(
		items,
		func(i T) Expression[T] {
			return V(i)
		})...)
}

// STV returns a Slice[Tag] with the given elements.
func STV(tagsMap map[string]string) Slice[Tag] {
	tagsSlice := make(Slice[Tag], 0, len(tagsMap))

	for k, v := range tagsMap {
		tagsSlice = append(tagsSlice, V(Tag{
			Key:   V(k),
			Value: V(v),
		}))
	}

	return tagsSlice
}

// ExpressionSlice implements the ExpressionSlice[T] interface.
func (Slice[T]) ExpressionSlice(_ T) {
	// intentionally empty
}

// GetShallowCopy implements the ExpressionSlice[T] interface.
func (e Slice[T]) GetShallowCopy() ExpressionSlice[T] {
	return Slice[T](memz.ShallowCopySlice(e))
}

// MarshalJSON implements the ExpressionSlice[T] interface.
func (e Slice[T]) MarshalJSON() ([]byte, error) {
	if e == nil {
		return []byte("null"), nil
	}

	return json.Marshal([]Expression[T](e))
}

// -------

type getAttExpressionSlice[T any] struct {
	logicalNameOfResource Expression[string]
	attributeName         Expression[string]
}

// GetAttSlice returns an ExpressionSlice[T] that resolves to the Fn::GetAtt intrinsic function.
// It works for attributes that return a List of values.
// Note that if the "AWS::LanguageExtensions" transform is not enabled, the parameters can only be a constant value.
func GetAttSlice[T any](logicalNameOfResource, attributeName Expression[string]) ExpressionSlice[T] {
	return &getAttExpressionSlice[T]{
		logicalNameOfResource: logicalNameOfResource,
		attributeName:         attributeName,
	}
}

// ExpressionSlice implements the ExpressionSlice[T] interface.
func (*getAttExpressionSlice[T]) ExpressionSlice(_ T) {
	// intentionally empty
}

// GetShallowCopy implements the ExpressionSlice[T] interface.
func (e *getAttExpressionSlice[T]) GetShallowCopy() ExpressionSlice[T] {
	return &getAttExpressionSlice[T]{
		logicalNameOfResource: e.logicalNameOfResource,
		attributeName:         e.attributeName,
	}
}

// MarshalJSON implements the ExpressionSlice[T] interface.
func (e *getAttExpressionSlice[T]) MarshalJSON() ([]byte, error) {
	errorz.Assertf(e != nil, "MarshalJSON is not called on nil pointers.")

	return json.Marshal(map[string]any{
		"Fn::GetAtt": []any{
			e.logicalNameOfResource, e.attributeName,
		},
	})
}

// -------

var (
	_ ExpressionMap[any] = (Map[any])(nil)
)

// Map is a map type that implements ExpressionMap[T].
type Map[T any] map[string]Expression[T]

// M returns a Map[T] with the given elements merged into one map.
func M[T any](maps ...map[string]Expression[T]) Map[T] {
	return memz.MergeMaps(maps...)
}

// MV returns a Map[T] with the given constant values merged into one map.
func MV[T any](maps ...map[string]T) Map[T] {
	return M(
		memz.TransformSlice(
			maps,
			func(m map[string]T) map[string]Expression[T] {
				return memz.TransformMapValues(
					m,
					func(_ string, i T) Expression[T] {
						return V(i)
					})
			})...)
}

// ExpressionMap implements the ExpressionMap[T] interface.
func (Map[T]) ExpressionMap(_ T) {
	// intentionally empty
}

// GetShallowCopy implements the ExpressionMap[T] interface.
func (e Map[T]) GetShallowCopy() ExpressionMap[T] {
	return Map[T](memz.ShallowCopyMap(e))
}

// MarshalJSON implements the ExpressionMap[T] interface.
func (e Map[T]) MarshalJSON() ([]byte, error) {
	if e == nil {
		return []byte("null"), nil
	}

	return json.Marshal(map[string]Expression[T](e))
}
