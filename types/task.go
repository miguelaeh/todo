package types

import (
  "fmt"
)

type Task struct {
	Title    string
  Priority int
	Alarm    string
}

type Tasks []Task

func (t Tasks) Len() int {
	return len(t)
}

func (t Tasks) Less(i, j int) bool {
	return t[i].Priority < t[j].Priority
}

func (t Tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Tasks) Print() {
    for i := 0; i < t.Len(); i++ {
      fmt.Printf("%d. %s\n", t[i].Priority, t[i].Title)
    }
}
