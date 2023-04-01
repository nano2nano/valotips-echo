package model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string `json:"title"`
	StandImg    string `json:"stand_img"`
	AimImg      string `json:"aim_img"`
	StrongLikes int    `json:"strong_likes" gorm:"default:0"`
	FunnyLikes  int    `json:"funny_likes" gorm:"default:0"`
	MapUUID     string `json:"map_uuid"`
	AgentUUID   string `json:"agent_uuid"`
	AbilitySlot string `json:"ability_slot"`
}
