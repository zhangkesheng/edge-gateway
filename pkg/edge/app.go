package edge

type App struct {
	apis []Api
}

func New() *App {
	return &App{}
}

func (app *App) Edges() []Api {
	return app.apis
}

func (app *App) Reload() {
}
