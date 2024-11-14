package unboundedqueue

import "sync"

type UnboundedQueen struct {
	queen []interface{}
	l     sync.Mutex
	cond  *sync.Cond
}

func NewUnboundedQueen() *UnboundedQueen {
	uq := UnboundedQueen{}
	uq.queen = make([]interface{}, 0)
	uq.cond = sync.NewCond(&uq.l)
	return &uq
}

func (uq *UnboundedQueen) Push(i interface{}) {
	uq.l.Lock()
	uq.queen = append(uq.queen, i)
	uq.l.Unlock()
	uq.cond.Broadcast()
}

func (uq *UnboundedQueen) PopBlock() interface{} {
	uq.l.Lock()
	for {
		if len(uq.queen) == 0 {
			uq.cond.Wait()
		} else {
			ret := uq.queen[0]
			uq.queen = uq.queen[1:]
			uq.l.Unlock()
			return ret
		}
	}
}
