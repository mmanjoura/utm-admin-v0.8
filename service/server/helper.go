package server

var ModeLookUp = map[string]string{
	"\u0001": "Self Test",
	"\u0002": "Commissioning",
	"\u0003": "Standard TRX",
	"\u0004": "Traffic Test",
}

var DiskSpaceLookUP = map[int]string{
	0: " Less 1GB",
	1: " More 1GB",
	2: " More 2GB",
	3: " More 4GB",
	4: " Max",
}

var EnergyLeftLookUP = map[int]string{
	0: " Less 5%",
	1: " Less 10%",
	2: " More 20%",
	3: " More 30%",
	4: " More 50%",
	5: " More 70%",
	6: " More 90%",
	7: " Max Left",
}
