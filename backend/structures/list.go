package structures

type List struct {
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Elements []RandomObject `json:"elements"`
}

type RandomObject struct {
	Name        string            `json:"name"`
	Probability int               `json:"probability"`
	Properties  map[string]string `json:"properties"`
}
