package resources

type Todo struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	Owner     uint   `json:"owner"`
}

func GetAllTodos() []Todo {
	atodo := Todo{
		Title:     "first note",
		Body:      "This is the first note for testing purposes.",
		CreatedAt: "2025-10-19",
		Owner:     1,
	}
	var todos []Todo
	todos = append(todos, atodo)
	return todos
}
