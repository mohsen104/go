package main

import (
	"fmt"
	"sync"
	"time"
)

var userList = []int{}
var ready = false

func main() {
	condition := sync.Cond{L: &sync.Mutex{}}

	for i := 0; i < 1_00; i++ {
		go NewRequest(i, &condition)
	}

	time.Sleep(time.Second * 5)
}

func NewRequest(userId int, condition *sync.Cond) {
	Checking(userId, condition)
	condition.L.Lock()
	defer condition.L.Unlock()
	if !ready {
		condition.Wait()
	}
	fmt.Println("UserId", userId, "start streaming")
}

func Checking(userId int, condition *sync.Cond) {
	fmt.Println("UserId", userId, "waiting for start streaming")
	time.Sleep(time.Millisecond * 300)
	condition.L.Lock()
	defer condition.L.Unlock()
	userList = append(userList, userId)
	if len(userList) == 55 {
		ready = true
		condition.Broadcast()
		fmt.Println("UserId", userId, "waiting end")
	}
}
