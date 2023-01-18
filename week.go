package main 
import (
"fmt"
"time"
)

func main() {

switch time.Now().Weekday() {
case time.Saturday, time.Sunday :
fmt.Println("weekend")

default :
fmt.Println("weekday")

}

}



