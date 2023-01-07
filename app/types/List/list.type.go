package List

import "sort"

type Number interface {
	int | int64 | float32 | float64
}

func Sum[V Number](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func Min[V Number](m []V) V {
	var s V = m[0]
	for _, v := range m {
		if s > v {
			s = v
		}
	}
	return s
}

func Max[V Number](m []V) V {
	var s V = m[0]
	for _, v := range m {
		if s < v {
			s = v
		}
	}
	return s
}

func Remove[V Number](m []V, index int) []V {
	return append(m[:index], m[index+1:]...)
}

func Push[V Number](m []V, item V) []V {
	return append(m, item)
}

func Pop[V Number](m []V) []V {
	return append(m[:len(m)-1])
}

func Sort[V Number](m []V) []V {
	sort.Slice(m, func(i, j int) bool {
		return m[i] < m[j]
	})
	return m
}
