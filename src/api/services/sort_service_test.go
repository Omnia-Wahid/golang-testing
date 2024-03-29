package services

import (
	"github.com/omniaw/golang-testing/src/api/utils/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConstants(t *testing.T) {
	if privateConst != "private" {
		t.Error("privateConst should be 'private'")
	}
}

func TestSort(t *testing.T) {
	//elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}
	elements := sort.GetElements(10)
	Sort(elements)
	assert.NotNil(t, elements)
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestSortMoreThan10000(t *testing.T) {
	//elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}
	elements := sort.GetElements(20001)
	Sort(elements)
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 20000 {
		t.Error("last element should be 20000")
	}
}

func BenchmarkBubbleSort10K(b *testing.B) {
	//elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}
	elements := sort.GetElements(20000)
	//fmt.Println(elements)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkSort100K(b *testing.B) {
	elements := sort.GetElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
