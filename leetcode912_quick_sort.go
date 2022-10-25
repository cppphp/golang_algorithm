package main

import (
	"fmt"
)

func main(){
	arr := []int{2,1,6,4,3,9,7,8,5,0}
	part(arr, 0, len(arr)-1)
	fmt.Printf("%v\n", arr)
}

func part(arr []int,l, r int) {
	if r - l < 3 {
		if arr[l] > arr[r]{
			arr[l],arr[r] = arr[r],arr[l]
		}
		return
	}
	mid := (r-l)/2+l
	arr[r],arr[mid] = arr[mid],arr[r]
	var i,j int
	for i,j=l,l; i<=r; i++ {
		if arr[j] <= arr[r] {
			j++
		}
		if arr[i] <= arr[r] && j < i {
			arr[i],arr[j] = arr[j],arr[i]
			j++
		}
	}
	println(l, r, mid, j)
	part(arr, l, j-1)
	part(arr, j, r)
}
