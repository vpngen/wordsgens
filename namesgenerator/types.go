package namesgenerator

// Person - struct with awardee, male, description and wiki URL.
type Person struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
}
