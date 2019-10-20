package testutils

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Update is the flag we pass to update the golden files for the test currently running
var Update = flag.Bool("update", false, "update golden files")

func fixturePath(t *testing.T, folder, fixture string) string {
	t.Helper()

	_, filename, _, _ := runtime.Caller(0)

	if _, err := os.Stat(filepath.Join(filepath.Dir(filename), "..", "testdata", folder)); os.IsNotExist(err) {
		err = os.Mkdir(filepath.Join(filepath.Dir(filename), "..", "testdata", folder), 0777)
		assert.Nil(t, err, "Error writing to testdata directory: %v", err)
	}

	return filepath.Join(filepath.Dir(filename), "..", "testdata", folder, fixture)
}

// WriteFixture creates our golden file for testing
func WriteFixture(t *testing.T, folder, fixture string, content []byte) {
	t.Helper()

	err := ioutil.WriteFile(fixturePath(t, folder, fixture), content, 0644)
	assert.Nil(t, err, "Error writing fixture: %v", err)
}

// LoadFixture is for loading the data from the golden file for comparison
func LoadFixture(t *testing.T, folder, fixture string) string {
	t.Helper()

	content, err := ioutil.ReadFile(fixturePath(t, folder, fixture))
	assert.Nil(t, err, "Error writing fixture: %v", err)

	return string(content)
}

// BoolsToBytes converts boolean arrays to byte arrays
func BoolsToBytes(t []bool) []byte {
	b := make([]byte, (len(t)+7)/8)
	for i, x := range t {
		if x {
			b[i/8] |= 0x80 >> uint(i%8)
		}
	}
	return b
}
