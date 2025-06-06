// Code generated by gowrap. DO NOT EDIT.
// template: ../templates/opentelemetry
// gowrap: http://github.com/hexdigest/gowrap

package templatestests

//go:generate gowrap gen -p github.com/hexdigest/gowrap/templates_tests -i GenericsTestInterface -t ../templates/opentelemetry -o interface_with_opentelemetry.go -v DecoratorName=TestInterfaceWithOpenTelemetry -l ""

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	_codes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// TestInterfaceWithOpenTelemetry implements GenericsTestInterface interface instrumented with open telemetry spans
type TestInterfaceWithOpenTelemetry[T any, U TestInterface] struct {
	GenericsTestInterface[T, U]
	_instance      string
	_spanDecorator func(span trace.Span, params, results map[string]interface{})
}

// NewTestInterfaceWithOpenTelemetry returns TestInterfaceWithOpenTelemetry[T, U]
func NewTestInterfaceWithOpenTelemetry[T any, U TestInterface](base GenericsTestInterface[T, U], instance string, spanDecorator ...func(span trace.Span, params, results map[string]interface{})) TestInterfaceWithOpenTelemetry[T, U] {
	d := TestInterfaceWithOpenTelemetry[T, U]{
		GenericsTestInterface: base,
		_instance:             instance,
	}

	if len(spanDecorator) > 0 && spanDecorator[0] != nil {
		d._spanDecorator = spanDecorator[0]
	}

	return d
}

// ContextNoError implements GenericsTestInterface
func (_d TestInterfaceWithOpenTelemetry[T, U]) ContextNoError(ctx context.Context, a1 string, a2 string) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "GenericsTestInterface.ContextNoError")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"a1":  a1,
				"a2":  a2}, map[string]interface{}{})
		}
		_span.End()
	}()
	_d.GenericsTestInterface.ContextNoError(ctx, a1, a2)
	return
}

// F implements GenericsTestInterface
func (_d TestInterfaceWithOpenTelemetry[T, U]) F(ctx context.Context, a1 T, a2 ...U) (result1 T, result2 string, err error) {
	ctx, _span := otel.Tracer(_d._instance).Start(ctx, "GenericsTestInterface.F")
	defer func() {
		if _d._spanDecorator != nil {
			_d._spanDecorator(_span, map[string]interface{}{
				"ctx": ctx,
				"a1":  a1,
				"a2":  a2}, map[string]interface{}{
				"result1": result1,
				"result2": result2,
				"err":     err})
		} else if err != nil {
			_span.RecordError(err)
			_span.SetStatus(_codes.Error, err.Error())
			_span.SetAttributes(
				attribute.String("event", "error"),
				attribute.String("message", err.Error()),
			)
		}

		_span.End()
	}()
	return _d.GenericsTestInterface.F(ctx, a1, a2...)
}
