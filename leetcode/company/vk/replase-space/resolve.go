package replase_space

//func replaceSpaces(s string) string {
//	var result []rune
//	for _, ch := range s {
//		if ch == ' ' {
//			result = append(result, '@', '4', '0')
//		} else {
//			result = append(result, ch)
//		}
//	}
//	return string(result)
//}

//func main() {
//	scanner := bufio.NewScanner(os.Stdin)
//
//	// Читаем количество тестов
//	scanner.Scan()
//	var T int
//	fmt.Sscanf(scanner.Text(), "%d", &T)
//
//	for i := 0; i < T; i++ {
//		scanner.Scan()
//		line := scanner.Text()
//
//		// Простая замена пробелов на "@40"
//		result := strings.ReplaceAll(line, " ", "@40")
//
//		fmt.Println(result)
//	}
//}
