package patches

import (
	//"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

func filePath() string {
	_, filename, _, _ := runtime.Caller(1)
	f := path.Join(path.Dir(filename), "../../patches/")
	return f
}

func GetPatches() ([]Patch, error) {
	p := Patches{}

	files, err := ioutil.ReadDir(filePath())
	if err != nil {
		log.Fatal(err)
		return p, err
	}

	for _, file := range files {
		p = append(p, Patch{Name: file.Name()})
		fmt.Println(file.Name())
	}

	return p, nil
}
