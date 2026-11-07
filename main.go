package main

import "fmt"

type Task struct {
	Title string
	Done  bool
}

type TaskList struct {
	Tasks []Task
}

func (list *TaskList) Add(title string) {
	list.Tasks = append(list.Tasks, Task{Title: title})
}

func (list *TaskList) Complete(index int) bool {
	if index < 0 || index >= len(list.Tasks) {
		return false
	}
	list.Tasks[index].Done = true
	return true
}

func (list TaskList) Print() {
	for i, task := range list.Tasks {
		status := " "
		if task.Done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Title)
	}
}

func main() {
	var tasks TaskList
	tasks.Add("Познакомиться со структурами")
	tasks.Add("Написать первую программу на Go")
	tasks.Complete(0)
	tasks.Print()
}
