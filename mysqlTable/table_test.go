package mysqlTable

import (
	"encoding/json"
	"time"
)

//
type Test struct {
	Id         int32     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"` //
	Name       string    `gorm:"column:name" json:"name"`                        //
	Card       string    `gorm:"column:card" json:"card"`                        //
	Num        int32     `gorm:"column:num" json:"num"`                          //
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`          //
}

func (This *Test) TableName() string {
	return "test"
}

func (This *Test) ToJson() string {
	if jsonByte, err := json.Marshal(This); err != nil {
		return ""
	} else {
		return string(jsonByte)
	}
}
