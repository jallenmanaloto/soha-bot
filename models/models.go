package models

type Manhwa struct {
	PK         string
	SK         string
	ID         string
	Chapters   string
	Image      string
	Rating     string
	Scanlation string
	Title      string
	Url        string
}

type ServerManhwa struct {
	PK         string
	SK         string
	ID         string
	ChanId     string
	ServerId   string
	TitleId    string
	Title      string
	TitleCh    string
	TitleImage string
	TitleUrl   string
}
