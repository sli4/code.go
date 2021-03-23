package sort

import "fmt"

// 冒泡排序
func BubbleSort(data []int) []int {
	n := len(data)
	for i:=0;i<n;i++ {
		for j:=0;j<n-i-1;j++ {
			if data[j]>data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}
// 插入排序
func InsertSort(data []int) []int {
	n := len(data)
	for i:=1;i<n;i++ {
		for j:=0;j<i;j++ {
			if data[i]<data[j] {
				tmp := data[i]
				for k:=i;k>j;k-- {
					data[k] = data[k-1]
				}
				data[j] = tmp
			}
		}
	}
	return data
}

func InsertSort2(data []int) []int {
	n := len(data)
	for i:=1;i<n;i++ {
		value := data[i]
		j:=i-1
		for ;j>=0;j-- {
			if data[j]>value{
				data[j+1] = data[j]
			} else {
				break
			}
		}
		data[j+1] = value   // 插入数据, 每次比较都要插入
	}
	return data
}

func SelectSort(data []int) []int {
	n := len(data)
	for i:=0;i<n;i++ {
		min := data[i]
		minIndex := i
		for j:=i+1; j<n; j++ {
			if data[j]< min {
				//data[j],data[i]=data[i],data[j]
				min = data[j]
				minIndex = j
			}
		}
		data[i],data[minIndex] = data[minIndex], data[i]
	}
	return data
}

// 归并排序
func MergeSort(data []int) []int {
	return mergeSort(data, 0,len(data)-1)
}
func mergeSort(data []int, p,r int) []int {
	if p>=r {
		return data[p:r+1]
	}

	q := (p+r)/2

	a := mergeSort(data, p,q)
	b := mergeSort(data, q+1, r)
	return merge(a,b)

}
func merge(a,b []int) []int {
	m := []int{}
	i,j := 0,0
	for ;i<len(a)&&j<len(b); {
		if a[i]<b[j] {
			m = append(m, a[i])
			i++
		} else {
			m = append(m, b[j])
			j++
		}
	}
	if i<len(a) {
		m = append(m, a[i:len(a)]...)
	}
	if j<len(b) {
		m = append(m, b[j:len(b)]...)
	}
	return m
}


// 快速排序
func QuickSort(data []int)  {
	quickSort(&data, 0, len(data)-1)
	fmt.Println("快速排序：", data)
}
func quickSort(data *[]int, p,r int)  {
	fmt.Println("QuickSort:",p, "...",r)
	if p>=r {
		return
	}
	q := partition(data, p, r)
	quickSort(data, p, q-1)
	quickSort(data, q+1, r)
}

func partition(dp *[]int, p,r int) int {
	data := *dp
	pivot := data[r]
	i := p  // 标记第一个大于pivot的元素下标
	for j:=p; j<r; j++{
		if data[j] < pivot {
			data[i],data[j] = data[j], data[i]
			i++
		}
	}
	data[i],data[r] = data[r], data[i]
	fmt.Println("Partition:",i)
	return i
}