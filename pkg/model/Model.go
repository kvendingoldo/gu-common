package model

import (
	"time"
)

type BaseModel struct {
	ID        int64     `json:"id" example:"23" gorm:"type:int;primaryKey"`
	CreatedAt time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
}
