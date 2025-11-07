package utils

import "math/rand"

type Position struct {
	X int
	Y int
}

type Status string

const (
	StatusWorking     = "WORKING"
	StatusCharging    = "CHARGING"
	StatusUnavailable = "UNAVAILABLE"
)

type SortOrder int

const (
	Descending SortOrder = iota
	Ascending
)

func SortRobotsDescending(possibleOffloaders []struct{id string; batteryLeft int}) {
	for i := 0; i < len(possibleOffloaders)-1; i++ {
		for j := i + 1; j < len(possibleOffloaders); j++ {
			if possibleOffloaders[i].batteryLeft < possibleOffloaders[j].batteryLeft {
				possibleOffloaders[i], possibleOffloaders[j] = possibleOffloaders[j], possibleOffloaders[i]
			}
		}
	}
}

func SortRobotsAscending(possibleOffloaders []struct{id string; batteryLeft int}) {
	for i := 0; i < len(possibleOffloaders)-1; i++ {
		for j := i + 1; j < len(possibleOffloaders); j++ {
			if possibleOffloaders[i].batteryLeft > possibleOffloaders[j].batteryLeft {
				possibleOffloaders[i], possibleOffloaders[j] = possibleOffloaders[j], possibleOffloaders[i]
			}
		}
	}
}

var allStatuses = []Status{StatusWorking, StatusCharging, StatusUnavailable}

func RandomStatus() Status {
	return allStatuses[rand.Intn(len(allStatuses)-1)]
}
