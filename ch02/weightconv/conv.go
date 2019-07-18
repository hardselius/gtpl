package weightconv

func LbToKg(lb Pound) Kilogram {
	return Kilogram(lb * 0.45359237)
}

func KgToLb(kg Kilogram) Pound {
	return Pound(kg / 0.45359237)
}
