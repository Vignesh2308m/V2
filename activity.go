package main

type dataset struct {
	data map[int]string
}

func NewDataSet(d map[int]string) *dataset {
	return &dataset{data: d}
}

func (d *dataset) apply(fn func(string) string) {
	for k, val := range d.data {
		d.data[k] = fn(val)
	}
}
