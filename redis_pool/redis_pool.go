package redis_pool

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

//redis pool
var once sync.Once

// MyRedis redis配置项
type redisConPool struct {
	conf *Conf
	// con  redis.Conn
	pool *redis.Pool
	mtx  sync.Mutex
	once sync.Once
	err  error
}

var redisConrtol *redisConPool = new(redisConPool)

// 封装 redis 实例，提供获取
func GetRedisInstance() *redisConPool {
	return redisConrtol
}

type Conf struct {
	Addr     string
	Password string

	Proto string

	Host string
	Port int

	MaxActive   int // 池在给定时间分配的最大连接数。当为零时，池中的连接数没有限制。
	MaxIdle     int // 池中空闲连接的最大数目。
	IdleTimeout time.Duration
	Wait        bool
	isLog       bool
	Options     []redis.DialOption
}

func InitRedisConfig(conf *Conf, options ...redis.DialOption) {

	ops := []redis.DialOption{
		redis.DialPassword(conf.Password),
	}
	conf.Options = append(ops, options...)

	if conf.Proto == "" {
		conf.Proto = "tcp"
	}
	if conf.MaxActive == 0 {
		conf.MaxActive = 100
	}
	if conf.MaxIdle == 0 {
		conf.MaxIdle = 100
	}
	if conf.IdleTimeout == 0 {
		conf.IdleTimeout = time.Second * 5
	}
	if conf.Addr == "" {
		conf.Addr = fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	}

	once.Do(func() {
		redisConrtol = &redisConPool{
			conf: conf,
		}
	})

}

//获取连接
func (rp *redisConPool) GetRedisClient() redis.Conn {
	if rp.conf.Proto == "" {
		rp.conf.Proto = "tcp"
	}
	if rp.conf.MaxActive == 0 {
		rp.conf.MaxActive = 100
	}
	if rp.conf.MaxIdle == 0 {
		rp.conf.MaxIdle = 100
	}
	if rp.conf.IdleTimeout == 0 {
		rp.conf.IdleTimeout = time.Second * 5
	}
	if rp.conf.Addr == "" {
		rp.conf.Addr = fmt.Sprintf("%s:%d", rp.conf.Host, rp.conf.Port)
	}

	rp.mtx.Lock()
	if rp.pool == nil { // 创建连接
		rp.pool = &redis.Pool{
			MaxIdle:     rp.conf.MaxIdle,
			MaxActive:   rp.conf.MaxActive,
			IdleTimeout: rp.conf.IdleTimeout,
			Wait:        rp.conf.Wait,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial(rp.conf.Proto, rp.conf.Addr, rp.conf.Options...)
				if err != nil {
					panic(err)
				}
				return c, nil
			},
		}
	}
	rp.mtx.Unlock()

	con := rp.pool.Get()
	if rp.conf.isLog {
		fmt.Println("ActiveCount:", rp.pool.ActiveCount())
	}

	return con
}

func (rp *redisConPool) Destory() {
	rp.mtx.Lock()
	defer rp.mtx.Unlock()

	if rp.pool != nil {
		rp.pool.Close()
	}
}

func (rp *redisConPool) Ping() bool {
	return rp.ping(rp.GetRedisClient())
}

func (rp *redisConPool) ping(con redis.Conn) bool {
	if con == nil {
		return false
	}

	_, err := con.Do("PING")
	if err != nil {
		fmt.Printf("ping redis error: %s", err)
		return false
	}
	fmt.Printf("ping redis success")
	return true
}

func (rp *redisConPool) Do(command string, params ...interface{}) (interface{}, error) {
	con := rp.GetRedisClient()
	defer con.Close()
	return con.Do(command, params...)
}

//获取数据
func (rp *redisConPool) GetValue(key string, value interface{}) (err error) {
	repy, err := rp.Do("GET", key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return rp.decodeValue(repy, value)
}

//转换内容
func (rp *redisConPool) decodeValue(in, out interface{}) (err error) {
	if in == nil {
		return fmt.Errorf("not fond")
	}

	var reply string
	switch t := in.(type) {
	case []byte:
		reply = string(t)
	default:
		return fmt.Errorf("decodeValue err in type not find:%v", t)
	}

	switch o := out.(type) {
	case *string: // string类型
		*o = reply
		return nil
	case *int32:
		i64, err := strconv.ParseInt(reply, 10, 0)
		*o = int32(i64)
		return err
	case *bool:
		b, err := strconv.ParseBool(string(reply))
		*o = b
		return err
	case *int:
		i64, err := strconv.ParseInt(reply, 10, 0)
		*o = int(i64)
		return err
	case *int8:
		i64, err := strconv.ParseInt(reply, 10, 0)
		*o = int8(i64)
		return err
	case *int16:
		i64, err := strconv.ParseInt(reply, 10, 0)
		*o = int16(i64)
		return err
	case *int64:
		i64, err := strconv.ParseInt(string(reply), 10, 64)
		*o = int64(i64)
		return err
	case *uint:
		i64, err := strconv.ParseUint(reply, 10, 0)
		*o = uint(i64)
		return err
	case *uint8:
		i64, err := strconv.ParseUint(reply, 10, 0)
		*o = uint8(i64)
		return err
	case *uint16:
		i64, err := strconv.ParseUint(reply, 10, 0)
		*o = uint16(i64)
		return err
	case *uint32:
		i64, err := strconv.ParseInt(string(reply), 10, 0)
		*o = uint32(i64)
		return err
	case *uint64:
		i64, err := strconv.ParseUint(reply, 10, 64)
		*o = uint64(i64)
		return err
	case *float32:
		f64, err := strconv.ParseFloat(string(reply), 32)
		*o = float32(f64)
		return err
	case *float64: // 基础类型
		f64, err := strconv.ParseFloat(string(reply), 64)
		*o = float64(f64)
		return err
	default:
		return json.Unmarshal([]byte(reply), out)

	}

}
