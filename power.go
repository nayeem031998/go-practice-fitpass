package main
import "fmt"

func main() {

var a int
var b int

fmt.Printf("Enter the number :")
fmt.Scanf("%d", &a)

fmt.Printf("Enter power to the number :")
fmt.Scanf("%d", &b)

c := a

for i:=1; i<b; i++{

     a = a*c 
}

fmt.Println(a)

}




