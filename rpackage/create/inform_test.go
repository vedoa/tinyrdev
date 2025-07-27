package create

import (
	"bytes"
	"log"
	"testing"
)

func TestInform(t *testing.T) {
	log.Println("Running TestInform")

	var buf bytes.Buffer
	originalOutput := log.Writer() // Save original output
	log.SetOutput(&buf)
	defer log.SetOutput(originalOutput) // Restore after test

	message := "test message"
	err := inform(message)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expected := "writing " + message + " \n"
	if !bytes.Contains(buf.Bytes(), []byte(expected)) {
		t.Errorf("Expected log to contain %q, got %q", expected, buf.String())
	}
}
