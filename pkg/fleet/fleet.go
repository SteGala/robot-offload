package fleet

import (
	"fmt"
	"os"
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
func NewFleet(n_robots int, env environment.Environment, taskSet *task.TaskSet, report_directory string) Fleet {
	f := Fleet{}
	f.robots = []robot.Robot{}
	f.taskSet = taskSet

	// create the subdirectory robots inside the report directory
	robotsReportDirectory := fmt.Sprintf("%s/robots", report_directory)
	err := os.MkdirAll(robotsReportDirectory, 0755)
	if err != nil {
		fmt.Println("Error creating robots report directory:", err)
	}

	for i := 0; i < n_robots; i++ {
		robotReportFile := fmt.Sprintf("%s/%s", robotsReportDirectory, fmt.Sprintf("robot-%d.csv", i))
		robotName := fmt.Sprintf("robot-%d", i)
		newRobot := robot.NewRobot(robotName, 100, &env, taskSet, robotReportFile)
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

func (f *Fleet) orchestrateTasks() error {
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
	possibleOffloaders := []utils.RobotInfo{}

	for i := 0; i < len(f.robots); i++ {
		rStatus := f.robots[i].GetStatus()
		if rStatus == utils.StatusWorking {
			batteryLeft := f.robots[i].CurrentBattery
			id := f.robots[i].Name

			task, err := f.taskSet.GetRobotTask(id)
			if err == nil {
				if task.HostRobotID == task.SourceRobotID {
					possibleOffloaders = append(possibleOffloaders, utils.NewRobotInfo(id, batteryLeft))
				}
			} else {
				return err
			}
		}
	}

	// if len(possibleOffloaders) > 0 {
	// 	fmt.Println("Possible offloaders:")
	// 	for _, offloader := range possibleOffloaders {
	// 		fmt.Printf("Robot %s\n", offloader.Id)
	// 	}
	// }
	// if len(availableRobotIDS) > 0 {
	// 	fmt.Println("Available robot IDs:")
	// 	for _, id := range availableRobotIDS {
	// 		fmt.Printf("Robot ID %d\n", id)
	// 	}
	// }

	if len(possibleOffloaders) > 0 {
		sortRobots(possibleOffloaders, utils.Descending)

		for i := 0 ; i < min(len(availableRobotIDS), len(possibleOffloaders)) ; i++ {
			availableRobotID := availableRobotIDS[i]
			offloader := possibleOffloaders[i]

			t, err := f.taskSet.GetRobotTask(offloader.Id)
			if err != nil {
				return err
			}
			t.Offload(fmt.Sprintf("robot-%d", availableRobotID))
		}
	}
	

	return nil
}

func sortRobots(possibleOffloaders []utils.RobotInfo, sortingType utils.SortOrder) {
	if sortingType == utils.Descending {
		utils.SortRobotsDescending(possibleOffloaders)
	} else {
		utils.SortRobotsAscending(possibleOffloaders)
	}
}
