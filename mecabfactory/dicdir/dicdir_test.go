package dicdir

import (
	"os"
	"testing"
)

func TestByMecabConfig(t *testing.T) {
	d := ByMecabConfig()
	dicDir, err := d()
	if err != nil {
		t.Errorf("want nil, got %#v", err)
		return
	}

	if len(dicDir) == 0 {
		t.Errorf("want a non-empty string, got an empty string")
		return
	}

	stat, err := os.Stat(dicDir)
	if err != nil {
		t.Errorf("want nil, got %#v", err)
		return
	}

	if !stat.IsDir() {
		t.Errorf("want a directory, got a not directory: %q", dicDir)
	}
}
