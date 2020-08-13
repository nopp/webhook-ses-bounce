package bounce

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"webhook-ses-bounce/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Received from Amazon SES + SNS Topic (subscription)
type rawBounce struct {
	FeedbackID string
	Item       struct {
		BouncedRecipients []struct {
			EmailAddress   string `json:"emailAddress"`
			Action         string `json:"action"`
			Status         string `json:"status"`
			DiagnosticCode string `json:"diagnosticCode"`
		} `json:"bouncedRecipients"`
		Timestamp  time.Time `json:"timestamp"`
		FeedbackID string    `json:"feedbackId"`
	} `json:"bounce"`
	Mail struct {
		Source   string `json:"source"`
		SourceIP string `json:"sourceIp"`
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
func PutBounce(w http.ResponseWriter, r *http.Request) {

	var rb rawBounce

	config := common.LoadConfiguration()

	awsConfig := &aws.Config{
		Region: aws.String(config.AwsRegion),
	}

	_ = json.NewDecoder(r.Body).Decode(&rb)

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
			common.Message(w, err.Error())
			os.Exit(1)
		}

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(config.TableName),
		}

		_, err = svc.PutItem(input)
		if err != nil {
			common.Message(w, err.Error())
		}

		common.Message(w, "Successfully added to table "+config.TableName)

	}
}
