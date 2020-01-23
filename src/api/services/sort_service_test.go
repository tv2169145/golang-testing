package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getElement(n int) []int {
	els := []int{}
	i := 0
	for j := n - 1; j >= i; j-- {
		els = append(els, j)
	}
	return els
}

func TestSortUseBubbleSort(t *testing.T) {
	s := getElement(100)
	Sort(s)
	fmt.Println(s)
	assert.EqualValues(t, 0, s[0])
}

func TestSortUseBasicSort(t *testing.T) {
	s := getElement(10001)
	Sort(s)
	assert.EqualValues(t, 10000, s[10000])
}

func BenchmarkSortUseBasicSort(b *testing.B) {
	s := getElement(100001)
	for i := 0; i < b.N; i++ {
		Sort(s)
	}
}

func BenchmarkSortUseBubbleSort(b *testing.B) {
	s := getElement(9999)
	for i := 0; i < b.N; i++ {
		Sort(s)
	}
}
