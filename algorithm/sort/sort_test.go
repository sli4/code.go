package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	data := []int{3,7,5,2,1,6,4}
	fmt.Println("冒泡排序：",BubbleSort(data))
}

func TestInsertSort(t *testing.T) {
	data := []int{3,7,5,2,1,6,4}
	fmt.Println("插入排序1：",InsertSort(data))
}

func TestInsertSort2(t *testing.T) {
	data := []int{3,7,5,2,1,6,4}
	fmt.Println("插入排序2：",InsertSort2(data))
}

func TestSelectSort(t *testing.T) {
	data := []int{3,7,5,2,1,6,4}
	fmt.Println("选择排序：",SelectSort(data))

	// Just Test Array Point
	p := &data
	fmt.Println("Point:", p)
	fmt.Println("data[1]:", (*p)[1])
}

func TestMergeSort(t *testing.T) {
	a := []int{1,3,5}
	b := []int{2,4,6}
	fmt.Println("归并：", merge(a,b))
	data := []int{3,7,5,2,1,6,4}
	fmt.Println("归并排序：", MergeSort(data))
}

func TestQuickSort(t *testing.T) {
	data := []int{3,7,5,2,1,6,4}
	QuickSort(data)
}