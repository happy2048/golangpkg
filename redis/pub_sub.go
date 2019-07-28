package redis

import (
	"fmt"
	redis "github.com/gomodule/redigo/redis"
)

func (rdb *RedisDBImpl) Subscribe(channel string,transferStatus func(string)) {
	conn := rdb.pool.Get()
	defer conn.Close()
	psc := redis.PubSubConn{conn}
	psc.Subscribe(channel)
	for {
		switch v := psc.Receive().(type) {
			case redis.Message:
				transferStatus(string(v.Data))
			case error:
				fmt.Println(v)
				return
		}
	}

}

func (rdb *RedisDBImpl) Publish(channel,data string) (bool,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("PUBLISH",channel,data))
}
