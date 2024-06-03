package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string,r *bufio.Reader)(string,error){
  fmt.Printf(prompt)
	input,err:=r.ReadString('\n')
	return strings.TrimSpace(input),err
}
func createBill() bill {
  reader:=bufio.NewReader(os.Stdin)
	// fmt.Printf("Create the bill for :")
	// name,_:=reader.ReadString('\n')
	// name=strings.TrimSpace(name)
	name,_:=getInput("Create the bill for: ",reader)
	b:=newBill(name)
	return b
}
func promptOpt(b *bill){
	reader:=bufio.NewReader(os.Stdin)
	opt,_:=getInput("Choose the options:\na)Add items\ns)Save bill\nt)Add tip\n",reader)
	switch opt {
	case "a":
	         name,_:=getInput("What's the item? ",reader)
					 price,_:=getInput("Add the price ",reader)
					 p,err:=strconv.ParseFloat(price,64)
					 if err != nil{
						fmt.Println("The price must be a number")
						promptOpt(b)
					 } 
					 
            b.AddItems(name,p)
					  promptOpt(b)
					//  fmt.Println(name,price)
	         break
	case "s":b.SaveBill()
		       fmt.Println("You saved the file- ",b.name)
	         break
	case "t":tip,_:=getInput("Enter the tip amount ",reader)
	         fmt.Println(tip)
					 t,err:=strconv.ParseFloat(tip,64)
					 if err != nil{
						fmt.Println("The tip should be number")
						promptOpt(b)
					 } 
						b.UpdateTip(t)
					 promptOpt(b)
	         break
	default: fmt.Println("Invalid choice")
	         promptOpt(b)
	}
}
 
func main(){
	  myBill:=createBill()
		promptOpt(&myBill)
	// myBill:=newBill("Moin khan")
	// myBill.AddItems("Tea",30.5)
	// myBill.AddItems("Idli",80.90)
	// myBill.AddItems("Toffee",20)
	// myBill.AddItems("Soup",90.99)
	// myBill.UpdateTip(20)
	fmt.Println(myBill.format())
	
	// fmt.Println(myBill)
}