package database

import (
	"fmt"
	"os"

	"github.com/AnjaneyuluBatta505/gin-graphql-postgres/graph/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type dbConfig struct {
	Path     string
	Username string
	Dbname   string
	Password string
	Config   string
}

var config = dbConfig{"127.0.0.1:3306", "root", "graphql", "123456", "charset=utf8mb4&parseTime=True&loc=Local"}

func getDatabaseUrl() string {
	// return fmt.Sprintf(
	// 	"host=%s port=%d user=%s dbname=%s password=%s",
	// 	config.host, config.port, config.user, config.dbname, config.password)

	// return fmt.Sprintf(
	// 	"host=%s port=%d user=%s dbname=%s password=%s",
	// 	config.host, config.port, config.user, config.dbname, config.password)
	var con string = config.Username + ":" + config.Password + "@(" + config.Path + ")/" + config.Dbname + "?" + config.Config
	fmt.Println(con)
	return con
}

func GetDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", getDatabaseUrl())
	if err != nil {
		fmt.Println("MySQL 启动异常！")
		os.Exit(0)
	}
	return db, err
}

func RunMigrations(db *gorm.DB) {
	if !db.HasTable(&model.Question{}) {
		db.CreateTable(&model.Question{})
	}
	if !db.HasTable(&model.Choice{}) {
		db.CreateTable(&model.Choice{})
		db.Model(&model.Choice{}).AddForeignKey("question_id", "questions(id)", "CASCADE", "CASCADE")
	}
}
