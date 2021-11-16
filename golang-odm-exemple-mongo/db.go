package database

import (
	"crypto/tls"
	"log"
	"net"
	"time"

	mgo "gopkg.in/mgo.v2"
)

//MongoSession ...
var mongoSession *mgo.Session
var dbName string

//InitDB ...
func initDB(config ConfigDB) error {
	if len(config.Servers) > 0 {

		var info mgo.DialInfo
		info.Database = config.Database
		info.Username = config.User
		info.Password = config.Password
		info.Addrs = config.Servers // Get HOST + PORT
		info.Timeout = 30 * time.Second

		log.Println(info.Database)
		log.Println(info.Addrs)

		if config.UseSSL == true {
			info.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
				return tls.Dial("tcp", addr.String(), &tls.Config{InsecureSkipVerify: true})
			}
		}
		session, err := mgo.DialWithInfo(&info)
		if err != nil {
			return err
		}
		// Optional. Switch the session to a monotonic behavior.
		session.SetMode(mgo.Monotonic, true)
		dbName = config.Database
		mongoSession = session
		return nil
	}
	return nil
}

type JSONDoc map[string]interface{}
