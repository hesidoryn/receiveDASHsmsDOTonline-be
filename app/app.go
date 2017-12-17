package app

import (
	"github.com/hesidoryn/receiveDASHsmsDOTonline/db"
	"github.com/hesidoryn/receiveDASHsmsDOTonline/model"
)

// App contains application business logic
type App struct {
	manager db.Manager
}

// Init returns App
func Init(manager db.Manager) *App {
	return &App{manager}
}

// Sms gets list of sms to `TO` number from databse
// using limitation and pagination
func (a *App) Sms(to string, limit, page int) ([]model.Sms, error) {
	return a.manager.Sms(to, limit, page)
}

// CreateSms inserts new sms message in databse
func (a *App) CreateSms(s *model.Sms) error {
	return a.manager.CreateSms(s)
}
