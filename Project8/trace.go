package main

import "strconv"

var parse_table = [16][11]string {
// Row and column:
// i		 +		-		*		/		(		 )		$		E		 T		F
	{"S5", "", "", "", "", "S4", "", "", "1", "2", "3"},			// 0
	{"", "S6", "S7", "", "", "", "", "ACC", "", "", ""},			// 1
	{"", "R3", "R3", "S8", "S9", "", "R3", "R3", "", "", ""},	// 2
	{"", "R6", "R6", "R6", "R6", "", "R6", "R6", "", "", ""},	// 3
	{"S5", "", "", "", "", "S4", "", "", "10", "2", "3"},			// 4
	{"", "R8", "R8", "R8", "R8", "", "R8", "R8", "", "", ""},	// 5
	{"S5", "", "", "", "", "S4", "", "", "", "11", "3"},			// 6
	{"S5", "", "", "", "", "S4", "", "", "", "12", "3"},			// 7
	{"S5", "", "", "", "", "S4", "", "", "", "", "13"},				// 8
	{"S5", "", "", "", "", "S4", "", "", "", "", "14"},				// 9
	{"", "S6", "S7", "", "", "", "S15", "", "", "", ""},			// 10
	{"", "R1", "R1", "S8", "S9", "", "R1", "R1", "", "", ""},	// 11
	{"", "R2", "R2", "S8", "S9", "", "R2", "R2", "", "", ""},	// 12
	{"", "R4", "R4", "R4", "R4", "", "R4", "R4", "", "", ""},	// 13
	{"", "R5", "R5", "R5", "R5", "", "R5", "R5", "", "", ""},	// 14
	{"", "R7", "R7", "R7", "R7", "", "R7", "R7", "", "", ""},	// 15
}

func trace(popped string, read string) string {
	poppedToInt, _ := strconv.Atoi(popped)
	row := int8(poppedToInt)

	switch read {
	case "i":
		return parse_table[row][0]
	case "+":
		return parse_table[row][1]
	case "-":
		return parse_table[row][2]
	case "*":
		return parse_table[row][3]
	case "/":
		return parse_table[row][4]
	case "(":
		return parse_table[row][5]
	case ")":
		return parse_table[row][6]
	case "$":
		return parse_table[row][7]
	case "E":
		return parse_table[row][8]
	case "T":
		return parse_table[row][9]
	case "F":
		return parse_table[row][10]
	default:
		return "NULL"
	}
}