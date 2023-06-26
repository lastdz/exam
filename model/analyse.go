package model

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/wire"
	"strconv"
)

func Analyse(examid string) wire.Resp {
	fmt.Println(examid)
	var equestions []wire.ExamQuestions
	equestions = dao.QueryExamQuestion(examid)

	for _, t := range equestions {
		fmt.Println(t)
	}
	var questions []wire.Questions
	for _, equestion := range equestions {
		q := dao.QueryQuestion(equestion)

		questions = append(questions, q)
	}
	diff := make([]int, 1000000)
	types := make(map[string]int)
	for _, q := range questions {
		diff[q.Difficulty]++
		types[q.Type]++
	}
	resp := wire.Resp{}
	for i := 0; i < 1000000; i++ {
		if diff[i] != 0 {
			tmp := "难度:" + strconv.Itoa(i) + "数量:" + strconv.Itoa(diff[i]) + "    "
			resp.Difficulty += tmp
		}
	}
	for i, v := range types {
		if v != 0 {
			tmp := "类型:" + i + "数量:" + strconv.Itoa(v) + "    "
			resp.Type += tmp
		}
	}
	for _, q := range questions {
		resp.QA += "题目:" + q.Content + "答案:" + q.Answer + "    "
	}
	return resp
}
