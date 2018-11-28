package main

import (
	"fmt"
)

type Tweet struct {
	User     string
	Message  string
	Retweets int
}

func shellSort(arr []int) []int {
	length := len(arr)
	gap := length / 2
	for gap > 0 {
		for index := gap; index < length; index++ {
			j := index
			for j-gap >= 0 && arr[j] < arr[j-gap] {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
				j -= gap
			}
		}
		gap /= 2
	}
	return arr
}
func main() {
	arr := []int{39, 80, 76, 41, 13, 29, 50, 78, 30, 11, 100, 7, 41, 86}
	var ar []int
	length := len(arr) - 1
	for index := 0; index < length; index++ {
		ArrayHeap(arr[:length-index])
		arr[0], arr[length-index] = arr[length-index], arr[0]
		ar = append(ar, arr[length-index])
	}
	fmt.Println(ar)
}
func ArrayHeap(arr []int) {

	for i := len(arr)/2 - 1; i >= 0; i-- {
		if arr[i] > arr[2*i+1] {
			arr[i], arr[2*i+1] = arr[2*i+1], arr[i]
		} else {
		}
		if 2*i+2 < len(arr) {
			if arr[i] > arr[2*i+2] {
				arr[i], arr[2*i+2] = arr[2*i+2], arr[i]
			}
		}
	}
	for i := 0; i <= len(arr)/2-1; i++ {
		if arr[i] > arr[2*i+1] {
			arr[i], arr[2*i+1] = arr[2*i+1], arr[i]
		} else {
		}
		if 2*i+2 < len(arr) {
			if arr[i] > arr[2*i+2] {
				arr[i], arr[2*i+2] = arr[2*i+2], arr[i]
			}
		}
	} ///

}

type TreeNode struct {
	Data      int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

func CreatHeap(arr []int) {
	tree := &TreeNode{Data: arr[0]}

	queue := []*TreeNode{tree}
	var curNode *TreeNode
	length := len(arr)
	var index int

	for len(queue) != 0 {
		curNode = queue[0]
		queue = queue[1:]
		if curNode.LeftNode == nil {
			if 2*index+1 < length {
				curNode.LeftNode = &TreeNode{Data: arr[2*index+1]}
				fmt.Println(curNode.LeftNode)
				queue = append(queue, curNode.LeftNode)

			} else {
				return
			}
		}
		if curNode.RightNode == nil {
			if (2*index + 2) >= length {
				return

			} else {
				curNode.LeftNode = &TreeNode{Data: arr[2*index+2]}
				queue = append(queue, curNode.RightNode)
			}
		}
		index++
	}
}

func test() {
	s1 := []int{5, 9, 12, 5, 24, 7, 8, 4, 17, 9, 2, 4, 13, 1, 24, 6, 2, 31, 12, 5, 10, 24}
	m1 := make(map[int]int)
	var s2 []int
	var max int
	var s3 []int

	for _, v := range s1 {
		if m1[v] != 0 {
			m1[v]++
		} else {
			m1[v] = 1
		}
	}
	fmt.Println(m1)

	for k, vv := range m1 {

		if vv > 1 {
			fmt.Println(k, vv)

		}
	}

	for _, v := range m1 {
		s2 = append(s2, v)
	}
	max = s2[0]
	for i := 0; i < len(s2); i++ {
		if max < s2[i] {
			max = s2[i]
		}
	}
	fmt.Println(max)

	for k, v := range m1 {
		if v == max {
			s3 = append(s3, k)
		}
	}
	fmt.Println(s3)
}
