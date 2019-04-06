package balance

import (
	"errors"
)

type RoundRobinBalance struct {
	curIndex int
}

func init() {
	RegisterBalancer("roundrobin", &RoundRobinBalance{})

}
func (r *RoundRobinBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) < 1 {
		err = errors.New("No Instance")
		return
	}

	lens := len(insts)
	if r.curIndex > lens {
		r.curIndex = 0
	}

	inst = insts[r.curIndex]
	r.curIndex = (r.curIndex + 1) % lens
	return
}
