package routes

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"thegoht/pkg/state"
)

func EditTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	r.ParseForm()
	button_val := r.FormValue("button")
	title := r.FormValue("title")
	completed := r.FormValue("completed") == "on"
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if button_val == "delete" {
		state.AppState.DeleteTodo(id)
        tmpt := template.Must(template.ParseFiles("templates/todos.tmpl", "templates/todos/todo.tmpl"))
        if len(state.AppState.Todos) == 0 {
            state.AppState.CurrId = 0
            state.AppState.AddTodo("Relax")
        }
		body := bytes.NewBuffer(nil)
		err = tmpt.Execute(body, state.AppState)
		if err != nil {
            fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(body.Bytes())
		return
	}
	state.AppState.EditTodo(title, completed, id)
	tmpt := template.Must(template.ParseFiles("templates/todos/todo.tmpl"))
	body := bytes.NewBuffer(nil)
	err = tmpt.Execute(body, state.AppState)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(body.Bytes())
}
