package faceit_cc

import (
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql" // go mysql driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql import driver for gorm
)

// Database represent the repository model
type Database struct {
	Db       *gorm.DB
	Host     string
	Port     int64
	User     string
	DBName   string
	Password string
}

// NewDatabase returns a Handler struct with a correcty database instance opened
func NewDatabase() *Handler {
	db := CreateDbInstance()
	err := db.RetrieveDbConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = db.Open()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &Handler{
		Database: db,
	}
}

// CreateDbInstance returns a Database struct initialized
// Return Database entity
func CreateDbInstance() Database {
	return Database{}
}

// RetrieveDbConfig retrieves the needed environment variables to
// create the connection string to conect to the database
// Return error | nil
func (d *Database) RetrieveDbConfig() error {

	DB_HOST, err := GetSetting("DB_HOST")
	if err != nil {
		log.Println(err)
		return err
	}

	DB_PORT, err := GetSetting("DB_PORT")
	if err != nil {
		log.Println(err)
		return err
	}
	portInt, err := strconv.ParseInt(DB_PORT, 10, 64)
	if err != nil {
		log.Println(err)
		return err
	}

	DB_USER, err := GetSetting("DB_USERNAME")
	if err != nil {
		log.Println(err)
		return err
	}

	DB_PASSWORD, err := GetSetting("DB_PASSWORD")
	if err != nil {
		log.Println(err)
		return err
	}

	DB_DATABASE, err := GetSetting("DB_DATABASE")
	if err != nil {
		log.Println(err)
		return err
	}

	d.Host = DB_HOST
	d.Port = portInt
	d.User = DB_USER
	d.Password = DB_PASSWORD
	d.DBName = DB_DATABASE

	return nil
}

// GenerateConnectionString generates the connection string
// Returns an string to use as connection string
func (d *Database) GenerateConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DBName,
	)
}

type Result struct {
	Success bool
	Error   error
	Db      *gorm.DB
	Kafka   *KafkaInstance
}

// Open tries to open a database connection. If the database is not available
// it will try again during 20 seconds
func (d *Database) Open() error {
	log.Println("Opening database connection ...")
	constr := d.GenerateConnectionString()
	c := make(chan Result)
	waitfor := 20

	for i := 0; i < waitfor; i++ {
		time.Sleep(1 * time.Second)

		go func(constr string) {
			db, err := gorm.Open("mysql", constr)
			if err != nil {
				log.Printf("error opening mysql connection. err: %v", err)
				result := Result{
					Success: false,
					Error:   err,
					Db:      nil,
				}
				c <- result
			}

			err = db.DB().Ping()
			if err == nil {
				result := Result{
					Success: true,
					Error:   nil,
					Db:      db,
				}
				c <- result
			}
		}(constr)

		select {
		case res := <-c:
			if res.Success {
				d.Db = res.Db
				fmt.Println("***** Database is ready *****")
				return nil
			}
		case <-time.After(time.Duration(waitfor) * time.Second):
			fmt.Println("timeout %n", waitfor)
		}
	}
	return fmt.Errorf("error opening database connection")
}
