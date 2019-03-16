package nosql

import (
	"github.com/garyburd/redigo/redis"
	"github.com/astaxie/beego/config"
	"zdy.local/utils/nlogs"
	"zdy.local/sql/mysql/mysqldriver"
)

type RedisConn struct {
	Dil redis.Conn
	err error
}

func (re *RedisConn)InitRedis()  {
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	dbHost := cfg.String("redis::host")
	dbPort := cfg.String("redis::port")
	if err != nil {
		nlogs.ConsoleLogs.Error(err.Error())
		panic(err)
	}
	re.Dil,re.err = redis.Dial("tcp", dbHost+":"+dbPort)
	if re.err != nil{
		nlogs.FileLogs.Error(re.err.Error())
	}
	re.Dil.Do("MULTI")
}

func (re *RedisConn)SetMessage(key string,value mysqldriver.Image){
	rep,err := re.Dil.Do("HMSET",redis.Args{}.Add(key).AddFlat(&value)...)
	if err != nil{
		nlogs.ConsoleLogs.Debug(err.Error())
	}else {
		nlogs.ConsoleLogs.Debug(nlogs.Fmt2String(rep))
	}
}

func (re *RedisConn)GetMessage(key string) (mysqldriver.Image,error){
	var img mysqldriver.Image
	val, err := redis.Strings(re.Dil.Do("KEYS", "*"+key))
	if err!=nil{
		nlogs.ConsoleLogs.Debug("not find key",key,err.Error())
		return img,err
	}
	v, err := redis.Values(re.Dil.Do("HGETALL", val))
	if err!=nil{
		nlogs.ConsoleLogs.Debug(err.Error())
		return img,err
	}
	if err := redis.ScanStruct(v, &img); err != nil{
		nlogs.ConsoleLogs.Info(nlogs.Fmt2String(img))
		return img,err
	}
	return img,nil
}

func (re *RedisConn)Close(){
	re.Dil.Do("EXEC")
	re.Dil.Close()
}