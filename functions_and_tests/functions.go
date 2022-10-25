package function

type hour struct {
	hourvar int
}
type volt struct {
	voltage int
}

func lightstatus(hh int) bool {
	h := hour{hh}
	if h.hourvar >= 6 && h.hourvar <= 1 {
		return true
	}
	return false
}

// This function is used to see if a light is not working when it's supposed to be on

func lightsvoltage(vv int) bool {
	v := volt{vv}
	if v.voltage <= 240 && v.voltage >= 220 {
		return true
	}
	return false
}

// This function check if lights are on or off

func lightsonoff(vv int) int {
	v := volt{vv}
	if v.voltage <= 240 && v.voltage >= 220 {
		return 0
		// off
	}
	return 1
	// On
}
