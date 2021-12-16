package forth

import "errors"
import "strconv"
import "strings"

func Forth(input []string) ([]int, error) {
	runner := newRunner()
	for _, line := range input {
		for _, s := range strings.Split(line, " ") {
			runner.run(strings.ToLower(strings.TrimSpace(s)))
			if runner.error != nil {
				return runner.stack, runner.error
			}
		}
	}
	return runner.stack, nil
}

type Stack = []int

type Defs = map[string][]Lit

type Runner struct {
	stack    Stack
	defs     Defs
	recorder *Recorder
	error    error
}

func newRunner() (runner Runner) {
	runner.stack = Stack{}
	runner.defs = Defs{
		"+":    []Lit{LitFn{func(x, y int) int { return x + y }}},
		"-":    []Lit{LitFn{func(x, y int) int { return x - y }}},
		"*":    []Lit{LitFn{func(x, y int) int { return x * y }}},
		"/":    []Lit{LitFn{func(x, y int) int { return x / y }}},
		"dup":  []Lit{LitMani{1, func(xs Stack) Stack { return append(xs, xs...) }}},
		"drop": []Lit{LitMani{1, func(xs Stack) Stack { return Stack{} }}},
		"swap": []Lit{LitMani{2, func(xs Stack) Stack { return Stack{xs[1], xs[0]} }}},
		"over": []Lit{LitMani{2, func(xs Stack) Stack { return append(xs, xs[0]) }}},
	}
	return
}

func (runner *Runner) run(s string) {
	defer func() {
		err := recover()
		if err != nil {
			runner.error = errors.New("Unknown error")
		}
	}()

	runner.push(parse(s))
}

func (r *Runner) push(l Lit) {
	if r.recorder != nil {
		if _, b := l.(LitRecEnd); b {
			key, inst := r.recorder.expand(r.defs)
			r.defs[key] = inst
			r.recorder = nil
		} else {
			r.recorder.push(l)
		}
		return
	}
	stack := r.stack
	switch x := l.(type) {
	case LitNum:
		r.stack = append(stack, x.num)
	case LitFn:
		xy := stack[len(stack)-2:]
		v := x.fn(xy[0], xy[1])
		r.stack = append(stack[:len(stack)-2], v)
	case LitMani:
		xs := stack[len(stack)-x.n:]
		r.stack = append(stack[:len(stack)-x.n], x.fn(xs)...)
	case LitRecStart:
		r.recorder = &Recorder{}
	case LitStr:
		if inst, b := r.defs[x.str]; b {
			for _, l := range inst {
				r.push(l)
			}
		} else {
			r.error = errors.New("Unknown word")
		}
	}
}

type Recorder struct {
	lits []Lit
}

func (r *Recorder) push(l Lit) {
	r.lits = append(r.lits, l)
}

func (r *Recorder) expand(defs Defs) (key string, inst []Lit) {
	if x, b := r.lits[0].(LitStr); b {
		key = x.str
	} else {
		panic("Invalid word")
	}
	for _, l := range r.lits[1:] {
		if x, b := l.(LitStr); b {
			if inst_, b := defs[x.str]; b {
				inst = append(inst, inst_...)
			}
		} else {
			inst = append(inst, l)
		}
	}
	return
}

type Lit interface{}

type LitNum struct {
	num int
}

type LitStr struct {
	str string
}

type LitFn struct {
	fn func(int, int) int
}

type LitMani struct {
	n  int
	fn func(Stack) Stack
}

type LitRecStart struct{}

type LitRecEnd struct{}

func parse(x string) Lit {
	switch x {
	case ":":
		return LitRecStart{}
	case ";":
		return LitRecEnd{}
	}
	if n, e := strconv.Atoi(x); e == nil {
		return LitNum{n}
	}
	return LitStr{x}
}
