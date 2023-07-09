package state;

type App struct {
    CurrId int
    Todos []Todo
}

type Todo struct {
    Id int
    Title string
    Completed bool
}

var AppState = &App{
    CurrId: 0,
    Todos: []Todo{
        {Id: 0, Title: "Learn HTMX", Completed: false},
    },
}

func (app *App) Increment() {
    app.CurrId++;
}

func (app *App) AddTodo(title string) {
    app.Todos = append(app.Todos, Todo{Id: app.CurrId, Title: title, Completed: false})
    app.Increment()
}

func (app *App) EditTodo(title string, completed bool, id int) {
    for i, todo := range app.Todos {
        if todo.Id == id {
            app.Todos[i] = Todo{Id: id, Title: title, Completed: completed}
        }
    }
}

func (app *App) DeleteTodo(id int) {
    for i, todo := range app.Todos {
        if todo.Id == id {
            app.Todos = append(app.Todos[:i], app.Todos[i+1:]...)
        }
    }
}
