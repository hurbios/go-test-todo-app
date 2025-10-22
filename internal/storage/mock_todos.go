package storage

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"createdAt"`
	Owner     int    `json:"owner"`
	Category  string `json:"category"`
}

func GetAllTodos() []Todo {
	var todos []Todo
	todos = append(todos, Todo{
		Id:        1,
		Title:     "night note",
		Body:      "This is the first note for testing purposes.",
		CreatedAt: "2025-10-19",
		Owner:     1,
		Category:  "night",
	})
	todos = append(todos, Todo{
		Id:        2,
		Title:     "day note",
		Body:      "Just another testnote...",
		CreatedAt: "2025-10-19",
		Owner:     1,
		Category:  "day",
	})
	return todos
}
