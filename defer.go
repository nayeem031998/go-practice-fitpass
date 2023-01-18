package main
import "fmt"

func add(a , b int) int {

  result := a + b
        
        fmt.Printf("Addition of %d and %d = %d  \n", a , b , result)

  return 0
}


func main() {

fmt.Println("start")

defer fmt.Println("End")

defer add(78 , 98)

defer add(56 , 32)

}



