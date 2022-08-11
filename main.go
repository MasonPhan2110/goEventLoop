package main

import (
	"fmt"
	"reflect"

	"github.com/minh211099/goEvenLoop/threadutils"
	ev "github.com/minh211099/goEventLoop/eventloop"
)

func main() {
	fmt.Println("Main Thread ID: ", threadutils.ThreadID())
	var app *ev.App
	app = &ev.App{}
	app.NewApp()
	fmt.Println("var1 = ", reflect.TypeOf(app))
}
