package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/wire"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var _ mysql.Config

func Init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@/online?charset=utf8&parseTime=True&loc=Local")
	fmt.Println("!!!!", err)
}
func QueryExamQuestion(examid string) []wire.ExamQuestions {
	var equestions []wire.ExamQuestions
	fmt.Println(examid)
	db.Raw("select * from examquestion where examid = ?", examid).Scan(&equestions)
	return equestions
}
func QueryQuestion(equestion wire.ExamQuestions) wire.Questions {
	q := wire.Questions{}
	p := wire.Questions{}
	db.Raw("select difficulty from question where questionid = ?", equestion.Questionid).Scan(&q)
	db.Raw("select type from question where questionid = ?", equestion.Questionid).Scan(&p)
	q.Type = p.Type
	db.Raw("select answer from question where questionid = ?", equestion.Questionid).Scan(&p)
	q.Type = p.Type
	db.Raw("select content from question where questionid = ?", equestion.Questionid).Scan(&p)
	q.Content = p.Content
	return q
}
