package main

import (
	"fmt"
)

// Struct Student dengan properti name dan score
type Student struct {
	Name  string
	Score int
}

// Method untuk menghitung skor rata-rata dari slice Student
func (students []Student) CalculateAverageScore() float64 {
	totalScore := 0
	for _, student := range students {
		totalScore += student.Score
	}
	return float64(totalScore) / float64(len(students))
}

// Method untuk mencari siswa dengan skor minimum
func (students []Student) FindMinScoreStudent() Student {
	minStudent := students[0]
	for _, student := range students {
		if student.Score < minStudent.Score {
			minStudent = student
		}
	}
	return minStudent
}

// Method untuk mencari siswa dengan skor maksimum
func (students []Student) FindMaxScoreStudent() Student {
	maxStudent := students[0]
	for _, student := range students {
		if student.Score > maxStudent.Score {
			maxStudent = student
		}
	}
	return maxStudent
}

func main() {
	var students []Student

	for i := 1; i <= 5; i++ {
		var name string
		var score int

		fmt.Printf("Masukkan nama siswa ke-%d: ", i)
		fmt.Scan(&name)
		fmt.Printf("Masukkan skor siswa ke-%d: ", i)
		fmt.Scan(&score)

		student := Student{Name: name, Score: score}
		students = append(students, student)
	}

	// Hitung skor rata-rata
	averageScore := students.CalculateAverageScore()
	fmt.Printf("Skor rata-rata: %0.2f\n", averageScore)

	// Cari siswa dengan skor minimum dan maksimum
	minStudent := students.FindMinScoreStudent()
	maxStudent := students.FindMaxScoreStudent()

	fmt.Printf("Siswa dengan skor minimum: %s (Skor: %d)\n", minStudent.Name, minStudent.Score)
	fmt.Printf("Siswa dengan skor maksimum: %s (Skor: %d)\n", maxStudent.Name, maxStudent.Score)
}
