package io

import (
	"bytes"
	"io"
	"testing"

	"code.uber.internal/infra/memtsdb"

	"github.com/stretchr/testify/require"
)

func TestSliceOfSliceReaderNilData(t *testing.T) {
	r := NewSliceOfSliceReader(nil)
	var buf []byte
	n, err := r.Read(buf)
	require.Equal(t, 0, n)
	require.NoError(t, err)

	buf = make([]byte, 1)
	n, err = r.Read(buf)
	require.Equal(t, 0, n)
	require.Equal(t, io.EOF, err)
}

func TestSliceOfSliceReaderOneSlice(t *testing.T) {
	bytes := [][]byte{
		{0x1, 0x2, 0x3},
	}
	r := NewSliceOfSliceReader(bytes)
	buf := make([]byte, 1)
	for i := 0; i < 3; i++ {
		n, err := r.Read(buf)
		require.Equal(t, 1, n)
		require.Equal(t, bytes[0][i:i+1], buf[:n])
		require.NoError(t, err)
	}
	for i := 0; i < 2; i++ {
		n, err := r.Read(buf)
		require.Equal(t, 0, n)
		require.Equal(t, io.EOF, err)
	}

	r = NewSliceOfSliceReader(bytes)
	buf = make([]byte, 2)
	for i := 0; i < 2; i++ {
		n, err := r.Read(buf)
		if i == 0 {
			require.Equal(t, 2, n)
			require.Equal(t, bytes[0][0:2], buf[:n])
		} else {
			require.Equal(t, 1, n)
			require.Equal(t, bytes[0][2:], buf[:n])
		}
		require.NoError(t, err)
	}
	for i := 0; i < 2; i++ {
		n, err := r.Read(buf)
		require.Equal(t, 0, n)
		require.Equal(t, io.EOF, err)
	}
}

func TestSliceOfSliceReaderMultiSlice(t *testing.T) {
	bytes := [][]byte{
		{0x1, 0x2, 0x3},
		{0x4, 0x5},
		{0x6, 0x7, 0x8, 0x9},
	}
	r := NewSliceOfSliceReader(bytes)
	buf := make([]byte, 3)
	expectedResults := []struct {
		n    int
		data []byte
	}{
		{3, []byte{0x1, 0x2, 0x3}},
		{3, []byte{0x4, 0x5, 0x6}},
		{3, []byte{0x7, 0x8, 0x9}},
	}
	for i := 0; i < 3; i++ {
		n, err := r.Read(buf)
		require.Equal(t, expectedResults[i].n, n)
		require.Equal(t, expectedResults[i].data, buf[:n])
		require.NoError(t, err)
	}
	for i := 0; i < 2; i++ {
		n, err := r.Read(buf)
		require.Equal(t, 0, n)
		require.Equal(t, io.EOF, err)
	}
}

func TestReaderSliceReader(t *testing.T) {
	readers := []io.Reader{
		bytes.NewReader([]byte{0x1, 0x2}),
		bytes.NewReader([]byte{0x3, 0x4, 0x5, 0x6}),
		bytes.NewReader([]byte{0x7, 0x8, 0x9}),
	}
	r := NewReaderSliceReader(readers)
	expected := []struct {
		n     int
		bytes []byte
	}{
		{4, []byte{0x1, 0x2, 0x3, 0x4}},
		{4, []byte{0x5, 0x6, 0x7, 0x8}},
		{1, []byte{0x9}},
	}
	var b [4]byte
	for i := 0; i < 3; i++ {
		n, err := r.Read(b[:])
		require.NoError(t, err)
		require.Equal(t, expected[i].n, n)
		require.Equal(t, expected[i].bytes, b[:n])
	}
	n, err := r.Read(b[:])
	require.Equal(t, io.EOF, err)
	require.Equal(t, 0, n)
	require.Equal(t, readers, r.Readers())
}

func TestSegmentReader(t *testing.T) {
	head := []byte{
		0x20, 0xc5, 0x10, 0x55, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x0, 0x0, 0x0, 0x0,
		0xa0, 0x90, 0xf, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x28, 0x3f, 0x2f,
		0xc0, 0x1, 0xf4, 0x1, 0x0, 0x0, 0x0, 0x2, 0x1, 0x2, 0x7, 0x10, 0x1e,
		0x0, 0x1, 0x0, 0xe0, 0x65, 0x58,
	}
	tail := []byte{
		0xcd, 0x3, 0x0, 0x0, 0x0, 0x0,
	}

	expected := []byte{
		0x20, 0xc5, 0x10, 0x55, 0x0, 0x0, 0x0, 0x0, 0x1f, 0x0, 0x0, 0x0, 0x0,
		0xa0, 0x90, 0xf, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x28, 0x3f, 0x2f,
		0xc0, 0x1, 0xf4, 0x1, 0x0, 0x0, 0x0, 0x2, 0x1, 0x2, 0x7, 0x10, 0x1e,
		0x0, 0x1, 0x0, 0xe0, 0x65, 0x58, 0xcd, 0x3, 0x0, 0x0, 0x0, 0x0,
	}

	r := NewSegmentReader(memtsdb.Segment{Head: head, Tail: tail})
	var b [100]byte
	n, err := r.Read(b[:])
	require.NoError(t, err)
	require.Equal(t, len(expected), n)
	require.Equal(t, expected, b[:n])

	n, err = r.Read(b[:])
	require.Equal(t, io.EOF, err)
	require.Equal(t, 0, n)

	require.Equal(t, head, r.Segment().Head)
	require.Equal(t, tail, r.Segment().Tail)
}