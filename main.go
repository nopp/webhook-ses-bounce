package main

import (
	"ses-bounces-webhook/bounce"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(bounce.PutBounce)

}
