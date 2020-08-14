SHELL=/bin/bash
.PHONY: init package install clean

LAMBDA_ZIP=lambda.zip

init:
	export GOPATH=${PWD}
	go get -d -v .
	GOARCH=amd64 GOOS=linux go build -o main

lambda.zip:
	zip -r9 $(LAMBDA_ZIP) . -x Makefile -x main.go -x .git/* -x .git

package: lambda.zip

install: lambda.zip
	aws lambda update-function-code --function-name nameOfYourLambdaFunction --zip-file fileb://lambda.zip --publish --region regionOfYourLambdaFunction

clean:
	rm -rf ./lambda.zip
	rm -rf main

