package fleet

import (
	"fmt"
	"robot-offload/pkg/environment"
	"robot-offload/pkg/robot"
)

type Fleet struct {
	robots []robot.Robot
}

// initialize fleet
func NewFleet(n_robots int, env environment.Environment) Fleet {
	f := Fleet{}
	f.robots = []robot.Robot{}

	for i := 0; i < n_robots; i++ {
		robotName := fmt.Sprintf("Robot-%d", i+1)
		newRobot := robot.NewRobot(robotName, 100, &env)
		f.robots = append(f.robots, newRobot)
	}
	
	return f
}

func (f *Fleet) Progress() {
	for i := range f.robots {
		f.robots[i].Progress()
	}
}