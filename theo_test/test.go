package test

type hour struct {
	hourvar int
}

func light1(hh int) bool {
	h := hour{hh}
	if h.hourvar >= 6 && h.hourvar <= 1 {
		return true
	}
	return false
}
