package tools

func EmailValidator(email string) bool {
	return true
}
func TelValidator(tel string) bool {
	if len(tel) < 5 {
		return false
	}
	return true
}
func PwdValidator(pwd string) bool {
	return true
}
func BilingValidator(addr string) bool {
	if len(addr) < 5 {
		return false
	}
	return true
}
