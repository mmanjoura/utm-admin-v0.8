package server

var ModeLookUp = map[string]string{
	"\u0001": "Self Test",
	"\u0002": "Commissioning",
	"\u0003": "Standard TRX",
	"\u0004": "Traffic Test",
}

var DiskSpaceLookUP = map[int]string{
	0: " < 1GB",
	1: " > 1GB",
	2: " > 2GB",
	3: " > 4GB",
	4: " max",
}

var EnergyLeftLookUP = map[int]string{
	0: " > 5%",
	1: " < 10%",
	2: " > 20%",
	3: " > 30%",
	4: " > 50%",
	5: " > 70%",
	6: " > 90%",
	7: " max left",
}
