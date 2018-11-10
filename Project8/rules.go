package main

const R1 string = "E+T"	// E -> E+T
const R2 string = "E-T"	// E -> E-T
const R3 string = "T"		// E -> T
const R4 string = "T*F"	// T -> T*F
const R5 string = "T/F"	// T -> T/F
const R6 string = "F"		// T -> F
const R7 string = "(E)"	// F -> (E)
const R8 string = "i"		// F -> i

func getRuleLengthAndNonTerm(rule string) (int, string) {
	switch rule {
	case "R1":
		return len(R1) * 2, "E"
	case "R2":
		return len(R2) * 2, "E"
	case "R3":
		return len(R3) * 2, "E"
	case "R4":
		return len(R4) * 2, "T"
	case "R5":
		return len(R5) * 2, "T"
	case "R6":
		return len(R6) * 2, "T"
	case "R7":
		return len(R7) * 2, "F"
	case "R8":
		return len(R8) * 2, "F"
	default:
		return 0, ""
	}
}