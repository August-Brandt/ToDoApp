package datatypes

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Dodate      string `json:"doDate"`
	Finished    bool   `json:"finished"`
}