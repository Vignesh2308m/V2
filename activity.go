package main

import "fmt"

type dataset struct {
	data map[int]string
}

func NewDataSet(d map[int]string) *dataset {
	return &dataset{data: d}
}

func (d *dataset) apply(fn func(string) string) *dataset {
	for k, val := range d.data {
		d.data[k] = fn(val)
	}

	return d
}

func (d *dataset) filter(fn func(string) bool) *dataset {
	for k, val := range d.data {
		if !fn(val) {
			delete(d.data, k)
		}
	}
	return d
}

func (d *dataset) reduce(fn func(string, string) string) *dataset {
	res := make(map[int]string)

	fmt.Println(len(d.data))

	for key, val := range d.data {
		res[key] = fn(val, res[key])
	}
	d.data = res

	return d
}
