package reckon

import (
	"common/utils/log"
	"common/service"
	"common/model"
	"common/utils/convert"
	"errors"
)

type ReckonHandler struct {
	handicap *model.Handicap
}

func NewReckonHandler() *ReckonHandler {
	return &ReckonHandler{}
}

func (reckonHandler *ReckonHandler) Run(msgs *model.Msgs) {
	reckonHandler.requestReckon(msgs)
	log.Info("[Reckon " + convert.ToStr(msgs.Han_id) + "] starting")
}

func (reckonHandler *ReckonHandler) requestReckon(msgs *model.Msgs) {
	service.GetHandicapById(msgs.Han_id, &reckonHandler.handicap)
	if reckonHandler.handicap.Han_id > 0 {
		if !service.HandicapHashIsValid(reckonHandler.handicap) {
			log.Error2Exit(errors.New("handicap["+convert.ToStr(msgs.Han_id)+"] HandicapHashIsValid faile"), "error")
		}
	} else {
		log.Error2Exit(errors.New("handicap["+convert.ToStr(msgs.Han_id)+"] not found in db"), "error")
	}
}
