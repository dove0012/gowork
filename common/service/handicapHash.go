package service

import (
	"common/model"
	"common/utils/convert"
	"common/utils/log"
	"common/utils/mgodb"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

func GetHandicapHashById(han_id int64, result interface{}) {
	err := mgodb.Single().S.DB(DB_DM_ADMIN).C(C_HANDICAP_HASH).Find(bson.M{"han_id": han_id}).One(result)
	log.Error2Exit(err, "GetHandicapHashById["+convert.ToStr(han_id)+"] error")
}

func HandicapHashIsValid(handicap *model.Handicap) bool {
	handicapHash := &model.HandicapHash{}
	GetHandicapHashById(handicap.Han_id, handicapHash)
	str := fmt.Sprintf("%d%d%d%d%s%d",
		handicap.Han_id,
		handicap.Han_type,
		handicap.Return_rate,
		handicap.Result,
		handicap.Resulttime.Format("2006-01-02 15:04:05"),
		handicap.Status,
	)
	return HashString(str) == handicapHash.Hash1
}

func HashString(str string) string {
	h := hmac.New(sha1.New, []byte(model.HANDICAP_HASH_KEY))
	h.Write([]byte(str))
	str = hex.EncodeToString(h.Sum(nil))
	log.Info("--------------" + str + "--------------")
	return str
}
