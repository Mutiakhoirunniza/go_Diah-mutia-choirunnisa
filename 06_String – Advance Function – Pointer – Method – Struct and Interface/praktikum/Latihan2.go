package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score int
}

type Students []Student

func (s *Students) AddStudent(name string, score int) {
	student := Student{Name: name, Score: score}
	*s = append(*s, student)
}

func (s Students) AverageScore() float64 {
	totalScore := 0
	for _, student := range s {
		totalScore += student.Score
	}
	return float64(totalScore) / float64(len(s))
}

func (s Students) MinScore() Student {
	minScore := math.MaxInt32
	var minStudent Student
	minIndex := -1

	for i, student := range s {
		if student.Score < minScore {
			minScore = student.Score
			minStudent = student
			minIndex = i
		}
	}

	if minIndex != -1 {
		s = append(s[:minIndex], s[minIndex+1:]...)
	}

	return minStudent
}

func (s Students) MaxScore() Student {
	maxScore := math.MinInt32
	var maxStudent Student
	maxIndex := -1

	for i, student := range s {
		if student.Score > maxScore {
			maxScore = student.Score
			maxStudent = student
			maxIndex = i
		}
	}

	if maxIndex != -1 {
		s = append(s[:maxIndex], s[maxIndex+1:]...)
	}

	return maxStudent
}

func main() {
	var students Students

	for i := 0; i < 5; i++ {
		var name string
		var score int
		fmt.Printf("Masukkan nama siswa ke-%d: ", i+1)
		fmt.Scan(&name)
		fmt.Printf("Masukkan skor siswa ke-%d: ", i+1)
		fmt.Scan(&score)
		students.AddStudent(name, score)
	}

	average := students.AverageScore()
	minStudent := students.MinScore()
	maxStudent := students.MaxScore()

	fmt.Printf("Skor rata-rata: %.2f\n", average)
	fmt.Printf("Siswa dengan skor minimum: %s (%d)\n", minStudent.Name, minStudent.Score)
	fmt.Printf("Siswa dengan skor maksimum: %s (%d)\n", maxStudent.Name, maxStudent.Score)
}
