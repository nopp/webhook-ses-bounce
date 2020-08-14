# Webhook SES Bouce - AWS Lambda function

### 1) Create a Amazon DynamoDB with keys:
```Partition Key: email(string) and Sort Key: timestamp(string)```
### 2) Create a Amazon Lambda function
  `Edit the table name and region on bounce.go`
  
  `Region: aws.String("sa-east-1")`
  `TableName: aws.String("prod-sesBounces")`
  
  Lambda role:
  
  ```
    {
      "Version": "2012-10-17",
      "Statement": [
          {
              "Sid": "ReadWriteTable",
              "Effect": "Allow",
              "Action": [
                  "dynamodb:BatchGetItem",
                  "dynamodb:GetItem",
                  "dynamodb:Query",
                  "dynamodb:Scan",
                  "dynamodb:BatchWriteItem",
                  "dynamodb:PutItem",
                  "dynamodb:UpdateItem"
              ],
              "Resource": "arn:aws:dynamodb:*:*:table/prod-sesBounces"
          },
          {
              "Sid": "CreateLogGroup",
              "Effect": "Allow",
              "Action": "logs:CreateLogGroup",
              "Resource": "*"
          }
      ]
  }
```
### 3) Create Amazon SNS Topic and subscription to Lambda.
