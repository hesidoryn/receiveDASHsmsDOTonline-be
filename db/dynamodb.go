package db

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/hesidoryn/receiveDASHsmsDOTonline/model"
)

// DynamoManager is struct that implements Manager interface
type DynamoManager struct {
	TableName string
	*dynamodb.DynamoDB
}

// CreateDynamoManager returns DynamoManager
func CreateDynamoManager(tableName string) (*DynamoManager, error) {
	sess, err := session.NewSession(&aws.Config{})
	if err != nil {
		return nil, err
	}

	db := dynamodb.New(sess)
	return &DynamoManager{tableName, db}, nil
}

// Sms is realisation of Manager's Sms method
func (m *DynamoManager) Sms(to string, limit, page int) (result []model.Sms, err error) {
	params := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":val": {
				S: aws.String(to),
			},
		},
		KeyConditionExpression: aws.String("numberTo = :val"),
		TableName:              aws.String(m.TableName),
		ScanIndexForward:       aws.Bool(false),
		Limit:                  aws.Int64(int64(limit)),
	}

	pageNum := 0
	err = m.QueryPages(params,
		func(o *dynamodb.QueryOutput, lastPage bool) bool {
			pageNum++

			if pageNum != page {
				return true
			}

			dynamodbattribute.UnmarshalListOfMaps(o.Items, &result)
			return false
		})

	return
}

// CreateSms is realisation of Manager's CreateSms method
func (m *DynamoManager) CreateSms(s *model.Sms) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"numberTo": {
				S: aws.String(s.To),
			},
			"numberFrom": {
				S: aws.String(s.From),
			},
			"body": {
				S: aws.String(s.Body),
			},
			"sentAt": {
				N: aws.String(strconv.Itoa(int(s.SentAt))),
			},
		},
		TableName: aws.String(m.TableName),
	}

	_, err := m.PutItem(input)
	return err
}
