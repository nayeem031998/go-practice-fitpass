package main 
import "fmt"

func main() {

abc := [7]int{1, 22, 13, 64, 85, 16, 37}  

l1 := abc[0]
                                            //To print second largest number
var l2 int

for j:=1; j<7; j++ {
 
 if l1 < abc[j] {
 
  l2 = l1
  l1 = abc[j]
  
 } else if l2 < abc[j] {
 
  l2 = abc[j]
 
 }

}
  
   fmt.Println(l2)
}
