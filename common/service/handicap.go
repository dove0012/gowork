package service

import (
	"gopkg.in/mgo.v2/bson"
	"gowork/common/utils/convert"
	"gowork/common/utils/log"
	"gowork/common/utils/mgodb"
)

func GetHandicapById(han_id int64, result interface{}) {
	err := mgodb.Single().S.DB(DB_DM_DATA).C(C_HANDICAP).Find(bson.M{"han_id": han_id}).One(result)
	log.Error2Exit(err, "GetHandicapById["+convert.ToStr(han_id)+"] error")
}
