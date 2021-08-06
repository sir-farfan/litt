package model

type Member struct {
	ID   int
	Name string
	Type int
}

type Tag struct {
	ID  int
	Tag string
}

type MemberTags struct {
	MemberID int
	TagID    int
}
