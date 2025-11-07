package fleet

import (
	"fmt"
	"robot-offload/pkg/environment"
	"robot-offload/pkg/robot"
	"robot-offload/pkg/task"
	"robot-offload/pkg/utils"
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
	availableRobotIDS := []int{}

	for i := 0; i < len(f.robots); i++ {
		rStatus := f.robots[i].GetStatus()
		if rStatus == utils.StatusCharging {
			count := 0
			for j := 0; j < len(f.taskSet.Tasks); j++ {
				task := &f.taskSet.Tasks[j]
				if task.HostRobotID == f.robots[i].Name {
					count++
				}
			}
			if count < 2 {
				availableRobotIDS = append(availableRobotIDS, i)
			}
		}
	}

	// find all robots that are in operation saving in a dedicated data structure the list containing for each robot the id and the remaining battery
	possibleOffloaders := []struct {
		id          string
		batteryLeft int
	}{}

	for i := 0; i < len(f.robots); i++ {
		rStatus := f.robots[i].GetStatus()
		if rStatus == utils.StatusWorking {
			batteryLeft := f.robots[i].CurrentBattery
			id := f.robots[i].Name

			count := 0
			for j := 0; j < len(f.taskSet.Tasks); j++ {
				task := &f.taskSet.Tasks[j]
				if task.HostRobotID == f.robots[i].Name {
					count++
				}
			}
			if count == 2 {
				possibleOffloaders = append(possibleOffloaders, struct {
					id          string
					batteryLeft int
				}{id: id, batteryLeft: batteryLeft})
			}
		}
	}

	sortRobots(possibleOffloaders, utils.Descending)
}

func sortRobots(possibleOffloaders []struct{id string; batteryLeft int}, sortingType utils.SortOrder) {
	if sortingType == utils.Descending {
		utils.SortRobotsDescending(possibleOffloaders)
	} else {
		utils.SortRobotsAscending(possibleOffloaders)
	}
}


