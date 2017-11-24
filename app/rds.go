package app

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tgsd96/cerviBack/models"
)

// connect to rds instance return the db instance

func ConnectToRDS(adapter, host, user, pass, name string) (*gorm.DB, error) {
	dbString := user + ":" + pass + "@tcp(" + host + ")/" + name + "?parseTime=true"
	//"admin:H?!A4gkm@tcp(cerbackmain.cpkl1sz5etxf.us-west-2.rds.amazonaws.com:3306)/cerback"
	db, err := gorm.Open(adapter, dbString)
	if err != nil {
		log.Fatalf("Error connecting to db: %s", err.Error())
		return nil, err
	}
	return db, nil
}

func CreateTable(db *gorm.DB, table interface{}) error {
	err := db.CreateTable(table).Error
	return err
}
func AddImageToTable(db *gorm.DB, entry *models.ImageStatus) error {

	// use gorm to add imageEntry to table
	err := db.Create(entry).Error
	return err
}

// func CreateNewDB(db *sql.DB, name string) error{
// 	db.
// }
