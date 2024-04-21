// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNameEpisode = "chii_episodes"

// Episode mapped from table <chii_episodes>
type Episode struct {
	ID        uint32  `gorm:"column:ep_id;type:mediumint(8) unsigned;primaryKey;autoIncrement:true" json:""`
	SubjectID uint32  `gorm:"column:ep_subject_id;type:mediumint(8) unsigned;not null" json:""`
	Sort      float32 `gorm:"column:ep_sort;type:float unsigned;not null" json:""`
	Type      uint8   `gorm:"column:ep_type;type:tinyint(1) unsigned;not null" json:""`
	Disc      uint8   `gorm:"column:ep_disc;type:tinyint(3) unsigned;not null;comment:碟片数" json:""` // 碟片数
	Name      string  `gorm:"column:ep_name;type:varchar(80);not null" json:""`
	NameCn    string  `gorm:"column:ep_name_cn;type:varchar(80);not null" json:""`
	Rate      int8    `gorm:"column:ep_rate;type:tinyint(3);not null" json:""`
	Duration  string  `gorm:"column:ep_duration;type:varchar(80);not null" json:""`
	Airdate   string  `gorm:"column:ep_airdate;type:varchar(80);not null" json:""`
	Online    string  `gorm:"column:ep_online;type:mediumtext;not null" json:""`
	Comment   uint32  `gorm:"column:ep_comment;type:mediumint(8) unsigned;not null" json:""`
	Resources uint32  `gorm:"column:ep_resources;type:mediumint(8) unsigned;not null" json:""`
	Desc      string  `gorm:"column:ep_desc;type:mediumtext;not null" json:""`
	Dateline  uint32  `gorm:"column:ep_dateline;type:int(10) unsigned;not null" json:""`
	Lastpost  uint32  `gorm:"column:ep_lastpost;type:int(10) unsigned;not null" json:""`
	Lock      uint8   `gorm:"column:ep_lock;type:tinyint(3) unsigned;not null" json:""`
	Ban       uint8   `gorm:"column:ep_ban;type:tinyint(3) unsigned;not null" json:""`
	Subject   Subject `gorm:"foreignKey:ep_subject_id;references:subject_id" json:"subject"`
}

// TableName Episode's table name
func (*Episode) TableName() string {
	return TableNameEpisode
}
