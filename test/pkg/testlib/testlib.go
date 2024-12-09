package testlib

func Истина(in string) bool {
	switch in {
	case "t", "true", "+", "да":
		return true
	}
	return false
}
