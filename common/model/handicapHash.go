package model

const HANDICAP_HASH_KEY = "de2%YGRwib)vde&)MD~!"

type HandicapHash struct {
	Han_id int64  `bson:"han_id"`
	Hash1  string `bson:"hash1"`
}
