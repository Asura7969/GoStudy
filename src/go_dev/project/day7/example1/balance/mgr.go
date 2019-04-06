package balance

import "fmt"

type BalanceMgr struct {
	allBalanace map[string]Balance
}

var mgr = BalanceMgr{
	allBalanace: make(map[string]Balance),
}

func (b *BalanceMgr) register(name string, ba Balance) {
	b.allBalanace[name] = ba
}

func RegisterBalancer(name string, b Balance) {
	mgr.register(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balancer, ok := mgr.allBalanace[name]
	if !ok {
		err = fmt.Errorf("Not found %s balancer", name)
		return
	}

	fmt.Printf("use %s balancer\n", name)
	inst, err = balancer.DoBalance(insts)
	return
}
