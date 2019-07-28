package redis
import(
	"time"
	redis "github.com/gomodule/redigo/redis"
)

type RedisDB interface {
	// string operation: get
	StringGet(key string) (string,error)
    // string operation: set
	StringSet(key,val string) (string,error)
	// string operation: set with expiration time
	StringSetWithExpTime(key,value string,ex int) (string,error)
    // hash operation: hash set
	HashSet(key,feild,value string) (bool,error)
    // hash operation: hash get 
	HashGet(key,feild string) (string,error)
    // hash operation: hash mset
	HashMSet(key string,value map[string]string) (string,error)
	// hash operation: hash mget
	HashMGet(key string,fields []string) ([]string,error)
	// hash operation: hash del
	HashHDel(key string,fields []string) (bool,error)
    // hash operation: hash exists
	HashFieldExist(key string,field string) (bool,error)
    // hash operation: hash vals
	HashGetAll(key string) (map[string]string,error)
    // hash operation: hash keys
    HashFields(key string) ([]string,error)
	// list operation: list lpush
	ListLPush(key,value string) (int,error)
	// list operation: list rpush
   	ListRPush(key,value string) (int,error)
	// list operation: list lset
	ListLSet(key,value string,index int) (string,error)
	// list operation: list lrange
	ListLRange(key string,start,end int) ([]string,error)
	// List operation: list lpop
	ListLPop(key string) (string,error)
	// List operation: list rpop
	ListRPop(key string) (string,error)
	// list operation: list linsert
	ListLInsert(key,value,inserType,pivot string) (int,error)
	// List operation: list llen
	ListLLen(key string) (int,error)
	// List operation: value of index 
	ListQueryValueByIndex(key string,index int) (string,error)
	// delete redis key
	DelKey(key string) (bool,error)
	// subscribe operation: subscribe
	Subscribe(redisChannel string,transferStatus func(string))
	// publish operation: publish
	Publish(channel,data string) (bool, error)
}

type RedisDBImpl struct {
	pool  *redis.Pool
}

type RedisConfigArgs struct {
	Protocol string   // protocol: tcp,udp
	Endpoint string   // endpoint: format is "ip:port"
	DB		 string   // select db 
	Password string   // db password
	MaxIdle  int 	 // MaxIdle: max idle connection
	IdleTimeout int  // second
}

func NewRedisDB(rconf RedisConfigArgs) (RedisDB,error) {
	var db RedisDB
    if err := RedisTestConnect(rconf.Endpoint);err != nil {
		return nil,err
	}
	redisDBImpl := &RedisDBImpl{
		pool: &redis.Pool{
				MaxIdle: rconf.MaxIdle,
				IdleTimeout: time.Duration(rconf.IdleTimeout) * time.Second,
				Dial: func() (redis.Conn,error) {
					conn,err := redis.Dial(rconf.Protocol,rconf.Endpoint); if err != nil {
						return nil,err
					}
					switch {
					case rconf.Password != "":
						if _,err := conn.Do("AUTH",rconf.Password); err != nil {
							conn.Close()
							return nil,err
						}
					case rconf.DB != "":
						if _,err := conn.Do("SELECT",rconf.DB); err != nil {
							conn.Close()
							return nil,err
						}
					
					}
					return  conn,nil
				},
		},
	}
	db = redisDBImpl
	return db,nil
}

func RedisTestConnect(addr string) (error) {
	c,err := redis.Dial("tcp",addr)
	if err != nil {
		return err
	}
	defer c.Close()
	return nil
}

func (rdb *RedisDBImpl) DelKey(key string) (bool,error) {
	conn := rdb.pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("DEL",key))
}

