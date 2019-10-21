package model

import "time"

type Handicap struct {
	Han_id      int64     `bson:"han_id"`
	Game_id     int64     `bson:"game_id"`
	Name        string    `bson:"name"`
	Han_type    int       `bson:"hantype"`
	Return_rate float32   `bson:"returnrate"`
	Result      int       `bson:"result"`
	Resulttime  time.Time `bson:"resulttime"`
	Status      int       `bson:"status"`
}
