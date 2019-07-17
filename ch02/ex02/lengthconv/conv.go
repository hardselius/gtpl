package lengthconv

func MToFt(m Meter) Foot {
	return Foot(m * 0.3048)
}

func FtToM(ft Foot) Meter {
	return Meter(ft / 0.3048)
}
