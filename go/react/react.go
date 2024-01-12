package react

import "math/rand"

type reactor struct {
	computeCells []*cell
}

func New() Reactor {
	return &reactor{}
}

func (r *reactor) CreateInput(initial int) InputCell {
	return &cell{val: initial, reactor: r}
}

func (r *reactor) CreateCompute1(cell1 Cell, compute func(int) int) ComputeCell {
	return r.newCompute(func() int { return compute(cell1.Value()) })
}

func (r *reactor) CreateCompute2(cell1, cell2 Cell, compute func(int, int) int) ComputeCell {
	return r.newCompute(func() int { return compute(cell1.Value(), cell2.Value()) })
}

func (r *reactor) newCompute(compute func() int) ComputeCell {
	cell := cell{reactor: r, val: compute(), compute: compute, callbacks: map[uint64]func(int){}}
	r.computeCells = append(r.computeCells, &cell)
	return &cell
}

type cell struct {
	reactor   *reactor
	val       int
	compute   func() int
	callbacks map[uint64]func(int)
}

func (c *cell) Value() int {
	return c.val
}

func (c *cell) SetValue(val int) {
	c.val = val
	for _, cc := range c.reactor.computeCells {
		if new := cc.compute(); new != cc.val {
			cc.val = new
			for _, callback := range cc.callbacks {
				callback(cc.val)
			}
		}
	}
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	key := rand.Uint64()
	c.callbacks[key] = callback
	return &canceler{func() { delete(c.callbacks, key) }}
}

type canceler struct {
	cancel func()
}

func (c *canceler) Cancel() {
	c.cancel()
}
