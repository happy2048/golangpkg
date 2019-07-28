package redis
import(
    redis "github.com/gomodule/redigo/redis"
)

func (rdb *RedisDBImpl) ListLPush(key,value string) (int,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	// return "OK <nil>"
	return redis.Int(conn.Do("LPUSH",key,value))
}

func (rdb *RedisDBImpl) ListRPush(key,value string) (int,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("RPUSH",key,value))
} 

func (rdb *RedisDBImpl) ListLSet(key,value string,index int) (string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("LSET",key,index,value))
}

func (rdb *RedisDBImpl) ListLRange(key string,start,end int) ([]string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.Strings(conn.Do("LRANGE",key,start,end))
}

func (rdb *RedisDBImpl) ListLPop(key string) (string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("LPOP",key))
}

func (rdb *RedisDBImpl) ListRPop(key string) (string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("RPOP",key))

}

func (rdb *RedisDBImpl) ListLInsert(key,value,inserType,pivot string) (int,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("LINSERT",key,inserType,pivot,value))

}

func (rdb *RedisDBImpl) ListLLen(key string) (int,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("LLEN",key))
}

func (rdb *RedisDBImpl) ListQueryValueByIndex(key string,index int) (string,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.String(conn.Do("LINDEX",key,index))
}


