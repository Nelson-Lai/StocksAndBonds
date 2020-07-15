package locks

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Locks is necessary to avoid race conditions with multiple updaters in dynamoDB gamestate.
// Because I suck and marshalled the entire gamestate as a single value, it's highly prone to races.

const (
	maxRetries = 5
	lockTable  = "StocksAndBondsLock"
)

func Lock() error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ddb := dynamodb.New(sess)

	lockRetries := 0

	for {
		output, err := ddb.GetItem(&dynamodb.GetItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"Name": {
					S: aws.String("Lock"),
				},
			},
			TableName: aws.String(lockTable),
		})
		if err != nil {
			return err
		}

		if output.GoString() == "Free" {
			break
		}
		if lockRetries >= maxRetries {
			return errors.New("Could not obtain ddb lock to update state")
		}

		lockRetries++
		time.Sleep(time.Duration(100*lockRetries) * time.Millisecond)
	}
	ddb.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String("Lock"),
			},
			"State": {
				S: aws.String("Taken"),
			},
		},
	})

	return nil
}

func Unlock() error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	ddb := dynamodb.New(sess)

	_, err := ddb.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(lockTable),
		Item: map[string]*dynamodb.AttributeValue{
			"Name": {
				S: aws.String("Lock"),
			},
			"State": {
				S: aws.String("Free"),
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
