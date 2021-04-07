package service

type pokemonsIDSorter []int

func (pis pokemonsIDSorter) Len() int { return len(pis) }

func (pis pokemonsIDSorter) Less(i, j int) bool { return pis[i] < pis[j] }

func (pis pokemonsIDSorter) Swap(i, j int) { pis[i], pis[j] = pis[j], pis[i] }
