package response

var ZhTwTranslate = map[int]string{
	Success:                      "成功",
	Error:                        "失敗",
	ParameterError:               "參數驗證失敗",
	SigMsgExpired:                "簽名消息過期",
	LoginError:                   "登陸失敗",
	TokenFailure:                 "請重新登陸",
	RateIsTooHigh:                "速率過高",
	TokenError:                   "登陸失敗",
	SignError:                    "簽名錯誤",
	UserNotExist:                 "用戶不存在",
	UploadError:                  "上傳失敗",
	HorseInfoNotExist:            "馬匹信息不存在",
	HorseParticipatedRacesError:  "選擇馬匹已參加其他比賽",
	MatchGameClosedError:         "比賽已關閉",
	MatchGameEndError:            "正在進行比賽，暫停報名",
	MatchGameNumberSelectedError: "此閘門已被選擇",
	NetworkBusy:                  "網絡繁忙，請稍後再試",
	RacesAchieveLimit:            "參賽馬匹已達數量限制",
	MatchGameInfoNotExist:        "比賽場次資訊不存在",
	EntryFailure:                 "加入比賽失敗",
}
