package pojo

type User struct {
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"pass_word,omitempty"`
}
