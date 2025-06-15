package main

func romanToInt(s string) int {
	res := 0
	hashMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && hashMap[s[i]] < hashMap[s[i+1]] {
			res -= hashMap[s[i]]
		} else {
			res += hashMap[s[i]]
		}
	}
	return res
}
