package pojo

// Follow 用户关系结构，对应用户关系表。
type Follow struct {
	Id         int64
	UserId     int64 //当前用户id
	FollowerId int64 //关注的用户id
	Cancel     int8  //是否关注(默认0为关注)
}
