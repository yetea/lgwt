package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Chrissie")
	want := "Hello, Chrissie"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
