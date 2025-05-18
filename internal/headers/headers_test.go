package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeaders(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("HoSt: localhost:42069\r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069", headers["host"])
	assert.Equal(t, 23, n)
	assert.False(t, done)

	// Test: Valid with extra whitespace
	headers = NewHeaders()
	data = []byte("          HosT: localhost:42069                    \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069", headers["host"])
	assert.Equal(t, 53, n)
	assert.False(t, done)

	// Test: Valid done
	headers = NewHeaders()
	data = []byte("\r\n my favorite game is starcraft")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Empty(t, headers)
	assert.Equal(t, 2, n)
	assert.True(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("Host : localhost:42069\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Invalid characters header
	headers = NewHeaders()
	data = []byte("Hst: localhost:42069\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Same header key
	headers = map[string]string{"host": "localhost:8000"}
	data = []byte("Host: localhost:42069\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:8000, localhost:42069", headers["host"])
	assert.Equal(t, 23, n)
	assert.False(t, done)
}
