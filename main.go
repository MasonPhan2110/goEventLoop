package main

import (
	"fmt"

	ev "github.com/minh211099/goEvenLoop/EventLoop"
	"github.com/nguyenzung/go-event-loop/threadutils"
)

func main() {
	fmt.Println("Main Thread ID: ", threadutils.ThreadID())
	var app *ev.App = nil
	fmt.Println(app)
}
