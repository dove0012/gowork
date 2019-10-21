package mgodb

import (
	"gopkg.in/mgo.v2"
	"gowork/common/utils/config"
	"gowork/common/utils/log"
	"sync"
)

type Mgodb struct {
	S *mgo.Session
}

var db = &Mgodb{}
var l sync.Mutex

func Single() *Mgodb {
	if db.S == nil {
		l.Lock()
		defer l.Unlock()
		if db.S == nil {
			var err error
			db.S, err = mgo.Dial(config.C.GetStringByDefault("mongodb_url"))
			log.Error2Exit(err, "NewMongodbSession Dial error")
		}
	}
	return db
}

func Close() {
	if db.S != nil {
		db.S.Close()
		db.S = nil
	}
}
