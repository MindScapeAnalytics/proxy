package entity

type Temperament struct {
	Data        int
	MinusSum    int
	PlusSum     int
	Accentuated bool
}

type TestRawResult struct {
	Raw map[string]Temperament
}

type TestTemplate struct {
	Questions []string `json:"questions"`
}

type Key struct {
	Type     string `json:"type"`
	Multiply int    `json:"multiply"`
	Plus     []int  `json:"plus"`
	Minus    []int  `json:"minus"`
}
