func createString(x int) string {
	tmpStr := "компьютеров"
	if x%10 == 1 {
		tmpStr = "компьютер"
	}
	if x%10 == 2 || x%10 == 3 || x%10 == 4 {
		tmpStr = "компьютера"
	}
	return fmt.Sprint(x, " ", tmpStr)
}


