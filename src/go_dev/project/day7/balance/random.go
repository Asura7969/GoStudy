package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {

}

func init()  {
	RegisterBalancer("Random",&RandomBalance{})
	
}

func (r *RandomBalance)DoBalance(insts []*Instance, key ...string)(inst *Instance,err error){

	if len(insts) < 1 {
		err = errors.New("No Instance")
		return
	}

	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}
