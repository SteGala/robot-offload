package utils

import "math/rand"

type Position struct {
	X int
	Y int
}

type Status string

const (
	StatusWorking    = "WORKING"
	StatusCharging   = "CHARGING"
	StatusUnavailable = "UNAVAILABLE"
)

var allStatuses = []Status{StatusWorking, StatusCharging, StatusUnavailable}

func RandomStatus() Status {
	return allStatuses[rand.Intn(len(allStatuses)-1)]
}