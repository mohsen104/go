package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var TodoList = []string{}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		// wg.Add(1)
		go GetTodo(i, &wg)
	}
	wg.Wait()
	fmt.Println(TodoList)
}

func GetTodo(id int, wg *sync.WaitGroup) {
	GetUrl("https://jsonplaceholder.typicode.com/todos/"+strconv.Itoa(id), wg)
}

func GetUrl(url string, wg *sync.WaitGroup) {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	resBody, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		panic(err)
	}

	TodoList = append(TodoList, string(resBody))
	wg.Done()
}

// [Done] exited with code=0 in 20.641 seconds
// [Done] exited with code=0 in 5.347 seconds
// [Done] exited with code=0 in 5.274 seconds
// [Done] exited with code=0 in 4.835 seconds

// ---

// 5879.53ms
// 4798.96ms
// 4833.92ms

// ---

// [Done] exited with code=0 in 1.063 seconds
// [Done] exited with code=0 in 1.017 seconds
