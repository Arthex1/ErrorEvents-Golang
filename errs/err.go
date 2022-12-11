package errs

import (
	"fmt"

	"golang.org/x/exp/slices"
)


type ErrorHandler struct {
	Listeners map[uint][]func(err *Error)
}                            
type Error struct {
	Name string 
	ID uint 
	Message string 
	
}
// Returns a new `Error` Struct Object with the Provided Message, Warning: ID and Error Name remain the same as before. This can be used while injecting variables into the message.
func (s *Error) ChangeMessage(msg string) *Error {
	object := s 
	
	object.Message = msg 
	return object

}

func (s *ErrorHandler) GetListeners(ecode uint) int  {
	return len(s.Listeners[ecode]) 
	 
	
}
func (s *ErrorHandler) GetAllListeners() int  {
	return len(s.Listeners) 
	 
	
}
func (s *ErrorHandler) DisconnectAllListeners(ecode uint) {
	delete(s.Listeners, ecode)  
}
// Avoid using this function at all costs, as it can delete unwanted listeners if the index is provided wrong
func (s *ErrorHandler) DisconnectListener(ecode uint, index int) {
	slices.Delete(s.Listeners[ecode], index, index) 
}


func (s *ErrorHandler) Listen(ecode uint, handler func(err *Error))  {
	
	_, exists := s.Listeners[ecode]
	if !exists {
		s.Listeners[ecode] = make([]func(err *Error), 0) 

	}
	listeners := s.Listeners[ecode] 
	s.Listeners[ecode] = append(listeners, handler)

	
}

var (
	Account = []*ErrorHandler{}
)



func (s *ErrorHandler) Init() {
	Account = append(Account, s)
	return 
}





func (s *ErrorHandler) Emit(ecode uint, err *Error) {
	listeners, exists := s.Listeners[ecode]
	if !exists {
		return 
	}
	for _, l := range listeners {
		l(err)

	}
}

func Emit(err *Error) error {  
	for _, v := range Account {
		v.Emit(err.ID, err) 

	}
	return fmt.Errorf(err.Message) 

}


