func IsEmail(email string) bool {
	if email != "" {
		if isOk, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", email); isOk {
			return true
		}
	}
	return false
}

func main() {
	s1 := "我是中国人hello word!,2020 street 188#"
	s2 := "aaaa"
	fmt.Println(IsEmail(s1))
	fmt.Println(IsEmail(s2))
}
