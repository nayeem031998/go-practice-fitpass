package main
import "fmt"

func main() {

var a [100]int
var total int

     fmt.Println("Enter total numbers : ")
     fmt.Scanf("%d", &total)
    
    for i:=0; i<total; i++ {
     
     fmt.Println("Enter the number: ")
     fmt.Scanf("%d", &a[i])

}
 
 for j:=1; j<total; j++ {
 
     if ( a[0] < a[j]) {
     
          a[0] = a[j]
     }
 
 }
     fmt.Println("largest number is ", a[0])
}





