package patches

import (
	//"bytes"
	"testing"
)

func TestGetPatches(t *testing.T) {

	p, err := GetPatches()

	//t.Logf("this is the patches %v", p)

	if len(p) <= 0 {
		t.Fatalf("empty patches dir: %d", len(p))
	}

	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}

}
