package main

import (
	"fmt"
	. "github.com/acastano/GOpenMP/gomp_lib"
)


var a, b, c = 1, 2, 3

func main() {
	fmt.Println("values of a,b and c before the parallelization: ", a, b, c)
	//pragma gomp parallel(5) private (a,b,c)
	ch := make(chan int)
	Gomp_set_num_routines(5)
	for i := 0; i < Gomp_get_num_routines(); i++ {
		go func(y int) { 
			var a, b, c int
			fmt.Println("values of a,b and c in parallelization :", a, b, c, "Goroutine number: ", Gomp_get_routine_num())
			
			ch <- 0
		}(i)
	}
	for i := 0; i < Gomp_get_num_routines(); i++ {
		<-ch
	}
}
