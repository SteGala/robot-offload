package fleet

import (
	"fmt"
	"math/rand"
	"robot-offload/pkg/environment"
	"robot-offload/pkg/robot"
	"robot-offload/pkg/task"
)

type Fleet struct {
	robots  []robot.Robot
	taskSet *task.TaskSet
}

// initialize fleet
func NewFleet(n_robots int, env environment.Environment, taskSet *task.TaskSet) Fleet {
	f := Fleet{}
	f.robots = []robot.Robot{}
	f.taskSet = taskSet

	for i := 0; i < n_robots; i++ {
		robotName := fmt.Sprintf("robot-%d", i)
		newRobot := robot.NewRobot(robotName, 100, &env, taskSet)
		f.robots = append(f.robots, newRobot)
	}

	return f
}

func (f *Fleet) Progress() {
	for i := range f.robots {
		f.robots[i].Progress()
	}

	f.orchestrateTasks()
}

func (f *Fleet) orchestrateTasks() {
	for i := 0; i < len(f.taskSet.Tasks); i++ {
		f.taskSet.Tasks[i].HostRobotID = fmt.Sprintf("robot-%d", rand.Intn(len(f.robots)))
	}
}
