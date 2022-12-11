package main

import (
	"events/errs"
	"fmt"
	"log"

	
)

// Warning!! Having Two errors with the same id will overwrite the former one.
var (
	AddrNotFound = errs.Error{Name: "AddrNotFound", ID: 501, Message: "Address Not Found"}
	DBNotExists = errs.Error{Name: "DBotExists", ID: 502, Message: "Does not exist on the database."}
)

// Setup Error Handler
func setup() errs.ErrorHandler {
	tp := errs.ErrorHandler{Listeners: map[uint][]func(err *errs.Error){}}; tp.Init()
	return tp 
}

var (
	ErrorHandler = setup() 
	Throw = errs.Emit
)



func main() {
	
	ErrorHandler.Listen(502, func(err *errs.Error) { 
		fmt.Println(err.Name, err.ID) 
		
	})
	ErrorHandler.Listen(501, func(err *errs.Error){
		fmt.Println(err.Name, err.ID)  
		
		
	})
	
	
	
	if pp := Test(); pp != nil {
		log.Fatal("eee")
	}	 

		

	
}

func Test() error {
	return Throw(AddrNotFound.ChangeMessage("Elemental Goddess does not exist in the database") )
}
