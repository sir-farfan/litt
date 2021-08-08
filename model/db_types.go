package model

type Member struct {
	ID   int
	Name string
	Type int
}

type Tag struct {
	ID  int    `json:"id"`
	Tag string `json:"tag"`
}

type MemberTags struct {
	MemberID int
	TagID    int
}
