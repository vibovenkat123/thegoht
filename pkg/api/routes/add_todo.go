package routes

import (
	"bytes"
	"html/template"
	"net/http"
	"thegoht/pkg/state"
)

func AddTodo(w http.ResponseWriter, r *http.Request) {
    state.AppState.AddTodo("test")
	tmpt := template.Must(template.ParseFiles("templates/todos.tmpl", "templates/todos/todo.tmpl"))
	body := bytes.NewBuffer(nil)
	err := tmpt.Execute(body, state.AppState)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "text/html")
	w.Write(body.Bytes())
}
