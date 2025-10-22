package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test-todo-app/internal/storage"

	"github.com/gorilla/mux"
)

func AllTodos(w http.ResponseWriter, r *http.Request) {
	cat := r.URL.Query().Get("cat")
	alltodos := storage.GetAllTodos()

	if cat != "" {
		var alltodosincategory []storage.Todo
		for i := 0; i < len(alltodos); i++ {
			if alltodos[i].Category == cat {
				alltodosincategory = append(alltodosincategory, alltodos[i])
			}
		}
		json.NewEncoder(w).Encode(alltodosincategory)
	} else {
		json.NewEncoder(w).Encode(alltodos)
	}
}

func OneTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Bad id", http.StatusBadRequest)
		return
	}
	alltodos := storage.GetAllTodos()

	for i := 0; i < len(alltodos); i++ {
		if alltodos[i].Id == id {
			json.NewEncoder(w).Encode(alltodos[i])
			return
		}
	}
	http.Error(w, "Id not found", http.StatusNotFound)
}
