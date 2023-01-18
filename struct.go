package main
import "fmt"

type a struct{
    name string
    height string
    weight string
    age string
    contact string
    membership b
}

type b struct{
    fitpass string
    fitfeast string
    fitcoach string
    fitprime string         
}

func main() {

     c := a{membership:b{fitpass: "Active", fitfeast: "Inactive", fitcoach: "Active", fitprime: "Active"}}
             fmt.Println(c)
             
     d := a{name: "Name - Hemant\n" , height: "Height - 173\n" , weight: "Weight - 75\n" , age: "Age - 31\n" , contact: "Contact - 9826373871"}
             fmt.Println(d)
             
             fmt.Println("fitpass membership:", c.membership.fitpass, 
                         "\nfitfeast membership:", c.membership.fitfeast,
                         "\nfitcoach membership:", c.membership.fitcoach,
                         "\nfitprime membership:", c.membership.fitprime)
     f:= &d
     fmt.Println(f)                        
      d.name = "Akash"  
      fmt.Println(d.name)                 
              fmt.Println(d.name, d.contact) 
     e := b{fitpass: "Active", fitfeast: "Inactive", fitcoach: "Inactive", fitprime: "Inactive"}
     fmt.Println(e)                  
}

