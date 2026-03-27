package main

import (
	"robot-offload/pkg/environment"
	"robot-offload/pkg/simulation"
	"robot-offload/pkg/utils"
)

func main() {
	epochs := 3000
	n_robots := 30

    environment := environment.NewEnvironment(20, 20)
    sim := simulation.NewSimulation(epochs, n_robots, environment, utils.Descending, true)   
	sim.Run() 
}
