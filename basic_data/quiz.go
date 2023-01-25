package basic_data

import "fmt"

func main() {
	// num 1 - hitung rata-rata
	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	// total / n
	// var total int
	// for _, score := range scores {
	// 	total += score
	// }
	// result := float64(total) / float64(len(scores))
	// fmt.Println(total)
	// fmt.Println(result)

	// num 2 - good scores
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int

	for _, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}
	fmt.Println(goodScores)
}