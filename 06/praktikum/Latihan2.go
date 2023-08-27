package main

import (
	"fmt"
	"math"
)

type Student struct {
	name  string
	score []float64
}

func (s Student) AverageScore() float64 {
	total := 0.0
	for _, sc := range s.score {
		total += sc
	}
	return total / float64(len(s.score))
}

func (s Student) MinScore() float64 {
	minScore := math.MaxFloat64
	for _, sc := range s.score {
		if sc < minScore {
			minScore = sc
		}
	}
	return minScore
}

func (s Student) MaxScore() float64 {
	maxScore := -math.MaxFloat64
	for _, sc := range s.score {
		if sc > maxScore {
			maxScore = sc
		}
	}
	return maxScore
}

func main() {
	students := make([]Student, 0)

	for i := 1; i <= 5; i++ {
		var name string
		var scores []float64

		fmt.Printf("Enter name of student %d: ", i)
		fmt.Scan(&name)

		for j := 1; j <= 3; j++ {
			var score float64
			fmt.Printf("Enter score %d for student %d: ", j, i)
			fmt.Scan(&score)
			scores = append(scores, score)
		}

		student := Student{
			name:  name,
			score: scores,
		}

		students = append(students, student)
	}

	// Input 1 Student's Name and 1 Student's Score
	var newName string
	var newScore float64

	fmt.Print("Enter name of a new student: ")
	fmt.Scan(&newName)

	fmt.Print("Enter score of the new student: ")
	fmt.Scan(&newScore)

	newStudent := Student{
		name:  newName,
		score: []float64{newScore},
	}

	students = append(students, newStudent)

	totalAverage := 0.0
	minScore := math.MaxFloat64
	maxScore := -math.MaxFloat64

		for _, student := range students {
		average := student.AverageScore()
		if average < minScore {
			minScore = average
		}
		if average > maxScore {
			maxScore = average
		}
		totalAverage += average
	}

	fmt.Printf("\nAverage scores:\n")
	for student := range students {
		average := student.AverageScore()
		fmt.Printf("%s: %.2f\n", student.name, average)
	}

	fmt.Printf("\nAverage of all students: %.2f\n", totalAverage/float64(len(students)))
	fmt.Printf("Student with minimum average score: %.2f\n", minScore)
	fmt.Printf("Student with maximum average score: %.2f\n", maxScore)
}
