package greedy

func Optimize(tasks map[string]string) *map[string]string {
	A := make(map[string]string)
	for len(tasks) > 0 {
		s, e := findMinAndRemove(tasks)
		A[e] = s
		for k, v := range tasks {
			if e > v {
				delete(tasks, k)
			}
		}
	}
	return &A
}

func findMinAndRemove(tasks map[string]string) (s string, e string) {
	e = "24:00"
	for k, _ := range tasks {
		if e > k {
			e = k
		}
	}
	s = tasks[e]
	delete(tasks, e)
	return
}
