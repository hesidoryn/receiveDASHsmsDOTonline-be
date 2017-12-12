package db

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/hesidoryn/receiveDASHsmsDOTonline/model"
)

// DynamoManager is struct that implements Manager interface
type DynamoManager struct {
	*dynamodb.DynamoDB
}

// CreateDynamoManager returns DynamoManager
func CreateDynamoManager() (*DynamoManager, error) {
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		return nil, err
	}

	db := dynamodb.New(sess)
	return &DynamoManager{db}, nil
}

// Sms is realisation of Manager's Sms method
func (m *DynamoManager) Sms(page, limit int) ([]model.Sms, error) {
	return []model.Sms{}, nil
}

// CreateSms is realisation of Manager's CreateSms method
func (m *DynamoManager) CreateSms(s *model.Sms) error {
	return nil
}
