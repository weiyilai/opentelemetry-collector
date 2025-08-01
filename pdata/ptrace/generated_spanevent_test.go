// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package ptrace

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/pdata/internal"
	otlptrace "go.opentelemetry.io/collector/pdata/internal/data/protogen/trace/v1"
	"go.opentelemetry.io/collector/pdata/internal/json"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestSpanEvent_MoveTo(t *testing.T) {
	ms := generateTestSpanEvent()
	dest := NewSpanEvent()
	ms.MoveTo(dest)
	assert.Equal(t, NewSpanEvent(), ms)
	assert.Equal(t, generateTestSpanEvent(), dest)
	dest.MoveTo(dest)
	assert.Equal(t, generateTestSpanEvent(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newSpanEvent(&otlptrace.Span_Event{}, &sharedState)) })
	assert.Panics(t, func() { newSpanEvent(&otlptrace.Span_Event{}, &sharedState).MoveTo(dest) })
}

func TestSpanEvent_CopyTo(t *testing.T) {
	ms := NewSpanEvent()
	orig := NewSpanEvent()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestSpanEvent()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newSpanEvent(&otlptrace.Span_Event{}, &sharedState)) })
}

func TestSpanEvent_MarshalAndUnmarshalJSON(t *testing.T) {
	stream := json.BorrowStream(nil)
	defer json.ReturnStream(stream)
	src := generateTestSpanEvent()
	src.marshalJSONStream(stream)
	require.NoError(t, stream.Error())

	// Append an unknown field at the start to ensure unknown fields are skipped
	// and the unmarshal logic continues.
	buf := stream.Buffer()
	assert.EqualValues(t, '{', buf[0])
	iter := json.BorrowIterator(append([]byte(`{"unknown": "string",`), buf[1:]...))
	defer json.ReturnIterator(iter)
	dest := NewSpanEvent()
	dest.unmarshalJSONIter(iter)
	require.NoError(t, iter.Error())

	assert.Equal(t, src, dest)
}

func TestSpanEvent_Timestamp(t *testing.T) {
	ms := NewSpanEvent()
	assert.Equal(t, pcommon.Timestamp(0), ms.Timestamp())
	testValTimestamp := pcommon.Timestamp(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.Equal(t, testValTimestamp, ms.Timestamp())
}

func TestSpanEvent_Name(t *testing.T) {
	ms := NewSpanEvent()
	assert.Empty(t, ms.Name())
	ms.SetName("test_name")
	assert.Equal(t, "test_name", ms.Name())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newSpanEvent(&otlptrace.Span_Event{}, &sharedState).SetName("test_name") })
}

func TestSpanEvent_Attributes(t *testing.T) {
	ms := NewSpanEvent()
	assert.Equal(t, pcommon.NewMap(), ms.Attributes())
	internal.FillTestMap(internal.Map(ms.Attributes()))
	assert.Equal(t, pcommon.Map(internal.GenerateTestMap()), ms.Attributes())
}

func TestSpanEvent_DroppedAttributesCount(t *testing.T) {
	ms := NewSpanEvent()
	assert.Equal(t, uint32(0), ms.DroppedAttributesCount())
	ms.SetDroppedAttributesCount(uint32(13))
	assert.Equal(t, uint32(13), ms.DroppedAttributesCount())
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { newSpanEvent(&otlptrace.Span_Event{}, &sharedState).SetDroppedAttributesCount(uint32(13)) })
}

func generateTestSpanEvent() SpanEvent {
	tv := NewSpanEvent()
	fillTestSpanEvent(tv)
	return tv
}

func fillTestSpanEvent(tv SpanEvent) {
	tv.orig.TimeUnixNano = 1234567890
	tv.orig.Name = "test_name"
	internal.FillTestMap(internal.NewMap(&tv.orig.Attributes, tv.state))
	tv.orig.DroppedAttributesCount = uint32(13)
}
