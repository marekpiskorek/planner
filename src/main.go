package main

import (
    "fmt"
    "errors"
)

type Item struct {
    Title, Text string
}

type Plan struct {
    Title string
    Todo List
    Done List
}

type List struct {
    Name string
    Items []Item
}

func NewPlan(title string) *Plan {
    todo := List{Name: "Todo", Items: make([]Item, 0)}
    done := List{Name: "Done", Items: make([]Item, 0)}
    return &Plan{Title: title, Todo: todo, Done: done}
}

func (p *Plan) Show() {
    fmt.Print(p.Title, "\n", "TODO:\n")
    for i := 0; i < len(p.Todo.Items); i+=1 {
        item := p.Todo.Items[i]
        fmt.Print(i, ". ", item.Title, ": ", item.Text, "\n")
    }
    fmt.Print("\nDONE:\n")
    for i := 0; i < len(p.Done.Items); i+=1 {
        item := p.Done.Items[i]
        fmt.Print(i, ". ", item.Title, ": ", item.Text, "\n")
    }
    fmt.Print("\n")
}

func (p *Plan) AddTodoItem(title, text string) {
    item := Item{Title: title, Text: text}
    p.Todo.Items = append(p.Todo.Items, item)
}

func (p *Plan) RemoveTodoItem(id int) error {
    if id > len(p.Todo.Items) {
        return errors.New("The value of id is out of range") 
    }
    p.Todo.Items = append(p.Todo.Items[:id], p.Todo.Items[id+1:]...)
    return nil
}

func main() {
    // Initially (in the POC) the main function can serve as just an example of usage.
    plan := NewPlan("example")
    plan.AddTodoItem("workout", "Perform 100 push-ups")
    plan.AddTodoItem("play", "Win 2 games in Fifa")
    plan.Show()
    plan.RemoveTodoItem(1)
    plan.AddTodoItem("books", "Read one chapter of book")
    plan.Show()
}
