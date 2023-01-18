package main 
import "fmt"

func main() {
 
var total int
 
var a [100]int

var b int

var avg int

fmt.Printf("Enter total number of weights : \n")
fmt.Scanf("%d", &total)
 
for i:=0; i<total; i++ {

fmt.Println("Enter the weight : \n")
fmt.Scanf("%d", &a[i])

   b += a[i]
}

   avg = (b/total)

fmt.Printf("Average weight of %d is : %d " , total , avg )
fmt.Println()

}







