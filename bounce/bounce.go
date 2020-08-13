package bounce

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type bounce struct {
	NotificationType string `json:"notificationType"`
	Bounce           struct {
		BounceType        string `json:"bounceType"`
		BounceSubType     string `json:"bounceSubType"`
		BouncedRecipients []struct {
			EmailAddress   string `json:"emailAddress"`
			Action         string `json:"action"`
			Status         string `json:"status"`
			DiagnosticCode string `json:"diagnosticCode"`
		} `json:"bouncedRecipients"`
		Timestamp    time.Time `json:"timestamp"`
		FeedbackID   string    `json:"feedbackId"`
		RemoteMtaIP  string    `json:"remoteMtaIp"`
		ReportingMTA string    `json:"reportingMTA"`
	} `json:"bounce"`
	Mail struct {
		Timestamp        time.Time `json:"timestamp"`
		Source           string    `json:"source"`
		SourceArn        string    `json:"sourceArn"`
		SourceIP         string    `json:"sourceIp"`
		SendingAccountID string    `json:"sendingAccountId"`
		MessageID        string    `json:"messageId"`
		Destination      []string  `json:"destination"`
		HeadersTruncated bool      `json:"headersTruncated"`
		Headers          []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"headers"`
		CommonHeaders struct {
			From      []string `json:"from"`
			Date      string   `json:"date"`
			To        []string `json:"to"`
			MessageID string   `json:"messageId"`
			Subject   string   `json:"subject"`
		} `json:"commonHeaders"`
	} `json:"mail"`
}

// NewBounce lala
func NewBounce(w http.ResponseWriter, r *http.Request) {

	config := &aws.Config{
		Region:   aws.String("sa-east-1"),
		Endpoint: aws.String("http://localhost:8000"),
	}

	var t1 bounce

	_ = json.NewDecoder(r.Body).Decode(&t1)

	tableName := "bounces"

	sess := session.Must(session.NewSession(config))
	svc := dynamodb.New(sess)

	av, err := dynamodbattribute.MarshalMap(t1)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
	}

	fmt.Println("Successfully added to table " + tableName)

	fmt.Println(t1)
	json.NewEncoder(w).Encode(t1)
}
