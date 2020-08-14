package bounce

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type rawBounce struct {
	Item struct {
		BouncedRecipients []struct {
			EmailAddress   string `json:"emailAddress"`
			Action         string `json:"action"`
			Status         string `json:"status"`
			DiagnosticCode string `json:"diagnosticCode"`
		} `json:"bouncedRecipients"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"bounce"`
	Mail struct {
		Source string `json:"source"`
	} `json:"mail"`
}

// Bounced recipients informations
type bouncedrecipients struct {
	EmailAddress   string `json:"email"`
	Timestamp      string `json:"timestamp"`
	DiagnosticCode string `json:"diagnosticCode"`
	From           string `json:"source"`
	Description    string `json:"description"`
}

// PutBounce - responsible for put bounce on DynamoDB
func PutBounce(ctx context.Context, snsEvent events.SNSEvent) {

	awsConfig := &aws.Config{
		Region: aws.String("sa-east-1"),
	}

	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		in := []byte(snsRecord.Message)

		var rb rawBounce

		if err := json.Unmarshal(in, &rb); err != nil {
			panic(err)
		}

		sess := session.Must(session.NewSession(awsConfig))
		svc := dynamodb.New(sess)

		var b bouncedrecipients
		for _, brData := range rb.Item.BouncedRecipients {
			b.DiagnosticCode = brData.Status
			b.EmailAddress = brData.EmailAddress
			b.Timestamp = rb.Item.Timestamp.String()
			b.From = rb.Mail.Source
			b.Description = brData.DiagnosticCode

			av, err := dynamodbattribute.MarshalMap(b)
			if err != nil {
				log.Println(err.Error())
				os.Exit(1)
			}

			input := &dynamodb.PutItemInput{
				Item:      av,
				TableName: aws.String("prod-sesBounces"),
			}

			_, err = svc.PutItem(input)
			if err != nil {
				log.Println(err.Error())
			}

			log.Println("Successfully added to table prod-sesBounces")

		}

	}
}
