package redis
import (
	"strconv"
	redis "github.com/gomodule/redigo/redis"
)

func (rdb *RedisDBImpl) StringGet(key string) (string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("GET",key))
}

func (rdb *RedisDBImpl) StringSet(key,val string) (string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	// return "OK <nil>"
	status,err := redis.String(conn.Do("SET",key,val))
	return status,err
}

func (rdb *RedisDBImpl) StringSetWithExpTime(key,val string,ex int) (string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	// return "OK <nil>"
	return redis.String(conn.Do("SET",key,val,"EX", strconv.Itoa(ex)))
}

