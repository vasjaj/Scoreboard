package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var rankQuery string = "SELECT scores.name, (SELECT COUNT(*) FROM user_scores AS sc WHERE sc.points <= scores.points) AS position, scores.points FROM user_scores AS scores"

type UserScore struct {
	gorm.Model

	Name   string `gorm:"not null;unique_index:name_index"`
	Points int64  `gorm:"not null"`
}

type UserScoreWithPosition struct {
	UserScore
	Position int64
}

func setupDB() (*gorm.DB, error) {
	// cURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))
	cURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "user", "password", "127.0.0.1:3306", "scoreboard")
	fmt.Println(cURL)
	db, err := gorm.Open("mysql", cURL)
	if err != nil {
		return db, err
	}

	db.LogMode(true)

	db.AutoMigrate(&UserScore{})

	return db, nil
}

func findScore(name string) (UserScore, error) {
	var score UserScore

	err := db.First(&score, "name = ?", name).Error

	return score, err
}

func updateScore(userScore UserScore, newScore int64) error {
	return db.Model(userScore).Update("points", newScore).Error
}

func saveScore(name string, score int64) error {
	return db.Create(&UserScore{Name: name, Points: score}).Error
}

func getRanks(name string) ([]UserScoreWithPosition, error) {
	var scores []UserScoreWithPosition

	err := db.Raw(rankQuery).Scan(&scores).Error

	return scores, err
}
