package main 
import "fmt"

type square struct {
      s int
   geometry struct{
     area int
    }
}

type circle struct {
      r float64
   geometry struct{
     area float64
    }
  }
  
type rectangle struct {
      l int
      b int
    geometry struct{
      area int
       }
}

type cuboid struct {
      l int
      b int
      h int
    geometry struct{
      area int
    }
}

type tank struct {
      h float64
      r float64 
    geometry struct{
      area float64
    }  
}

func main () {
    var sq square
    sq.s = 10
    sq.geometry.area = sq.s * sq.s
         
         fmt.Println("Area of Square:", sq)
 
    var cr circle
    cr.r = 5
    cr.geometry.area = 3.14*cr.r*cr.r
    
         fmt.Println("Area of Circle:", cr) 
 
    var rec rectangle
    rec.l = 20
    rec.b = 5
    rec.geometry.area = rec.l*rec.b
    
         fmt.Println("Area of Rectangle:", rec)
 
    var cub cuboid
    cub.l = 20
    cub.b = 5
    cub.h = 3
    cub.geometry.area = cub.l*cub.b*cub.h
          
         fmt.Println("Area of Cuboid:", cub)
     
    var tnk tank
    tnk.h = 10
    tnk.r = 5
    tnk.geometry.area = 2*3.14*5*5 + 2*3.14*5*10 
    
         fmt.Println("Area of Tank:", tnk)
         

}
