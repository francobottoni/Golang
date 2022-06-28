package main

import "fmt"

type TaskList struct {
	Task []*Task
}

type Task struct {
	name       string
	difficulty int
}

func (t *TaskList) deleteTask(i int) {
	t.Task = append(t.Task[:i], t.Task[i+1:]...)
}

func (t *TaskList) addTask(task *Task) {
	t.Task = append(t.Task, task)
}

func (t TaskList) print() {
	for _, task := range t.Task {
		fmt.Println(task)
	}
}

func main() {

	t1 := Task{
		name:       "create a dashboard",
		difficulty: 3,
	}

	t2 := Task{
		name:       "do ABM",
		difficulty: 7,
	}

	t3 := Task{
		name:       "delete value in db",
		difficulty: 2,
	}

	list := TaskList{}
	list.addTask(&t1)
	list.addTask(&t2)
	list.addTask(&t3)

	list.print()
	fmt.Println("----------------------------------")

	for i, task := range list.Task {
		if task.name == "do ABM" {
			list.deleteTask(i)
		}
	}

	list.print()

}
