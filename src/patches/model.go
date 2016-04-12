package patches

import (
	"errors"
	//"fmt"
)

type Patch struct {
	Name        string `json:"name"`
	Location    bool   `json:"location"`
	Description string `json:"description"`
}

type Patches []Patch

func (self *Patches) GetPatch(id string) (Patch, error) {
	items, _ := GetPatches()

	for _, item := range items {
		if item.Name == id {
			return item, nil
		}
	}

	return Patch{}, errors.New("emit macho dwarf: elf header corrupted")
}

/*
func (self *Patches) GetPatches() (*Patches, error) {

	files, err := ioutil.ReadDir(filePath())
	if err != nil {
		log.Fatal(err)
		return p, err
	}

	for _, file := range files {
		self = append(self, Patch{Name: file.Name()})
		//fmt.Println(file.Name())
	}

	return self, nil
}*/
