package main

import "fmt"

var tasks = []string{"Wash dishes", "Go outside"}

func main() {
	tasks = append(tasks, "Learn Go")
	for i, val := range(tasks) {
		fmt.Println(i+1, "-", val)
	}
}