package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func fetchTodo(id int, wg *sync.WaitGroup, todos chan<- Todo, errs chan<- error) {
	defer wg.Done()
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		errs <- fmt.Errorf("error fetching TODO %d: %v", id, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errs <- fmt.Errorf("error reading response body for TODO %d: %v", id, err)
		return
	}

	var todo Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		errs <- fmt.Errorf("error unmarshalling TODO %d: %v", id, err)
		return
	}

	todos <- todo
}

func main() {
	var wg sync.WaitGroup
	todos := make(chan Todo, 20)
	errs := make(chan error, 20)

	for i := 2; i <= 40; i += 2 {
		wg.Add(1)
		go fetchTodo(i, &wg, todos, errs)
	}

	wg.Wait()
	close(todos)
	close(errs)

	for todo := range todos {
		fmt.Printf("Title: %s, Completed: %v\n", todo.Title, todo.Completed)
	}

	for err := range errs {
		fmt.Printf("Error: %v\n", err)
	}
}
