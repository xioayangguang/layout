package model

type User struct {
	Base
	Uuid           uint   `gorm:"column:uuid;type:int(10) unsigned;default:0;comment:9位非0开头纯数字;NOT NULL" json:"uuid"`
	Serial         uint   `gorm:"column:serial;type:int(11) unsigned;default:0;comment:9位非0开头纯数字;NOT NULL" json:"serial"`
	Nickname       string `gorm:"column:nickname;type:varchar(80);comment:用户昵称;NOT NULL" json:"nickname"`
	Mail           string `gorm:"column:mail;type:varchar(255);comment:用户邮箱;NOT NULL" json:"mail"`
	Describe       string `gorm:"column:describe;type:varchar(255);comment:用户描述;NOT NULL" json:"describe"`
	Code           string `gorm:"column:code;type:varchar(10);comment:我是被谁邀请的 (邀请人的邀请码);NOT NULL" json:"code"` //
	InvitationCode string `gorm:"column:invitation_code;type:varchar(15);comment:邀请码;NOT NULL" json:"invitation_code"`
	Avatar         string `gorm:"column:avatar;type:varchar(200);comment:头像;NOT NULL" json:"avatar"`
	Status         int    `gorm:"column:status;type:tinyint(4);default:1;comment:状态 1正常 -1拉黑;NOT NULL" json:"status"`
}

func (u *User) TableName() string {
	return "t_user"
}
