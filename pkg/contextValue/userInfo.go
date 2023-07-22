package contextValue

type LoginUserInfo struct {
	Id             uint64 `json:"id"`
	Nickname       string `json:"nickname"`
	Uuid           int    `json:"uuid"`
	InvitationCode string `json:"invitation_code"`
	ApiAuth        string `json:"apiAuth"`
	Serial         uint   `json:"serial"`
}
