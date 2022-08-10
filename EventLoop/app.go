package eventloop

type App struct {
	events chan Ievent
}

func (app *App) exec() {
	for event := range app.events {
		event.process()
	}
}

func (app *App) NewApp() {}
