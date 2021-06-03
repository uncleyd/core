/*
* @Author: haodaquan
* @Date:   2017-06-20 09:44:44
* @Last Modified by:   haodaquan
* @Last Modified time: 2017-06-21 12:21:37
 */

package rdmodels

import (
	"github.com/go-redis/redis"
	"github.com/goinggo/mapstructure"
	"reflect"
	"time"
	"core/config"
)

var RdClient *redis.Client

func Init() {
	RedisInit()
}

func RedisInit() {
	var redisConfig = config.Get().Redis
	rdhost := redisConfig[0].Host
	rdport := redisConfig[0].Port
	rdpassword := redisConfig[0].Password
	rdidx := redisConfig[0].Index
	if rdpassword != "" {
		RdClient = redis.NewClient(&redis.Options{
			Addr:         rdhost + ":" + rdport,
			Password:     rdpassword, // no password set
			DB:           rdidx,      // use default DB
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			PoolSize:     10,
			PoolTimeout:  30 * time.Second,
			IdleTimeout:  500 * time.Millisecond,
		})
	} else {
		RdClient = redis.NewClient(&redis.Options{
			Addr: rdhost + ":" + rdport,
			//Password: rdpassword, 	// no password set
			DB:           rdidx, // use default DB
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			PoolSize:     10,
			PoolTimeout:  30 * time.Second,
			IdleTimeout:  500 * time.Millisecond,
		})
	}
}

func Struct2Map(obj interface{}) map[string]interface{} {
	obj_v := reflect.ValueOf(obj)
	v := obj_v.Elem()
	typeOfType := v.Type()
	var data = make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		data[typeOfType.Field(i).Name] = field.Interface()
	}
	return data
}

func Map2Struct(mapInstance interface{}, pointer interface{}) error {
	return mapstructure.Decode(mapInstance, pointer)
}
