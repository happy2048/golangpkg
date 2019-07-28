package redis
import (
	redis "github.com/gomodule/redigo/redis"
)

func (rdb *RedisDBImpl) HashSet(key,field,val string) (bool,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("HSET",key,field,val))
}

func (rdb *RedisDBImpl) HashGet(key,field string) (string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("HGET",key,field))
}

func (rdb *RedisDBImpl) HashMSet(hkey string,fieldsMap  map[string]string) (string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("HMSET",redis.Args{}.Add(hkey).AddFlat(fieldsMap)...))
}

func (rdb *RedisDBImpl) HashMGet(hkey string,fields []string) ([]string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("HMGET",redis.Args{}.Add(hkey).AddFlat(fields)...))
}
func (rdb *RedisDBImpl) HashHDel(key string,fields []string) (bool,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("HDEL",redis.Args{}.Add(key).AddFlat(fields)...)) 
}
func (rdb *RedisDBImpl) HashFieldExist(key string,field string) (bool,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("HEXISTS",redis.Args{}.Add(key).Add(field)...))
}

func (rdb *RedisDBImpl) HashGetAll(key string) (map[string]string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.StringMap(conn.Do("HGETALL",key))
}
func (rdb *RedisDBImpl) HashFields(key string) ([]string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("HKEYS",key))
}

