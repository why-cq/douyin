package pojo

type User struct {
	Id       int64  `json:"id,omitempty"`
	UserName string `json:"name,omitempty"`
	Password string `json:"pass_word,omitempty"`
}
