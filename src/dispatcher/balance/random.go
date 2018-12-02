package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(list []*Instance) (inst *Instance, err error) {
	if len(list) == 0 {
		err = errors.New("No instance")
		return
	}

	lens := len(list)
	index := rand.Intn(lens)
	inst = list[index]
	return
}
