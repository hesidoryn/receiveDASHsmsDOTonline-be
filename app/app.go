package app

import (
	"github.com/hesidoryn/receiveDASHsmsDOTonline/db"
)

// App contains application business logic
type App struct {
	manager db.Manager
}

// InitApp returns App
func InitApp(manager db.Manager) *App {
	return &App{manager}
}
