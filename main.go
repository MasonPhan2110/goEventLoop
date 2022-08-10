package main

import (
	"fmt"

	ev "github.com/minh211099/goEvenLoop/eventloop"
	"github.com/minh211099/goEvenLoop/threadutils"
)

func main() {
	fmt.Println("Main Thread ID: ", threadutils.ThreadID())
	var app *ev.App = nil
	fmt.Println(app)
}
