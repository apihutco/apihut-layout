package conf

func IsDevMode(conf Bootstrap_Mode) bool {
	if conf == Bootstrap_Dev {
		return true
	}
	return false
}
