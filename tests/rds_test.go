package tests

import (
	"testing"

	"github.com/tgsd96/cerviBack/app"
	"github.com/tgsd96/cerviBack/models"
)

func TestDbConnection(t *testing.T) {
	_, err := app.ConnectToRDS()
	if err != nil {
		t.Errorf("Test failed, error: %s", err.Error())
	}
}

// func TestTableCreate(t *testing.T) {
// 	db, err := app.ConnectToRDS()
// 	if err != nil {
// 		t.Errorf("Unable to connect, error: %s", err.Error())
// 	}
// 	err = app.CreateTable(db, models.ImageStatus{})
// 	if err != nil {
// 		t.Errorf("Error creating table, error : %s", err.Error())
// 	}
// }
func TestTableEntry(t *testing.T) {
	db, err := app.ConnectToRDS()
	if err != nil {
		t.Errorf("Test failed, error: %s", err.Error())
	}
	image := models.ImageStatus{
		ImageKey: "testingkey1.png",
		UserID:   "testinguserid2",
		Status:   "INQUEUE",
		Type1:    99.1,
		Type2:    0.1,
		Type3:    0.8,
	}
	err = app.AddImageToTable(db, &image)
	if err != nil {
		t.Errorf("Unable to add to table, error: %s", err.Error())
	}
	// var testRow models.ImageStatus
	// db.Last(&testRow)
	// t.Logf("\n Received the row : %v", testRow)

}
