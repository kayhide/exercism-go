package react

// import "fmt"

type reactor struct {
	deps    map[Cell][]*cell
	changes map[*cell]int
}

type cell struct {
	reactor   *reactor
	value     int
	update    func()
	callbacks []*func(int)
}

type canceler struct {
	cell *cell
	f    *func(int)
}

func (c *canceler) Cancel() {
	cbs := []*func(int){}
	for _, g := range c.cell.callbacks {
		if c.f != g {
			cbs = append(cbs, g)
		}
	}
	c.cell.callbacks = cbs
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) SetValue(value int) {
	c.reactor.changes = map[*cell]int{}
	if c.value != value {
		c.value = value
		c.reactor.fire(c)
	}
	for c, v := range c.reactor.changes {
		cur := c.Value()
		if cur != v {
			for _, cb := range c.callbacks {
				(*cb)(cur)
			}
		}
	}
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	c.callbacks = append(c.callbacks, &callback)
	return &canceler{c, &callback}
}

func New() Reactor {
	deps := map[Cell][]*cell{}
	changes := map[*cell]int{}
	return &reactor{deps, changes}
}

func (r *reactor) CreateInput(initial int) InputCell {
	c := cell{}
	c.reactor = r
	c.value = initial
	return &c
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	c := cell{}
	c.reactor = r
	c.callbacks = []*func(int){}
	c.update = func() {
		v := compute(dep.Value())
		if c.value != v {
			if _, b := c.reactor.changes[&c]; !b {
				c.reactor.changes[&c] = c.value
			}
			c.value = v
			c.reactor.fire(&c)
		}
	}
	c.update()
	if _, b := r.deps[dep]; !b {
		r.deps[dep] = []*cell{}
	}
	r.deps[dep] = append(r.deps[dep], &c)
	return &c
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	c := cell{}
	c.reactor = r
	c.callbacks = []*func(int){}
	c.update = func() {
		v := compute(dep1.Value(), dep2.Value())
		if c.value != v {
			if _, b := c.reactor.changes[&c]; !b {
				c.reactor.changes[&c] = c.value
			}
			c.value = v
			c.reactor.fire(&c)
		}
	}
	c.update()
	if _, b := r.deps[dep1]; !b {
		r.deps[dep1] = []*cell{}
	}
	r.deps[dep1] = append(r.deps[dep1], &c)
	if _, b := r.deps[dep2]; !b {
		r.deps[dep2] = []*cell{}
	}
	r.deps[dep2] = append(r.deps[dep2], &c)
	return &c
}

func (r *reactor) fire(c Cell) {
	for _, d := range r.deps[c] {
		if d.update != nil {
			d.update()
		}
	}
}
