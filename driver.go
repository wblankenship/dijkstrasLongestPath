package main

func main() {
	m:=Dam{
		[][]int{
		[]int{00,00,00,00,00,00,00,00,00,00,00,00,00,00},//0
		[]int{00,00,00,00,00,00,00,00,00,00,00,00,00,28},//1
		[]int{00,00,00,00,00,00,00,00,00,00,00,00,00,00},//2
		[]int{20,00,00,00,00,00,00,00,00,00,00,00,00,00},//3
		[]int{10,15,00,00,00,00,00,00,00,00,00,00,00,00},//4
		[]int{00,30,00,00,00,00,00,00,00,00,00,00,00,00},//5
		[]int{00,00,12,00,00,00,00,00,00,00,00,00,00,00},//6
		[]int{00,00,00,22,8,00,00,00,00,00,00,00,00,00},//7
		[]int{00,00,00,00,19,34,00,00,00,00,00,00,00,00},//8
		[]int{00,00,00,00,00,15,25,00,00,00,00,00,00,00},//9
		[]int{00,00,00,00,00,00,00,9,62,54,00,00,00,00},//10
		[]int{00,00,00,00,00,00,00,00,00,00,23,00,00,00},//11
		[]int{00,00,00,00,00,00,00,00,00,00,32,00,00,00},//12
		[]int{00,00,00,00,00,00,00,00,00,00,00,00,00,00}}}
	m.Dijkstra()
}
