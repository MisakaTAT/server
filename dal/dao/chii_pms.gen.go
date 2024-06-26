// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

const TableNamePrivateMessage = "chii_pms"

// PrivateMessage mapped from table <chii_pms>
type PrivateMessage struct {
	ID                uint32 `gorm:"column:msg_id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:""`
	SenderID          uint32 `gorm:"column:msg_sid;type:mediumint(8) unsigned;not null" json:""`
	ReceiverID        uint32 `gorm:"column:msg_rid;type:mediumint(8) unsigned;not null" json:""`
	Folder            string `gorm:"column:msg_folder;type:enum('inbox','outbox');not null;default:inbox" json:""`
	New               bool   `gorm:"column:msg_new;type:tinyint(1);not null" json:""`
	Title             string `gorm:"column:msg_title;type:varchar(75);not null" json:""`
	CreatedTime       uint32 `gorm:"column:msg_dateline;type:int(10) unsigned;not null" json:""`
	Content           string `gorm:"column:msg_message;type:text;not null" json:""`
	MainMessageID     uint32 `gorm:"column:msg_related_main;type:int(10) unsigned;not null" json:""`
	RelatedMessageID  uint32 `gorm:"column:msg_related;type:int(10) unsigned;not null" json:""`
	DeletedBySender   bool   `gorm:"column:msg_sdeleted;type:tinyint(1) unsigned;not null" json:""`
	DeletedByReceiver bool   `gorm:"column:msg_rdeleted;type:tinyint(1) unsigned;not null" json:""`
}

// TableName PrivateMessage's table name
func (*PrivateMessage) TableName() string {
	return TableNamePrivateMessage
}
