package tempconv

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenehit temperature to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k) + AbsoluteZeroC }

// CToK converts a Celsius temperature to Kelvin
func CToK (c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }
