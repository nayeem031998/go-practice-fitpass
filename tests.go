package main
import "fmt"

type fitpassHead struct {

     FP_ID      int
     gender     string    
     first_name string
     last_name  string
     contact    int
     email      string
          
}

func (f fitpassHead) fullname() string {
      return fmt.Sprintf("%s %s", f.first_name, f.last_name) 
}

type details struct {

     fitpassHead
     weight    float64
     height_Feet    int         
     height_Inches  int 
     age       int
     waist     int 
     fitpass   string
     fitfeast  string
     fitcoach  string   
   
}

func (d details) data() {

   fmt.Println("FITPASS ID:",d.FP_ID )
   fmt.Println("Gender:", d.gender)
   fmt.Println("Full Name:", d.fullname())
   fmt.Println("Contact:", d.contact)
   fmt.Println("Email_ID:", d.email)
   fmt.Println("Weight:", d.weight)
   fmt.Println("Height:", d.height_Feet, "'", d.height_Inches, "\" ")
   fmt.Println("Age:", d.age)
   fmt.Println("Waist:", d.waist)
   fmt.Println("FITPASS:", d.fitpass)
   fmt.Println("FITFEAST:", d.fitfeast)
   fmt.Println("FITCOACH:", d.fitcoach)
   
}

func main() {

fit := fitpassHead{1, "male", "Jnanendra", "Veer", 9717024007, "jnanendra.veer@fitpass.co.in"}

value := details{fit, 72.60, 5, 8, 34, 30, "Active", "Inactive", "Active"}

      value.data()
}
