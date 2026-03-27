package robot

import (
	"fmt"
	"math/rand"
	"os"
	"robot-offload/pkg/environment"
	"robot-offload/pkg/task"
	"robot-offload/pkg/utils"
)

type Robot struct {
	Name              string
	TotalBattery      int
	CurrentBattery    int
	Position          utils.Position
	Status            utils.Status
	Map               *environment.Environment
	chargingThreshold int
	consumptionRate   int
	rechargeRate      int
	taskSet           *task.TaskSet
	reportFile        string
}

type Task struct {
	ID          int
	Computation int
	Source      *Robot
	HostedBy    *Robot
}

var report_header = "Battery,PositionX,PositionY,Status,HostedTasks\n"

func NewRobot(name string, totalBattery int, env *environment.Environment, taskSet *task.TaskSet, reportFile string) Robot {
	// create the csv report file for the robot
	// if the file already exists, it should be overwritten
	_, err := os.Create(reportFile)
	if err != nil {
		fmt.Println("Error creating report file for robot", name, ":", err)
	}

	// write the header
	// the header should contain the following columns: Epoch, Battery, PositionX, PositionY, Status, HostedTasks, BatteryLevel
	f, err := os.OpenFile(reportFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening report file for robot", name, ":", err)
	}
	defer f.Close()

	_, err = f.WriteString(report_header)
	if err != nil {
		fmt.Println("Error writing header to report file for robot", name, ":", err)
	}

	return Robot{
		Name:              name,
		TotalBattery:      totalBattery,
		CurrentBattery:    rand.Intn(totalBattery + 1),
		Position:          utils.Position{X: rand.Intn(env.GetXSize()), Y: rand.Intn(env.GetYSize())},
		Status:            utils.RandomStatus(),
		Map:               env,
		consumptionRate:   5,
		rechargeRate:      20,
		chargingThreshold: 20,
		taskSet:           taskSet,
		reportFile:        reportFile,
	}
}

func (r *Robot) GetStatus() utils.Status {
	return r.Status
}

func (r *Robot) Print() {
	hostedTaskIDs := []int{}
	for i := 0; i < len(r.taskSet.Tasks); i++ {
		task := r.taskSet.Tasks[i]
		if task.HostRobotID == r.Name {
			hostedTaskIDs = append(hostedTaskIDs, task.ID)
		}
	}
	fmt.Printf("%s: Battery %d/%d, Position (%d, %d), Status %s, Hosted Tasks %v\n",
		r.Name, r.CurrentBattery, r.TotalBattery, r.Position.X, r.Position.Y, r.Status, hostedTaskIDs)
}

func (r *Robot) LogOnReportFile() {
	// log the current state of the robot to the report file
	hostedTaskIDs := []int{}
	for i := 0; i < len(r.taskSet.Tasks); i++ {
		task := r.taskSet.Tasks[i]
		if task.HostRobotID == r.Name {
			hostedTaskIDs = append(hostedTaskIDs, task.ID)
		}
	}

	f, err := os.OpenFile(r.reportFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening report file for robot", r.Name, ":", err)
		return
	}
	defer f.Close()

	logLine := fmt.Sprintf("%d,%d,%d,%s,%v\n",
		r.CurrentBattery, r.Position.X, r.Position.Y, r.Status, hostedTaskIDs)

	_, err = f.WriteString(logLine)
	if err != nil {
		fmt.Println("Error writing log line to report file for robot", r.Name, ":", err)
	}
}

func (r *Robot) Progress() {
	// Example logic to update robot status and battery
	switch r.Status {
	case utils.StatusWorking:
		r.CurrentBattery -= r.consumptionRate

		for i := 0; i < len(r.taskSet.Tasks); i++ {
			task := &r.taskSet.Tasks[i]
			if task.HostRobotID == r.Name {
				r.CurrentBattery -= task.Computation
			}
		}

		if r.CurrentBattery <= r.chargingThreshold {
			r.Status = utils.StatusUnavailable
		}

		r.move()

	case utils.StatusCharging:
		r.CurrentBattery += r.rechargeRate
		if r.CurrentBattery >= r.TotalBattery {
			r.CurrentBattery = r.TotalBattery
			r.Status = utils.StatusWorking

			// unoffload tasks in case we're fully charged and ready for operation
			for i := 0; i < len(r.taskSet.Tasks); i++ {
				task := &r.taskSet.Tasks[i]
				if task.HostRobotID == r.Name {
					task.HostRobotID = task.SourceRobotID
				}
			}
		}

	case utils.StatusUnavailable:
		r.moveTowardsChargingStation()
	}

	//r.Print()
	r.LogOnReportFile()
}

func (r *Robot) moveTowardsChargingStation() {
	xCharging, yCharging := r.Map.GetChargingPosition()

	if r.Position.X == xCharging && r.Position.Y == yCharging {
		r.Status = utils.StatusCharging
		return
	}

	if r.Position.X < xCharging {
		r.Position.X++
	} else if r.Position.X > xCharging {
		r.Position.X--
	}

	if r.Position.Y < yCharging {
		r.Position.Y++
	} else if r.Position.Y > yCharging {
		r.Position.Y--
	}
}

func (r *Robot) move() {
	destX := rand.Intn(r.Map.GetXSize())
	destY := rand.Intn(r.Map.GetYSize())

	if r.Position.X < destX {
		r.Position.X++
	} else if r.Position.X > destX {
		r.Position.X--
	}

	if r.Position.Y < destY {
		r.Position.Y++
	} else if r.Position.Y > destY {
		r.Position.Y--
	}
}
