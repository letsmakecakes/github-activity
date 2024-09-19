package model

type Actor struct {
	ID           string `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type Event struct {
	ID    int64  `json:"id"`
	Type  string `json:"type"`
	Actor Actor  `json:"actor"`
}
