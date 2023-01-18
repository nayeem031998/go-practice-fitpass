package main
import "fmt"

func main() {

var f float64
var c float64

fmt.Println("Enter the temperature in Fahrenheit : ")
fmt.Scanf("%g", &f)

c = (f-32)*0.556

fmt.Println("Temperature in Celcius : ", c)

}








