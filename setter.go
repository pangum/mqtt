package mqtt

func setString(setter func(value string), value string, null string) {
	if "" != value {
		setter(value)
	} else {
		setter(null)
	}
}
