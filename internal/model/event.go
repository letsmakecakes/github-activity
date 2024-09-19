package model

type Actor struct {
	ID string `json:"id"`
}

type Event struct {
	ID    int64  `json:"id`
	Type  string `json:"type"`
	Actor Actor  `json:"actor"`
}
