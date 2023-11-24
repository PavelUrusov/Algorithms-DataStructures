package main

import (
	greedy "algos/greedyalgs"
	"fmt"
)

func main() {

	tasks := make(map[string]string)
	tasks["07:00"] = "05:00"
	tasks["05:15"] = "05:00"
	tasks["05:30"] = "05:00"
	tasks["14:00"] = "08:00"
	tasks["13:00"] = "06:00"
	tasks["20:30"] = "18:00"

	res := greedy.Optimize(tasks)

	for k, v := range *res {
		fmt.Println(v, k)
	}
}
