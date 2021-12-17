package bookstore

func Cost(books []int) int {
	min := len(books) * 800
	for _, combo := range possibilities(Combo{}, books) {
		if p := combo.price(); p < min {
			min = p
		}
	}
	return min
}

func possibilities(cur Combo, books []int) (res []Combo) {
	if len(books) == 0 {
		res = append(res, cur)
		return
	}
	x := books[0]
	xs := books[1:]

	inserted := false
	for i, set := range cur {
		if !set.Has(x) {
			next := cur.clone()
			next[i] = cur[i].Insert(x)
			res = append(res, possibilities(next, xs)...)
			inserted = true
		}
	}
	if !inserted {
		set := IntSet(0)
		set = set.Insert(x)
		next := append(cur.clone(), set)
		res = append(res, possibilities(next, xs)...)
	}
	return
}

type IntSet int

func (i IntSet) Len() (res int) {
	for ; 0 < i; i = i >> 1 {
		if i&1 == 1 {
			res++
		}
	}
	return
}

func (i IntSet) Has(x int) bool {
	return i>>x&1 == 1
}

func (i IntSet) Insert(x int) IntSet {
	return i | 1<<x
}

type Combo []IntSet

func (combo Combo) price() (res int) {
	for _, set := range combo {
		switch set.Len() {
		case 1:
			res += 800
		case 2:
			res += (800 - 40) * 2
		case 3:
			res += (800 - 80) * 3
		case 4:
			res += (800 - 160) * 4
		case 5:
			res += (800 - 200) * 5
		}
	}
	return
}

func (combo Combo) clone() Combo {
	res := make([]IntSet, len(combo))
	copy(res, []IntSet(combo))
	return Combo(res)
}
