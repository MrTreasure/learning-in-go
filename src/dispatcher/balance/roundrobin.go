package balance

import (
	"errors"
)

type RoundRobinBalance struct {
	curIndex int
}

func (p *RoundRobinBalance) DoBalance(list []*Instance) (inst *Instance, err error) {
	if len(list) == 0 {
		err = errors.New("No instance")
		return
	}

	lens := len(list)
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	inst = list[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens
	return
}
