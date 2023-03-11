package model

import (
	"time"

	"gorm.io/gorm"
)

type ControlBy struct {
	CreateBy string `json:"createBy" gorm:"comment:创建者"`
	UpdateBy string `json:"updateBy" gorm:"comment:更新者"`
}

func (e *ControlBy) SetCreateBy(createBy string) {
	e.CreateBy = createBy
}

func (e *ControlBy) SetUpdateBy(updateBy string) {
	e.UpdateBy = updateBy
}

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`   // 在创建时，如果该字段值为零值，则使用当前时间填充。自动填充的字段都是内定命名的，当然也给定可以使用特定标签来指定特定命名的字段自动填充，比如CAt int64 `gorm:"autoUpdateTime:milli"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"` // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`     // 在删除时值为零，则自动填充
}
