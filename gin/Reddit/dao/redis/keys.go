package redis

// redis key
// redis key 注意使用命名空间的方式，方便查询和拆分
const (
	Prefix                 = "bluebell:"
	KeyPostTimeZSet        = "post:time"   // ZSet: 帖子及发帖时间
	KeyPostScoreZSet       = "post:score"  // ZSet: 帖子及投票的分数
	KeyPostVotedZSetPrefix = "post:voted:" // ZSet: 记录用户及投票类型，参数是 post_id
	KeyCommunitySetPrefix  = "community:"  // set: 保存每个分区下帖子的 id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
