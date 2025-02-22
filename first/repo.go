package first

type Repo struct {
	storage map[int]int
	unique  map[int]struct{}
}

func (r *Repo) init() {
	r.storage = make(map[int]int)
	r.unique = make(map[int]struct{})
}

func (r *Repo) add(x int) {
	r.storage[x]++

	if counter := r.storage[x]; counter == 1 {
		r.unique[x] = struct{}{}
	} else {
		delete(r.unique, x)
	}
}

func (r *Repo) get_unique() int {
	for key, _ := range r.unique {
		return key
	}

	return 0
}

func (r *Repo) delete(x int) {
	if _, ok := r.unique[x]; ok {
		delete(r.unique, x)
		delete(r.storage, x)
	} else if counter, ok := r.storage[x]; ok {
		r.storage[x]--

		if counter-1 == 1 {
			r.unique[x] = struct{}{}
		}
	}
}
