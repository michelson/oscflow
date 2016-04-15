package patches

import (
	//"bytes"
	"testing"
)

func TestGetPatchSuccess(t *testing.T) {

	pp := Patches{}

	p, err := pp.GetPatch("hello_world.pd")

	t.Logf("this is the patches %v", p)

	if err != nil {
		t.Fatalf("empty patches dir: %v", p)
	}

}

func TestGetPatchFail(t *testing.T) {

	pp := Patches{}

	p, err := pp.GetPatch("foo.pd")

	t.Logf("this is the patches %v", p)

	if err == nil {
		t.Fatalf("should respond with error here!: %v", p)
	}

}
