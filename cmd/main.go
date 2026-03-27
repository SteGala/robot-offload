package main

import (
	"robot-offload/pkg/environment"
	"robot-offload/pkg/simulation"
)

func main() {
    environment := environment.NewEnvironment(20, 20)
    sim := simulation.NewSimulation(30, 10, environment)   
	sim.Run() 
}
