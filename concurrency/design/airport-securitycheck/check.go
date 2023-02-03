package main

import (
	"fmt"
	"time"
)

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTmCost = 180
)

func idCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	fmt.Printf("\tgoroutine-%s-idCheck: idCheck ok\n", id)
	return idCheckTmCost
}

func bodyCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	fmt.Printf("\tgoroutine-%s-bodyCheck: bodyCheck ok\n", id)
	return bodyCheckTmCost
}

func xRayCheck(id string) int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTmCost))
	fmt.Printf("\tgoroutine-%s-xRayCheck: xRayCheck ok\n", id)
	return xRayCheckTmCost
}
