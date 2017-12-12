package db

import "github.com/hesidoryn/receiveDASHsmsDOTonline/model"

// Manager is interface
type Manager interface {
	Sms(page, limit int) ([]model.Sms, error)
	CreateSms(*model.Sms) error
}
