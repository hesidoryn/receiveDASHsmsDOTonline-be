package db

import "github.com/hesidoryn/receiveDASHsmsDOTonline/model"

// Manager is interface
type Manager interface {
	Sms(to string, limit, page int) ([]model.Sms, error)
	CreateSms(*model.Sms) error
}
