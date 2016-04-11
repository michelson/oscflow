package patches

type Patch struct {
	Name        string `json:"name"`
	Location    bool   `json:"location"`
	Description string `json:"description"`
}

type Patches []Patch
