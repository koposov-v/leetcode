package main

//type event struct {
//	pos  int
//	typ  int // -1 = начало отрезка, 0 = точка, 1 = конец отрезка
//	idx  int // индекс точки (только для тип == 0)
//}

//func main() {
//	var n, m int
//	fmt.Scan(&n, &m)
//
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Split(bufio.ScanWords)
//
//	events := make([]event, 0, 2*n+m)
//
//	// Чтение отрезков
//	for i := 0; i < n; i++ {
//		var a, b int
//		fmt.Scan(&a, &b)
//		if a > b {
//			a, b = b, a
//		}
//		events = append(events, event{a, -1, -1}) // начало
//		events = append(events, event{b, 1, -1})  // конец
//	}
//
//	points := make([]int, m)
//	// Чтение точек
//	for i := 0; i < m; i++ {
//		fmt.Scan(&points[i])
//		events = append(events, event{points[i], 0, i})
//	}
//
//	// Сортируем события
//	sort.Slice(events, func(i, j int) bool {
//		if events[i].pos == events[j].pos {
//			return events[i].typ < events[j].typ // порядок: начало < точка < конец
//		}
//		return events[i].pos < events[j].pos
//	})
//
//	result := make([]int, m)
//	active := 0
//
//	for _, e := range events {
//		switch e.typ {
//		case -1:
//			active++
//		case 1:
//			active--
//		case 0:
//			result[e.idx] = active
//		}
//	}
//
//	// Печать результата
//	for i := 0; i < m; i++ {
//		fmt.Printf("%d ", result[i])
//	}
//	fmt.Println()
//}
