package server

var ModeLookUp = map[string]string{
	"\u0001": "Self Test",
	"\u0002": "Commissioning",
	"\u0003": "Standard TRX",
	"\u0004": "Traffic Test",
}

var DiskSpaceLookUP = map[int]string{
	0: " less 1GB",
	1: " more 1GB",
	2: " more 2GB",
	3: " more 4GB",
	4: " max",
}

var EnergyLeftLookUP = map[int]string{
	0: " less 5%",
	1: " less 10%",
	2: " more 20%",
	3: " more 30%",
	4: " more 50%",
	5: " more 70%",
	6: " more 90%",
	7: " max left",
}
