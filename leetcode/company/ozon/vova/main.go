package vova

import "fmt"

type Entry struct {
	userId int
	steps  int
}

type Champions struct {
	userIds []int
	steps   int
}

// findChampions принимает статистику по шагам и возвращает победителей
func findChampions(statistics [][]Entry) Champions {
	res := Champions{}
	userSteps := make(map[int]int)
	userDays := make(map[int]int)
	totalDays := len(statistics)

	for _, stat := range statistics {
		for _, day := range stat {
			userSteps[day.userId] += day.steps
			userDays[day.userId]++
		}
	}

	maxSteps := 0
	for userId, days := range userDays {
		if days != totalDays {
			continue
		}
		res.userIds = append(res.userIds, userId)

		if userSteps[userId] > maxSteps {
			maxSteps = userSteps[userId]
		}
	}

	filtered := make([]int, 0, len(res.userIds))

	for _, userId := range res.userIds {
		if maxSteps == userSteps[userId] {
			filtered = append(filtered, userId)
		}
	}

	res.userIds = filtered
	res.steps = maxSteps

	return res
}

func main() {
	statistics := [][]Entry{
		{{userId: 1, steps: 2000}, {userId: 2, steps: 1500}},
		{{userId: 2, steps: 4000}, {userId: 1, steps: 3500}},
	}

	fmt.Println(findChampions(statistics))
}
