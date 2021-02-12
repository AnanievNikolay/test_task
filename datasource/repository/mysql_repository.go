package repository

import (
	"log"

	repoModel "github.com/AnanievNikolay/test_task/datasource/model"
	"github.com/AnanievNikolay/test_task/presentation/model"
	sqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//NewMySQLRepository ...
func NewMySQLRepository(_connectionString string) IRepository {
	conn, err := gorm.Open(mysql.Open(_connectionString), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("[Error] connecting to database:" + err.Error())
	}
	return &MySQLRepository{
		connection: conn,
	}
}

//MySQLRepository ...
type MySQLRepository struct {
	connection *gorm.DB
}

//Save ...
func (repo *MySQLRepository) Save(response *model.ExternalAPIResponse) bool {
	dbModel := repoModel.New(response)
	tx := repo.connection.Scopes(repoModel.DBModelTable(dbModel)).Create(dbModel)
	err := tx.Error
	if err != nil {
		log.Println("[Error] Error while inserting to database. Error:", err.Error())
		merr, ok := err.(*sqlDriver.MySQLError)
		if ok && merr.Number == 1146 {
			if cerr := repo.createTable(); cerr == nil {
				err = repo.connection.Create(dbModel).Error
			}
		}
	}
	return err == nil
}

func (repo *MySQLRepository) createTable() error {
	return repo.connection.Exec("CREATE TABLE IF NOT EXISTS `test`.`currency_quotes` (" +
		"`id` INT NOT NULL AUTO_INCREMENT," +
		"`response` BLOB NOT NULL," +
		"`received_at` DATETIME NOT NULL," +
		"PRIMARY KEY (`id`));").Error
}
