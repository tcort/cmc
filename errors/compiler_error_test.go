package errors

import (
	"testing"
)

func TestNew(t *testing.T) {

	error := New("message", "filename", 42)
	if error.Error() != "ERROR filename:42 message" {
		t.Errorf("Unexpected output")
	}
}
