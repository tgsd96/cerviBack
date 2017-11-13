package tests

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/tgsd96/cerviBack/app"
	"github.com/tgsd96/cerviBack/models"
)

func TestSQSPublish(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		t.Errorf("Unable to connect to aws, error : %s", err.Error())
	}
	msg := models.SqsMessage{
		ImageKey: "testingImageKey",
		UserID:   "testinguserid",
	}
	err = app.PushToSQS(sess, "https://sqs.us-west-2.amazonaws.com/907743384002/jobQueue", msg)
	if err != nil {
		t.Errorf("Unable to push to SNS, error: %s", err.Error())
	}
}
