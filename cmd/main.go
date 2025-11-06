package main

import (
	"robot-offload/pkg/environment"
	"robot-offload/pkg/simulation"
)

func main() {
    environment := environment.NewEnvironment(20, 20)
    simulation.NewSimulation(50, 1, environment).Run()    
}
