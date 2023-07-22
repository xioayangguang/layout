package response

var EnTranslate = map[int]string{
	Success:        "Success",
	Error:          "Failed",
	ParameterError: "Parameter validation failed",
	SigMsgExpired:  "Signature message expired",
	LoginError:     "Login failed",
	TokenFailure:   "Please log in again",
	RateIsTooHigh:  "Rate is too high",
	TokenError:     "Login failed",
	SignError:      "Signature error",
	UserNotExist:   "User does not exist",
	UploadError:    "Upload failed",
}
