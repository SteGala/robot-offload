package robot

import (
	"fmt"
	"math/rand"
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
}

type Task struct {
	ID          int
	Computation int
	Source      *Robot
	HostedBy    *Robot
}

func NewRobot(name string, totalBattery int, env *environment.Environment, taskSet *task.TaskSet) Robot {
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

func (r *Robot) Progress() {
	// Example logic to update robot status and battery
	if r.Status == utils.StatusWorking {
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
	} else if r.Status == utils.StatusCharging {
		r.CurrentBattery += r.rechargeRate
		if r.CurrentBattery >= r.TotalBattery {
			r.CurrentBattery = r.TotalBattery
			r.Status = utils.StatusWorking

			// unoffload tasks in case we're fully charged and ready for operation
			for i := 0 ; i < len(r.taskSet.Tasks); i++ {
				task := &r.taskSet.Tasks[i]
				if task.HostRobotID == r.Name {
					task.HostRobotID = task.SourceRobotID
				}
			}
		}
	} else if r.Status == utils.StatusUnavailable {
		r.moveTowardsChargingStation()
	}

	r.Print()
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
