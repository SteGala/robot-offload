package simulation

import (
	"fmt"
	"os"
	"robot-offload/pkg/environment"
	"robot-offload/pkg/fleet"
	"robot-offload/pkg/task"
	"time"
)

type Simulation struct {
	epochs      int
	curr_epoch  int
	n_robots    int
	fleet       fleet.Fleet
	environment environment.Environment
	report_directory string
}

func NewSimulation(epochs int, n_robots int, environment environment.Environment) *Simulation {
	fmt.Println("Initializing simulation with", n_robots, "robots for", epochs, "epochs.")

	taskSet := task.NewTaskSet(n_robots)

	// create e report directory for the current simulation run using the current timestamp as name
	report_directory := fmt.Sprintf("res/report-%d", time.Now().Unix())
	err := os.MkdirAll(report_directory, 0755)
	if err != nil {
		fmt.Println("Error creating report directory:", err)
		return nil
	}

	return &Simulation{
		n_robots:    n_robots,
		epochs:      epochs,
		curr_epoch:  0,
		fleet:       fleet.NewFleet(n_robots, environment, taskSet, report_directory),
		environment: environment,
		report_directory: report_directory,
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
