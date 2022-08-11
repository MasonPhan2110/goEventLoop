package eventloop

import "fmt"

type App struct {
	events chan IEvent
}

func (app *App) exec() {
	for event := range app.events {
		event.process()
	}
}

func (app *App) NewApp() {
	fmt.Println("New App run")
}
