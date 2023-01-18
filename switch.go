package main 
import (
"fmt"
"time"
)


func main() {

for i:=0; i<7; i++{
switch {
case i == 0:
fmt.Println("Sunday")
case i == 1:
fmt.Println("Monday")
case i == 2:
fmt.Println("Tuesday")
case i == 3:
fmt.Println("Wednesday")
case i == 4:
fmt.Println("Thursday")
case i == 5:
fmt.Println("Friday")
case i == 6:
fmt.Println("Saturday")
}

}

switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
fmt.Println("It's a Weekend")

default:
fmt.Println("It's a Weekday")
}

Now := time.Now()
switch {
case Now.Hour() < 12:
fmt.Println("It's before noon")
default:
fmt.Println("It's after noon")
}

}
