package app

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/tgsd96/cerviBack/models"
)

func PushToSQS(sess *session.Session, queueUrl string, message models.SqsMessage) error {
	sqsClient := sqs.New(sess)
	msg, _ := json.Marshal(message)
	stringMsg := string(msg)
	params := &sqs.SendMessageInput{
		MessageBody: aws.String(stringMsg),
		QueueUrl:    aws.String(queueUrl),
	}
	resp, err := sqsClient.SendMessage(params)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}
