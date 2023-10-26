package main

import (
	"fmt"
	"sync"
	"time"
)

// 2. Горутины для вывода сообщений
// messages := []string{"Привет", "Здравствуй", "Приветствую"}
// Есть слайс, надо вывести каждое сообщение в отдельной горутине через n секунд,
// где n это индекс сообщения в слайсе. При проверки я могу добавлять в слайс элементы или убирать

func printMsg(s string, t int) {
	

	secs := time.Duration(t)
	time.Sleep(secs * time.Second)
	fmt.Printf("%s was printed after %d seconds\n", s, t)
}

func main() {
	messages := []string{"Привет", "Здравствуй", "Приветствую","a","dasda","dasdasddddddddd"}
	var wg sync.WaitGroup

	wg.Add(len(messages))

	for i, m := range messages {
		go func(i int, m string){
			defer wg.Done()
			printMsg(m, i)
		}(i, m)
	}

	wg.Wait()
}