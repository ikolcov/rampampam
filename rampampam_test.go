package rampampam

import "testing"

func TestRampampam(t *testing.T) {
    want := "Hello, world."
    if got := Rampampam(); got != want {
        t.Errorf("Rampampam() = %q, want %q", got, want)
    }
}
