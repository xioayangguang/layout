package response

const (
	Success                      = 0
	Error                        = 1
	ParameterError               = 1000
	LoginError                   = 1001
	TokenFailure                 = 1002
	RateIsTooHigh                = 1003
	TokenError                   = 1004
	SignError                    = 1005
	SigMsgExpired                = 1006
	UserNotExist                 = 1007
	UploadError                  = 1008
	HorseInfoNotExist            = 1009
	HorseParticipatedRacesError  = 1010
	MatchGameClosedError         = 1011
	MatchGameEndError            = 1012
	MatchGameNumberSelectedError = 1013
	NetworkBusy                  = 1014
	RacesAchieveLimit            = 1015
	MatchGameInfoNotExist        = 1016
	EntryFailure                 = 1017
)
