package wire

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

type Resp struct {
	Difficulty string
	Type       string

	QA string
}
