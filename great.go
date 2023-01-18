package main 
import "fmt"

func main() {

var x int
a := [6]int{8,4,24,12,43,78}
for index:= range a {
	if x<a[index]{
		x=a[index]
	}
	
}
fmt.Println(x)
}

