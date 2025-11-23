package main

import "fmt"

type Student struct {
	Name   string
	Scores []int
}

func (student Student) Average() int {
	meanScore := 0
	totalScore, scoreCount := 0, 0
	for _, score := range student.Scores {
		totalScore += score
		scoreCount += 1
	}
	if scoreCount != 0 {
		meanScore = totalScore / scoreCount
	}
	return meanScore
}

func (student Student) Max() int {
	maxScore := 0
	for _, score := range student.Scores {
		if maxScore < score {
			maxScore = score
		}
	}
	return maxScore
}

func (student Student) Min() int {
	minScore := 100
	for _, score := range student.Scores {
		if score < minScore {
			minScore = score
		}
	}
	return minScore
}

func main() {
	students := []Student{
		{"Aki", []int{80, 70, 90}},
		{"Ben", []int{60, 75, 65, 70}},
		{"Cara", []int{100, 95}},
		{"Dan", []int{50, 40, 55}},
	}
	totalScore := 0
	for _, student := range students {
		meanScore := student.Average()
		totalScore += meanScore
		maxScore := student.Max()
		minScore := student.Min()
		fmt.Println(student.Name)
		fmt.Printf("平均点：%d点\n", meanScore)
		fmt.Printf("最高点：%d点\n", maxScore)
		fmt.Printf("最低点：%d点\n", minScore)
	}
	fmt.Println("全体")
	fmt.Printf("平均点：%d点\n", totalScore/len(students))
}
