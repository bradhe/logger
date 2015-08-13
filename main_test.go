package logger

import (
	"bytes"
	"strings"
	"sync"
	"testing"
)

func assertMissingString(t *testing.T, buf *bytes.Buffer, needle string) {
	if strings.Index(buf.String(), needle) != -1 {
		t.Fatalf("found string [%s]", needle)
	}
}

func assertContainsString(t *testing.T, buf *bytes.Buffer, needle string) {
	if strings.Index(buf.String(), needle) == -1 {
		t.Fatalf("found string [%s]", needle)
	}
}

func TestDefaultSetup(t *testing.T) {
	// Always reset the setup sync.Once so that we can re-run setup
	defer func() {
		setup = sync.Once{}
	}()

	buf := bytes.NewBuffer([]byte{})

	Setup(buf)

	Info.Println("Info")
	Debug.Println("Debug")

	assertMissingString(t, buf, "Debug")
	assertContainsString(t, buf, "Info")
}

func TestSetupWithLogLevel(t *testing.T) {
	// Always reset the setup sync.Once so that we can re-run setup
	defer func() {
		setup = sync.Once{}
	}()

	buf := bytes.NewBuffer([]byte{})

	// Super aggressive!
	SetupWith(LogLevelError, buf)

	// All these should end up empty.
	Warning.Println("Warning")
	Info.Println("Info")
	Debug.Println("Debug")

	assertMissingString(t, buf, "Warning")
	assertMissingString(t, buf, "Info")
	assertMissingString(t, buf, "Debug")
}

func TestPrefixes(t *testing.T) {
	t.SkipNow()
}

func BenchmarkPrefixes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Push("TEST")
		Pop()
	}
}
