package main

import (
	"fmt"

	"github.com/penkovski/statemachine/pkg/statemachine"
)

func main() {
	sm := statemachine.New()

	done := make(chan chan struct{})
	go sm.Run(done)

	fmt.Println(sm.State())

	sm.NewState("hello")
	fmt.Println(sm.State())

	sm.NewState("bye")
	fmt.Println(sm.State())

	q := make(chan struct{})
	done <- q
	<-q
}
