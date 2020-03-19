package mongo

import (
	"fmt"
	"github.com/globalsign/mgo"
	//  "labix.org/v2/mgo/bson"
)

/**
 * 参考：http://www.cnblogs.com/shenguanpu/p/5318727.html
 * 参考：http://www.jyguagua.com/?p=3126
 */

var (
	session      *mgo.Session
	ip           = "127.0.0.1"
	port         = "27017"
	databaseName = "blogs"
	maxPoolSize  = "10"
	poolLimit    = 100
)

//func init(){
//	ip              = config.Get("mgo_ip")
//	port            = config.Get("mgo_port")
//	databaseName    = config.Get("mgo_databaseName")
//	maxPoolSize     = config.Get("mgo_maxPoolSize")
//	poolLimitStr   := config.Get("mgo_poolLimit")
//	poolLimit, _    = strconv.Atoi(poolLimitStr)
//}
func Session() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(ip + ":" + port + "?maxPoolSize=" + maxPoolSize)
		if err != nil {
			fmt.Println("出现错误")
			return nil
		}
		session.SetPoolLimit(poolLimit)
	}
	return session.Clone()
}

// 可以指定collection，database使用配置中的值
func MC(collection string) *mgo.Collection {
	session := Session()
	//defer func() {
	//	session.Close()
	//}()
	c := session.DB(databaseName).C(collection)
	return c
}

// 可以指定database和collection
func MDC(dbName string, collection string) *mgo.Collection {
	session := Session()
	//defer func() {
	//	session.Close()
	//}()
	return session.DB(dbName).C(collection)
}
