package sort

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getElement(n int) []int {
	els := []int{}
	i := 0
	for j := n -1 ; j >= i; j--  {
		els = append(els, j)
	}
	return els
}

func TestBubbleSortAsc(t *testing.T) {
	s := []int{1, 3, 2, 6, 5, 4}
	BubbleSortAsc(s)
	fmt.Println(s)
	assert.EqualValues(t, []int{1, 2, 3, 4, 5, 6}, s)
}

func BenchmarkBubbleSortAsc(b *testing.B) {
	els := getElement(100000)
	for i := 0; i < b.N; i++ {
		BubbleSortAsc(els)
	}
}

func TestBasicSortAsc(t *testing.T) {
	els := getElement(10)
	fmt.Println(els)
	BasicSortAsc(els)
	fmt.Println(els)
	assert.EqualValues(t, els[0], 0)
}

func BenchmarkBasicSortAsc(b *testing.B) {
	els := getElement(100000)
	for i := 0; i <= b.N; i++ {
		BasicSortAsc(els)
	}
}
