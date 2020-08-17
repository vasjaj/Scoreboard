package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/now"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

var PositionsQuery string = "select user_scores.name, user_scores.points, Position () over (order by user_scores.points desc) AS position from user_scores"

// var PositionQuery string = "SELECT scores.name, scores.points, Position () OVER (PARTITION BY scores.name ORDER BY scores.points DESC) AS position FROM user_scores AS scores WHERE scores.name = ?"

type UserScore struct {
	gorm.Model

	Name     string `gorm:"not null;unique_index:name_index"`
	Points   int64  `gorm:"not null"`
	Position int64  `gorm:"-"`
}

func setupDB() (*gorm.DB, error) {
	connURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_DATABASE"))

	log.Println("DB connection URL: ", connURL)

	db, err := gorm.Open("mysql", connURL)

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

func findScoreWithPosition(name string) (UserScore, error) {
	score := UserScore{}
	scores, err := findScoresWithPositions(false)

	if err != nil {
		return score, err
	}

	for i := range scores {
		if scores[i].Name == name {
			return scores[i], nil
		}
	}

	return score, errors.New("user not found by name")
}

func updateScore(name string, points int64) error {
	return db.Table("user_scores").Where("name = ?", name).Update("points", points).Error
}

func saveScore(name string, points int64) error {
	return db.Create(&UserScore{Name: name, Points: points}).Error
}

func findScoresWithPositions(monthly bool) ([]UserScore, error) {
	var scores []UserScore
	var err error

	if monthly {
		timeFrom := now.BeginningOfMonth()

		err = db.Order("points DESC").Where("updated_at IS NULL AND created_at > ? || updated_at IS NOT NULL AND updated_at > ?", timeFrom, timeFrom).Find(&scores).Error

	} else {
		err = db.Order("points DESC").Find(&scores).Error
	}

	if err != nil {
		return scores, err
	}

	assignPositions(scores)

	return scores, nil
}

func findScoresWithPositionsByPage(name string, pageNum, pageSize int64, monthly bool) ([]UserScore, int64, bool, error) {
	var nextPage int64
	var includesName bool

	scores, err := findScoresWithPositions(monthly)
	if err != nil {
		return scores, nextPage, includesName, err
	}

	scoresLen := int64(len(scores))
	skipScores := (pageNum - 1) * pageSize

	if scoresLen <= skipScores {
		return scores, nextPage, includesName, status.Error(codes.InvalidArgument, "invalid page number")
	}

	scoresFrom := skipScores

	var scoresTo int64
	if skipScores+pageSize > scoresLen {
		scoresTo = scoresLen
	} else {
		scoresTo = skipScores + pageSize
	}

	if scoresTo < scoresLen {
		nextPage = pageNum + 1
	}

	for i := range scores[:scoresTo] {
		if scores[i].Name == name {
			includesName = true

			break
		}
	}

	return scores[scoresFrom:scoresTo], nextPage, includesName, nil
}

func findScoresAround(name string, monthly bool) ([]UserScore, error) {
	aroundSize := 5

	scores, err := findScoresWithPositions(monthly)
	if err != nil {
		return scores, err
	}

	nameIndex := -1
	for i := range scores {
		if scores[i].Name == name {
			nameIndex = i

			break
		}
	}

	if nameIndex < 0 {
		return scores, errors.New("user not found by name")
	}

	var aroundFrom, aroundTo int

	if nameIndex-aroundSize > 0 {
		aroundFrom = nameIndex - aroundSize
	} else {
		aroundFrom = 0
	}

	if nameIndex+aroundSize < len(scores) {
		aroundTo = nameIndex + aroundSize
	} else {
		aroundTo = len(scores)
	}

	return scores[aroundFrom:aroundTo], nil
}

// func printScores(scores []UserScore) {
// 	for i := range scores {
// 		fmt.Println("==================")
// 		fmt.Println("Index: ", i)
// 		fmt.Println("Name:", scores[i].Name)
// 		fmt.Println("Points: ", scores[i].Points)
// 		fmt.Println("Position: ", scores[i].Position)
// 		fmt.Println("==================")
// 		fmt.Println("")
// 	}
// }

func assignPositions(scores []UserScore) {
	for i := range scores {
		if i == 0 {
			scores[i].Position = 1
			continue
		}

		if scores[i].Points == scores[i-1].Points {
			scores[i].Position = scores[i-1].Position
			continue
		}

		scores[i].Position = scores[i-1].Position + 1
	}
}

func seed() {
	db.Unscoped().Delete(&UserScore{})
	db.Create(&UserScore{Name: "John", Points: 10})
	db.Create(&UserScore{Name: "Alice", Points: 12})
	db.Create(&UserScore{Name: "Bob", Points: 1})
	db.Create(&UserScore{Name: "FPosition", Points: 5})
	db.Create(&UserScore{Name: "Jane", Points: 1})
	db.Create(&UserScore{Name: "Kane", Points: 11})
	db.Create(&UserScore{Name: "Maria", Points: 24})
	db.Create(&UserScore{Name: "Robert", Points: 13})
	db.Create(&UserScore{Name: "Michael", Points: 24})
	db.Create(&UserScore{Name: "Zed", Points: 8})
	db.Create(&UserScore{Name: "Anthony", Points: 4})
	db.Create(&UserScore{Name: "Anna", Points: 16})
	db.Create(&UserScore{Name: "Carl", Points: 19})
	db.Create(&UserScore{Name: "Fred", Points: 22})
	db.Create(&UserScore{Name: "Leo", Points: 32})
	db.Create(&UserScore{Name: "Peter", Points: 11})
	db.Create(&UserScore{Name: "Jet", Points: 11})
}
