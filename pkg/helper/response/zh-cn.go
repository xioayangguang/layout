package response

var ZhCnTranslate = map[int]string{
	Success:                      "成功",
	Error:                        "失败",
	ParameterError:               "参数验证失败",
	SigMsgExpired:                "签名消息过期",
	LoginError:                   "登陆失败",
	TokenFailure:                 "请重新登陆",
	RateIsTooHigh:                "速率过高",
	TokenError:                   "登陆失败",
	SignError:                    "签名错误",
	UserNotExist:                 "用户不存在",
	UploadError:                  "上传失败",
	HorseInfoNotExist:            "马匹信息不存在",
	HorseParticipatedRacesError:  "选择马匹已参加其他比赛",
	MatchGameClosedError:         "比赛已关闭",
	MatchGameEndError:            "正在进行比赛，暂停报名",
	MatchGameNumberSelectedError: "此闸门已被选择",
	NetworkBusy:                  "网络繁忙, 请稍后再试",
	RacesAchieveLimit:            "参赛马匹已达数量限制",
	MatchGameInfoNotExist:        "比赛场次信息不存在",
	EntryFailure:                 "加入比赛失败",
}
