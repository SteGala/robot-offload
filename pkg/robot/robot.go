package robot

import (
	"fmt"
	"math/rand"
	"robot-offload/pkg/environment"
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
}

func NewRobot(name string, totalBattery int, env *environment.Environment) Robot {
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
	}
}

func (r *Robot) Print() {
	fmt.Printf("Robot %s: Battery %d/%d, Position (%d, %d), Status %s\n",
		r.Name, r.CurrentBattery, r.TotalBattery, r.Position.X, r.Position.Y, r.Status)
}

func (r *Robot) Progress() {
	// Example logic to update robot status and battery
	if r.Status == utils.StatusWorking {
		r.CurrentBattery -= r.consumptionRate
		if r.CurrentBattery <= r.chargingThreshold {
			r.Status = utils.StatusUnavailable
		}

		r.move()
	} else if r.Status == utils.StatusCharging {
		r.CurrentBattery += r.rechargeRate
		if r.CurrentBattery >= r.TotalBattery {
			r.CurrentBattery = r.TotalBattery
			r.Status = utils.StatusWorking
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

// func (r *Robot) moveNext() {
// 	// Example logic to move robot randomly within the environment bounds

// 	xMove := rand.Intn(3) - 1
// 	yMove := rand.Intn(3) - 1

// 	newX := r.Position.X + xMove
// 	newY := r.Position.Y + yMove
// 	if newX >= 0 && newX < r.Map.GetXSize() {
// 		r.Position.X = newX
// 	}
// 	if newY >= 0 && newY < r.Map.GetYSize() {
// 		r.Position.Y = newY
// 	}
// }

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
