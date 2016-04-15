package patches

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var app_dir = "/Applications/Pd-extended.app/Contents/Resources/bin/pd"

// -nogui -verbose main.pd

func InitMainPatch() {

	cmd := exec.Command(app_dir, "-nogui", "-verbose", "main.pd")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())

}