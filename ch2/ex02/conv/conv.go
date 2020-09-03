package conv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func MToFt(m Meter) Feet {
	return Feet(m * 3.2808)
}

func FtToM(ft Feet) Meter {
	return Meter(ft / 3.2808)
}

func KgToLb(kg Kilogram) Pound {
	return Pound(kg / 0.45359237)
}

func LbToKg(lb Pound) Kilogram {
	return Kilogram(lb * 0.45359237)
}
