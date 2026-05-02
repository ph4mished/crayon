package termcolor

import (
	"testing"
)

func TestCapability(t *testing.T) {
	caps := Capability()
	//Should be ColorNone, Color8, Color16, Color256, ColorTrue
	t.Logf("Capability: %d", caps)
}

func TestCapability_ColorSupport(t *testing.T) {
	caps := Capability()
	if caps == ColorNone {
		t.Log("Terminal doesn't support color")
	}
}
