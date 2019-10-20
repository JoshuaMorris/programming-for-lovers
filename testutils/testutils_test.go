package testutils

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	folder = "testutils"
)

func TestWriteLoadFixture(t *testing.T) {
	var tests = []struct {
		name    string
		data    string
		fixture string
	}{
		{
			"Should write the string and then read it back",
			"Use the force",
			"writeLoadFixture.golden"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteFixture(t, folder, tt.fixture, []byte(tt.data))

			actual := tt.data
			expected := LoadFixture(t, folder, tt.fixture)

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("actual = %s, expected = %s", actual, expected)
			}
		})
	}
}

func TestFixturePath(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)

	fixturePath(t, "temp", "temp.golden")

	assert.DirExists(t, filepath.Join(filepath.Dir(filename), "..", "testdata", "temp"), "Should have created a 'temp' directory but did not")
	defer os.RemoveAll(filepath.Join(filepath.Dir(filename), "..", "testdata", "temp"))
}

func TestBoolsToBytes(t *testing.T) {
	var tests = []struct {
		name    string
		data    []bool
		fixture string
	}{
		{
			name:    "Should create a byte array from a bool array",
			data:    []bool{true, false, true, true},
			fixture: "BoolsToBytesFixture.golden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BoolsToBytes(tt.data)

			if *Update {
				WriteFixture(t, folder, tt.fixture, result)
			}

			actual := result
			expected := []byte(LoadFixture(t, folder, tt.fixture))

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("actual = %v, expected = %v", actual, expected)
			}
		})
	}
}
