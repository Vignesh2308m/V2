package reader

type csvfile struct {
	f_path string
	sep    string
	end    string
}

func (c *csvfile) CSV(f_path, sep, end string) *csvfile {
	return &csvfile{
		f_path: c.f_path,
		sep:    c.sep,
		end:    c.end,
	}
}
