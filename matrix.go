package main

import "fmt"

type Connections [9]uint8
type Matrix [9]Connections

func NewMatrix() *Matrix {
	return &Matrix{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
}

func (m Matrix) String() string {
	res := "   0 1 2 3 4 5 6 7 8\n"
	for i, connection := range m {
		res += fmt.Sprintf("%d %v\n", i, connection)
	}
	return res
}

func (m *Matrix) AddingIsLessThan(m2 *Matrix, max uint8) bool {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j]+m2[i][j] > max {
				return false
			}
		}
	}
	return true
}

func (m *Matrix) Equals(m2 *Matrix) bool {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] != m2[i][j] {
				return false
			}
		}
	}
	return true
}

func (m *Matrix) Add(m2 *Matrix) *Matrix {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			m[i][j] += m2[i][j]
		}
	}
	return m
}

func (m *Matrix) Sub(m2 *Matrix) *Matrix {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			m[i][j] -= m2[i][j]
		}
	}
	return m
}
