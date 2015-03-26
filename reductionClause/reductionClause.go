package main

import (
	"fmt"
	"runtime"
	. "github.com/acastano/GOpenMP/gomp_lib"
)

var n=100
var wch = make(chan int)
var sum =5
func main() {
//pragma gomp parallel default(none) shared(sum,a,m,wch) reduction(+:sum)
	Gomp_set_num_routines(4)	
	runtime.GOMAXPROCS(12)
	g:=Gomp_get_num_routines()
	for i := 0; i < g; i++ {
		go func(b int) {
			sum:=0
			fmt.Println("value of sum in goroutine",b,"is: ",sum)
			sum++
			wch <- sum
		}(i)
	}

	for i := 0; i < g; i++ {
		sum += <-wch
	}
	fmt.Println("value of sum after parallelization is: ",sum)
}
