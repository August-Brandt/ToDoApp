package datatypes

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Dodate      string `json:"doDate"`
	Finished    bool   `json:"finished"`
}
