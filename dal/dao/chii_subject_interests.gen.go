// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"github.com/bangumi/server/dal/utiltype"
)

const TableNameSubjectCollection = "chii_subject_interests"

// SubjectCollection mapped from table <chii_subject_interests>
type SubjectCollection struct {
	ID           uint32                     `gorm:"column:interest_id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:""`
	UserID       uint32                     `gorm:"column:interest_uid;type:mediumint(8) unsigned;not null" json:""`
	SubjectID    uint32                     `gorm:"column:interest_subject_id;type:mediumint(8) unsigned;not null" json:""`
	SubjectType  uint8                      `gorm:"column:interest_subject_type;type:smallint(6) unsigned;not null" json:""`
	Rate         uint8                      `gorm:"column:interest_rate;type:tinyint(3) unsigned;not null" json:""`
	Type         uint8                      `gorm:"column:interest_type;type:tinyint(1) unsigned;not null" json:""`
	HasComment   bool                       `gorm:"column:interest_has_comment;type:tinyint(1) unsigned;not null" json:""`
	Comment      utiltype.HTMLEscapedString `gorm:"column:interest_comment;type:mediumtext;not null" json:""`
	Tag          string                     `gorm:"column:interest_tag;type:mediumtext;not null" json:""`
	EpStatus     uint32                     `gorm:"column:interest_ep_status;type:mediumint(8) unsigned;not null" json:""`
	VolStatus    uint32                     `gorm:"column:interest_vol_status;type:mediumint(8) unsigned;not null;comment:卷数" json:""` // 卷数
	WishTime     uint32                     `gorm:"column:interest_wish_dateline;type:int(10) unsigned;not null" json:""`
	DoingTime    uint32                     `gorm:"column:interest_doing_dateline;type:int(10) unsigned;not null" json:""`
	DoneTime     uint32                     `gorm:"column:interest_collect_dateline;type:int(10) unsigned;not null" json:""`
	OnHoldTime   uint32                     `gorm:"column:interest_on_hold_dateline;type:int(10) unsigned;not null" json:""`
	DroppedTime  uint32                     `gorm:"column:interest_dropped_dateline;type:int(10) unsigned;not null" json:""`
	CreateIP     string                     `gorm:"column:interest_create_ip;type:char(15);not null" json:""`
	LastUpdateIP string                     `gorm:"column:interest_lasttouch_ip;type:char(15);not null" json:""`
	UpdatedTime  uint32                     `gorm:"column:interest_lasttouch;type:int(10) unsigned;not null" json:""`
	Private      uint8                      `gorm:"column:interest_private;type:tinyint(1) unsigned;not null" json:""`
}

// TableName SubjectCollection's table name
func (*SubjectCollection) TableName() string {
	return TableNameSubjectCollection
}
