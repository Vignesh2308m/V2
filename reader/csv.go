package reader

import (
	"encoding/csv"
)


type csvfile struct {
	f_path string
}

type csv_context struct{
	
}

func (c *csvfile) CSV(f_path, sep, end string) *csvfile {
	return &csvfile{
		f_path: c.f_path,
		sep:    c.sep,
		end:    c.end,
	}
}

func (c *)

func (c *csvfile) Row_Count() int{

}