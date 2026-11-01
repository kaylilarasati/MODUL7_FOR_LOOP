package main

import "fmt"
func main() {
	var x, y, j, jumlah int
	fmt.Scan(&x, &y)
	jumlah = 0
	for j = x; j <= y; j++ {
		jumlah = jumlah + j
	}
	fmt.Println(jumlah)
}