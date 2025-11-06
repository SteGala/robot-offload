package simulation

import (
	"fmt"
	"robot-offload/pkg/environment"
	"robot-offload/pkg/fleet"
)

type Simulation struct {
	epochs      int
	curr_epoch  int
	n_robots    int
	fleet       fleet.Fleet
	environment environment.Environment
}

func NewSimulation(epochs int, n_robots int, environment environment.Environment) *Simulation {
	fmt.Println("Initializing simulation with", n_robots, "robots for", epochs, "epochs.")
	return &Simulation{
		n_robots:    n_robots,
		epochs:      epochs,
		curr_epoch:  0,
		fleet:       fleet.NewFleet(n_robots, environment),
		environment: environment,
	}
}

func (s *Simulation) Run() {
	fmt.Println("Starting simulation...")
	for s.curr_epoch < s.epochs {
		fmt.Println("Epoch", s.curr_epoch+1)
		s.fleet.Progress()
		s.curr_epoch++
	}
}
