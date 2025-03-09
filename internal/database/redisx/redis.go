package redisx

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var RedisClient *redis.Client

//var ctx = context.Background()

func InitRedisConn(ctx context.Context, options redis.Options) *redis.Client {
	RedisClient = redis.NewClient(&options)

	// 验证连接是否成功
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return RedisClient
}

// 设置 bitmap 使用
// 语法：setbit like_id1 100 1 设置 ID 为 100 为 1
func GiveLike(ctx context.Context, keys string, userID int64) (bool, error) {
	result, err := RedisClient.GetBit(ctx, keys, userID-1).Result()
	if err != nil {
		return false, err
	}

	if result == 1 {
		return true, nil
	}

	_, err = RedisClient.SetBit(ctx, keys, userID-1, 1).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

// 查询 bitmap 使用
func GiveLikeSelect(ctx context.Context, keys string, userID int64) (bool, error) {
	result, err := RedisClient.GetBit(ctx, keys, userID-1).Result()
	if err != nil {
		return false, err
	}

	if result == 1 {
		return true, nil
	}

	return false, err
}

// https://github.com/bsm/redislock
// redis加锁
func Lock(ctx context.Context, key string) bool {
	var mutex sync.Locker
	mutex.Lock()
	defer mutex.Unlock()
	result, err := RedisClient.SetNX(ctx, key, 1, 10*time.Second).Result()
	if err != nil {
		panic(err)
	}

	return result
}

// 释放锁
func UnLock(ctx context.Context, key string) int64 {
	nums, err := RedisClient.Del(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return nums
}

// set 设置 key 值
func Set(ctx context.Context, key, value string) error {
	_, err := RedisClient.Set(ctx, key, value, 0).Result()
	return err
}

// SetEX 设置 key的值并指定过期时间
func SetEX(ctx context.Context, key, value string, ex time.Duration) error {
	_, err := RedisClient.Set(ctx, key, value, ex).Result()

	return err
}

// Get 获取 key的值
func Get(ctx context.Context, key string) (string, error) {
	result, err := RedisClient.Get(ctx, key).Result()
	return result, err
}

// GetSet 设置新值获取旧值
func GetSet(ctx context.Context, key, value string) (string, error) {
	result, err := RedisClient.GetSet(ctx, key, value).Result()
	return result, err
}

// Incr key值每次加一 并返回新值
func Incr(ctx context.Context, key string) (int64, error) {
	result, err := RedisClient.Incr(ctx, key).Result()
	return result, err
}

// IncrBy key值每次加指定数值 并返回新值
func IncrBy(ctx context.Context, key string, incr int64) (int64, error) {
	result, err := RedisClient.IncrBy(ctx, key, incr).Result()
	return result, err
}

// IncrByFloat key值每次加指定浮点型数值 并返回新值
func IncrFloatBy(ctx context.Context, key string, incrFloat float64) (float64, error) {
	result, err := RedisClient.IncrByFloat(ctx, key, incrFloat).Result()
	return result, err
}

// Decr key值每次递减 1 并返回新值
func Decr(ctx context.Context, key string) (int64, error) {
	result, err := RedisClient.Decr(ctx, key).Result()
	return result, err
}

// DecrBy key值每次递减指定数值 并返回新值
func DecrBy(ctx context.Context, key string, decr int64) (int64, error) {
	result, err := RedisClient.Decr(ctx, key).Result()
	return result, err
}

// Del 删除 key
func Del(ctx context.Context, key string) (int64, error) {
	result, err := RedisClient.Del(ctx, key).Result()
	return result, err
}

// Expire 设置 key的过期时间
func Expire(ctx context.Context, key string, ex time.Duration) (bool, error) {
	result, err := RedisClient.Expire(ctx, key, ex).Result()
	return result, err
}

/*------------------------------------ list 操作 ------------------------------------*/

// LPush 从列表左边插入数据，并返回列表长度
func LPush(ctx context.Context, key string, date ...interface{}) (int64, error) {
	return RedisClient.LPush(ctx, key, date).Result()
}

// RPush 从列表右边插入数据，并返回列表长度
func RPush(ctx context.Context, key string, date ...interface{}) (int64, error) {
	return RedisClient.RPush(ctx, key, date).Result()
}

// LPop 从列表左边删除第一个数据，并返回删除的数据
func LPop(ctx context.Context, key string) (string, error) {
	return RedisClient.LPop(ctx, key).Result()
}

// RPop 从列表右边删除第一个数据，并返回删除的数据
func RPop(ctx context.Context, key string) (string, error) {
	return RedisClient.RPop(ctx, key).Result()
}

// LIndex 根据索引坐标，查询列表中的数据
func LIndex(ctx context.Context, key string, index int64) (string, error) {
	return RedisClient.LIndex(ctx, key, index).Result()
}

// LLen 返回列表长度
func LLen(ctx context.Context, key string) (int64, error) {
	return RedisClient.LLen(ctx, key).Result()
}

// LRange 返回列表的一个范围内的数据，也可以返回全部数据
func LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return RedisClient.LRange(ctx, key, start, stop).Result()
}

// LRem 从列表左边开始，删除元素data， 如果出现重复元素，仅删除 count次
func LRem(ctx context.Context, key string, count int64, data interface{}) (int64, error) {
	return RedisClient.LRem(ctx, key, count, data).Result()
}

// LInsert 在列表中 pivot 元素的后面插入 data
func LInsert(ctx context.Context, key string, pivot int64, data interface{}) (int64, error) {
	return RedisClient.LInsert(ctx, key, "after", pivot, data).Result()
}

/*------------------------------------ set 操作 ------------------------------------*/

// SAdd 添加元素到集合中
func SAdd(ctx context.Context, key string, data interface{}) (int64, error) {
	return RedisClient.SAdd(ctx, key, data).Result()
}

// SCard 获取集合元素个数
func SCard(ctx context.Context, key string) (int64, error) {
	return RedisClient.SCard(ctx, key).Result()
}

// SIsMember 判断元素是否在集合中
func SIsMember(ctx context.Context, key string, data interface{}) (bool, error) {
	return RedisClient.SIsMember(ctx, key, data).Result()
}

// SMembers 获取集合所有元素
func SMembers(ctx context.Context, key string) ([]string, error) {
	return RedisClient.SMembers(ctx, key).Result()
}

// SRem 删除 key集合中的 data元素
func SRem(ctx context.Context, key string, data ...interface{}) (int64, error) {
	return RedisClient.SRem(ctx, key, data).Result()
}

// SPopN 随机返回集合中的 count个元素，并且删除这些元素
func SPopN(ctx context.Context, key string, count int64) ([]string, error) {
	return RedisClient.SPopN(ctx, key, count).Result()
}

/*------------------------------------ hash 操作 ------------------------------------*/

// HSet 根据 key和 field字段设置，field字段的值
func HSet(ctx context.Context, key, field, value string) (int64, error) {
	return RedisClient.HSet(ctx, key, field, value).Result()
}

// HGet 根据 key和 field字段，查询field字段的值
func HGet(ctx context.Context, key, field string) (string, error) {
	return RedisClient.HGet(ctx, key, field).Result()
}

// HMGet 根据key和多个字段名，批量查询多个 hash字段值
func HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	return RedisClient.HMGet(ctx, key, fields...).Result()
}

// HGetAll 根据 key查询所有字段和值
func HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return RedisClient.HGetAll(ctx, key).Result()
}

// HKeys 根据 key返回所有字段名
func HKeys(ctx context.Context, key string) ([]string, error) {
	return RedisClient.HKeys(ctx, key).Result()
}

// HLen 根据 key，查询hash的字段数量
func HLen(ctx context.Context, key string) (int64, error) {
	return RedisClient.HLen(ctx, key).Result()
}

// HMSet 根据 key和多个字段名和字段值，批量设置 hash字段值
func HMSet(ctx context.Context, key string, data map[string]interface{}) (bool, error) {
	return RedisClient.HMSet(ctx, key, data).Result()
}

// HSetNX 如果 field字段不存在，则设置 hash字段值
func HSetNX(ctx context.Context, key, field string, value interface{}) (bool, error) {
	return RedisClient.HSetNX(ctx, key, field, value).Result()
}

// HDel 根据 key和字段名，删除 hash字段，支持批量删除
func HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	return RedisClient.HDel(ctx, key, fields...).Result()
}

// HExists 检测 hash字段名是否存在
func HExists(ctx context.Context, key, field string) (bool, error) {
	return RedisClient.HExists(ctx, key, field).Result()
}
