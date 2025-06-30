package constants

const (
	MysqlDefaultDsn = "root:pass@tcp(127.0.0.1:3306)/db0?charset=utf8&parseTime=True&loc=Local"
)

const (
	CommentsTableName  = "comments"
	FollowsTableName   = "follows"
	FavoritesTableName = "favorites"
	UsersTableName     = "users"
	VideosTableName    = "videos"
	MessagesTableName  = "messages"

	VideoFeedCount       = 30
	FavoriteActionType   = 1
	UnFavoriteActionType = 2

	MinioVideoBucketName = "video_bucket"
	MinioImageBucketName = "image_bucket"

	TestSignature       = "测试用户签名"
	TestAvatar          = "avatar/test1.jpg"
	TestBackgroundImage = "background/test1.png"
)
