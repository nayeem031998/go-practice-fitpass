package main
import "fmt"

type cuboid interface {

       area() float64
     volume() float64
  }   
     
type value struct {

      length float64
      width  float64
     height  float64
  }

func (m value) area() float64 {

  return 2*(m.length*m.width + m.length*m.height + m.width*m.height)

  }
  
func (m value) volume() float64 {

  return m.length*m.width*m.height
  
}   

func main () {

var t cuboid
  t = value{12, 8, 10}
  fmt.Println("Area of the cuboid:", t.area())
  fmt.Println("Volume of the cuboid:", t.volume())
}
