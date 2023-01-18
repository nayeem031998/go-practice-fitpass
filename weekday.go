package main 
import "fmt"

func main() {

for days := 1; days <= 7; days++ {
fmt.Printf("%d", days)

switch {
case days == 1:
fmt.Println("Monday") 

case days == 2:
fmt.Println("Tuesday")

case days == 3:
fmt.Println("Wednesday")
 
case days == 4:
fmt.Println("Thursday")

case days == 5:
fmt.Println("Friday")

case days == 6:
fmt.Println("Saturday")

case days == 7:
fmt.Println("Sunday")

 } 
 }
 }
