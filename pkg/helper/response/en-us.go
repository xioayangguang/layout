package response

var EnTranslate = map[int]string{
	Success:                      "Success",
	Error:                        "Failed",
	ParameterError:               "Parameter validation failed",
	SigMsgExpired:                "Signature message expired",
	LoginError:                   "Login failed",
	TokenFailure:                 "Please log in again",
	RateIsTooHigh:                "Rate is too high",
	TokenError:                   "Login failed",
	SignError:                    "Signature error",
	UserNotExist:                 "User does not exist",
	UploadError:                  "Upload failed",
	HorseInfoNotExist:            "Horse information does not exist",
	HorseParticipatedRacesError:  "The selected horses is already in aother race.",
	MatchGameClosedError:         "The race is closed.",
	MatchGameEndError:            "The race is underway. The registration is suspended.",
	MatchGameNumberSelectedError: "This gate has been selected",
	NetworkBusy:                  "The network is busy, please try again later",
	RacesAchieveLimit:            "The horse number in the race has reached the limit.",
	MatchGameInfoNotExist:        "The race not exist",
	EntryFailure:                 "Entry failure",
}
