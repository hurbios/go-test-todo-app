package resources

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	Owner     int    `json:"owner"`
}

func GetAllTodos() []Todo {
	var todos []Todo
	todos = append(todos, Todo{
		Id:        1,
		Title:     "first note",
		Body:      "This is the first note for testing purposes.",
		CreatedAt: "2025-10-19",
		Owner:     1,
	})
	todos = append(todos, Todo{
		Id:        2,
		Title:     "second note",
		Body:      "Just another testnote...",
		CreatedAt: "2025-10-19",
		Owner:     1,
	})
	return todos
}
