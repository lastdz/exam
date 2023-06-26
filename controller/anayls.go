package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
import "github.com/jinzhu/gorm"

type ExamQuestions struct {
	Examquestionid int
	Examid         int
	Questionid     int
	Score          float64
}

type Questions struct {
	Questionid     int
	Content        string
	Type           string
	Answer         string
	Difficulty     int
	Courseid       int
	Chapterid      int
	Errorrate      float64
	Commonmistakes string
	Teacherid      int
}

func Analys(c *gin.Context) {
	examid := c.Query("examid")

	db, err := gorm.Open("mysql", "root:123456@/online?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	var equestions []ExamQuestions
	db.Raw("select * from examquestion where examid = ?", examid).Scan(&equestions)

	for _, t := range equestions {
		fmt.Println(t)
	}
	var questions []Questions
	for _, equestion := range equestions {
		q := Questions{}
		db.Raw("select difficulty from examquestion where questionid = ?", equestion.Questionid).Scan(&q.Difficulty)
		db.Raw("select type from examquestion where questionid = ?", equestion.Questionid).Scan(&q.Type)
		fmt.Println(q)
		questions = append(questions, q)
	}
}
