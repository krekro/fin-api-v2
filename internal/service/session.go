package service

func ValidateSesssion(id string) bool {
	if id != "" {
		return true
	} else {
		return false
	}
}
