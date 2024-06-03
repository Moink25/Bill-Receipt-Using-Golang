package main
import (
	"fmt"
	"os"
)
type bill struct{
	name string
	items map[string]float64
	tip float64
}

func newBill(n string) bill{
	b:= bill{
		name:n,
		items: map[string]float64{},
    tip:0,
	}
	return b
}

func (b *bill) format() string{ //receiver function for the bill structure and cannot be accessed without bill struct
	//automatic dereferencing of the pointer in the function,we dont need to do
	var fs string =""
	fs+=fmt.Sprintf(b.name+"\n")
	fs+=fmt.Sprintf("Bill breakdown\n")
  
	var total float64=0
  for k,v:=range b.items{
		fs+=fmt.Sprintf("%-25v...$%0.2f\n",k+":",v)
		total+=v
	}
	fs+=fmt.Sprintf("%-25v...$%0.2f\n","tip:",b.tip)
	fs+=fmt.Sprintf("%-25v...$%0.2f\n","total:",total+b.tip)
	return fs
}

func (b *bill) AddItems(name string,price float64){
	b.items[name]=price
}

func(b *bill) UpdateTip(tip float64){
	b.tip=tip
}

//save bill
func (b *bill) SaveBill(){
	data:=[]byte(b.format())
	err:=os.WriteFile("./"+b.name+".txt",data,0644)//location, byte-version, permissions
	// err:=os.WriteFile("bill/"+b.name+".txt",data,0644)//location, byte-version, permissions
	if err != nil{
		panic(err)//stop the flow and print error
	}
	fmt.Println("The bill is saved")
}