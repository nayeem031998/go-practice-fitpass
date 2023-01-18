package main 
import (
         "fmt"
         "math"
)

func xyz(b float64) float64{
if b < 0 {
        return -b
}

        return b
}


func main() {

var a float64
a = -632.873

fmt.Printf("Absolute value of %f is : %f \n" , a , math.Abs(a))

var b float64
b = -83278323

fmt.Printf("Absolute value of %d is : %d \n", b , xyz(b))

}





