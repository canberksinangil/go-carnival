package main

func main() {
	var mySlice1 = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySLice2 := mySlice1[5:]
	mySLice2[0] = 99
	println(mySlice1[5])
}
