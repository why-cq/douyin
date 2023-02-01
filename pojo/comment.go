package pojo

// Comment
type Comment struct {
	Id         int64  `json:"id,omitempty"`
	UserId     int64  //评论用户id
	VideoId    int64  //视频id
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	Cancel     int32  //取消评论为1，发布评论为0
}
