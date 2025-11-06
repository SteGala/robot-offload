package task

import "fmt"

type Task struct {
	ID            int
	Computation   int
	SourceRobotID string
	HostRobotID   string
}

type TaskSet struct {
	Tasks []Task
}

func NewTask(id int, computation int, sourceRobotID string, hostRobotID string) Task {
	return Task{
		ID:            id,
		Computation:   computation,
		SourceRobotID: sourceRobotID,
		HostRobotID:   hostRobotID,
	}
}

func NewTaskSet(n_tasks int) TaskSet {
	tasks := make([]Task, n_tasks)
	for i := 0; i < n_tasks; i++ {
		tasks[i] = NewTask(i, 1, fmt.Sprintf("robot-%d", i), fmt.Sprintf("robot-%d", i))
	}
	return TaskSet{Tasks: tasks}
}

func (ts *TaskSet) Print() {
	for _, task := range ts.Tasks {
		fmt.Printf("Task %d: Computation %d, SourceRobot %s, HostRobot %s\n",
			task.ID, task.Computation, task.SourceRobotID, task.HostRobotID)
	}
}
