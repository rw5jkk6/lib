package lib

func Swap(x int, y int)(int, int){
	tmp := x
	x = y
	y = tmp
	return x, y
}