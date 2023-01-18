package main
import "fmt"

func main() {

days := []string{"Monday" , "Tuesday" , "Wednesday" , "Thursday" , "Friday" , "Saturday" , "Sunday"}

fmt.Println(days)

for i:=0; i<7; i++ {
fmt.Println(days[i])
 }
 
 a := [5]int{1, 2, 3, 4, 5}
 
 fmt.Println(a[1], a[2])
 fmt.Println(a)
 
  var twoD [4][4]int
    for i := 0; i <= 3; i++ {
        for j := 0; j < 4; j++ {
            twoD[i][j] = i + j
        }
    }
 fmt.Println(twoD)
    
 emp := [6]string{"a", "b", "c", "d", "e"}  

 fmt. Println(emp)
 fmt.Println(len(emp))
 fmt.Println(cap(emp))
 
 }

