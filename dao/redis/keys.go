package redis


// redis key--命名空间
const (
	KeyPostTimeZSet = "reddit:post:time"	// 帖子以发帖时间为分数的zset
	KeyPostScoreZSet = "reddit:post:score"	// 帖子以投票分数为分数的zset
	KeyPostVotedZSetPF = "reddit:post:voted:"	// 记录用户投票类型的zset；参数是post id
)

////