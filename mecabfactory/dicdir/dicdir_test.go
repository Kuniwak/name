package dicdir

import (
	"testing"
)

func TestByMecabConfig(t *testing.T) {
	d, err := searchPathByMecabConfig()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if d == "" {
		t.Errorf("want non-empty, but empty")
		return
	}
}
