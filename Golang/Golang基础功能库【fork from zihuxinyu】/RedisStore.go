package Library

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
)

const SS_PREFIX = "ss:"
var StoragePtr *Storage
type redisUser struct {
	Name     string
	Password string

	Method   string

	EndDate  time.Time   //到期删除时间
	State 	 string        //账号状态 ok del stop


	Port     int
	Limit    string
}

type Storage struct {
	pool *redis.Pool
}

func NewStorage() *Storage {
	var server string
	server=":6379"
	pool := redis.NewPool(func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", server)
			return
		}, 3)
	return &Storage{pool}
}

func (s *Storage) Del(key string) (err error) {
	
	var conn = s.pool.Get()
	defer conn.Close()
	_, err = conn.Do("DEL", SS_PREFIX + key)
	return
	
	 
}
func (s *Storage) Get(key string) (user redisUser, err error) {
	fullkey := SS_PREFIX + key

	return s.get(fullkey)
}

func (s *Storage) get(fullkey string) (user redisUser, err error) {
	var data []byte
	var conn = s.pool.Get()
	defer conn.Close()
	data, err = redis.Bytes(conn.Do("GET", fullkey))
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &user)
	return
}

func (s *Storage) GetList() (userList []redisUser, err error) {
	//var data []byte
	var conn = s.pool.Get()
	defer conn.Close()
	key := "user*"
	data, err := redis.Strings(conn.Do("keys", SS_PREFIX+key))

	if err != nil {
		fmt.Printf("GetList()\n", err)
		return
	}

	userList = make([]redisUser, len(data))
	for k, v := range data {
		_user, _ := s.get(v)
		userList[k] = _user
	}

	return
}
func (s *Storage) Set(key string, user redisUser) (err error) {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	conn := s.pool.Get()
	defer conn.Close()
	_, err = conn.Do("SET", SS_PREFIX+key, data)
	return
}

func (s *Storage) IncrSize(key string, incr int) (score int64, err error) {
	var conn = s.pool.Get()
	defer conn.Close()
	score, err = redis.Int64(conn.Do("INCRBY", SS_PREFIX+key, incr))
	return
}

func (s *Storage) GetSize(key string) (score int64, err error) {
	var conn = s.pool.Get()
	defer conn.Close()
	score, err = redis.Int64(conn.Do("GET", SS_PREFIX+key))
	return
}

func (s *Storage) ZincrbySize(key, member string, incr int) (err error) {
	var conn = s.pool.Get()
	defer conn.Close()
	var score int64
	var year, month, day int
	var real_key string

	now := time.Now()
	year = now.Year()
	month = int(now.Month())
	day = now.Day()

	// store year
	real_key = fmt.Sprintf("%s%s:%d", SS_PREFIX, key, year)
	score, err = redis.Int64(conn.Do("ZINCRBY", real_key, incr, member))
	if score < 0 || err != nil {
		conn.Do("ZADD", real_key, incr, member)
	}
	// store year:month
	real_key = fmt.Sprintf("%s%s:%d:%d", SS_PREFIX, key, year, month)
	score, err = redis.Int64(conn.Do("ZINCRBY", real_key, incr, member))
	if score < 0 || err != nil {
		conn.Do("ZADD", real_key, incr, member)
	}
	// store year:month:day
	real_key = fmt.Sprintf("%s%s:%d:%d:%d", SS_PREFIX, key, year, month, day)
	score, err = redis.Int64(conn.Do("ZINCRBY", real_key, incr, member))
	if score < 0 || err != nil {
		conn.Do("ZADD", real_key, incr, member)
	}

	// store year total
	real_key = fmt.Sprintf("%s%s:%d", SS_PREFIX, key, year)
	score, err = redis.Int64(conn.Do("ZINCRBY", real_key, incr, "total"))
	if score < 0 || err != nil {
		conn.Do("ZADD", real_key, incr, "total")
	}
	// store year:month total
	real_key = fmt.Sprintf("%s%s:%d:%d", SS_PREFIX, key, year, month)
	score, err = redis.Int64(conn.Do("ZINCRBY", real_key, incr, "total"))
	if score < 0 || err != nil {
		conn.Do("ZADD", real_key, incr, "total")
	}
	// store year:month:day total
	real_key = fmt.Sprintf("%s%s:%d:%d:%d", SS_PREFIX, key, year, month, day)
	score, err = redis.Int64(conn.Do("ZINCRBY", real_key, incr, "total"))
	if score < 0 || err != nil {
		conn.Do("ZADD", real_key, incr, "total")
	}
	return
}
