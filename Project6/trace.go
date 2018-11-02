package main

var table = [5][8]string{
	//  a	  +	     -		  *	      /	     (	    )	   $
	{"TQ", "NULL", "NULL", "NULL", "NULL", "TQ", "NULL", "NULL"}, // E
	{"NULL", "+TQ", "-TQ", "NULL", "NULL", "NULL", "1", "1"},     // Q
	{"FR", "NULL", "NULL", "NULL", "NULL", "FR", "NULL", "NULL"}, // T
	{"NULL", "1", "1", "*FR", "/FR", "NULL", "1", "1"},           // R
	{"a", "NULL", "NULL", "NULL", "NULL", "(E)", "NULL", "NULL"}, // F
}

func trace(popped string, read string) string {
	var row, col int8

	switch popped {
	case "E":
		row = 0
	case "Q":
		row = 1
	case "T":
		row = 2
	case "R":
		row = 3
	case "F":
		row = 4
	}

	switch read {
	case "a":
		col = 0
	case "+":
		col = 1
	case "-":
		col = 2
	case "*":
		col = 3
	case "/":
		col = 4
	case "(":
		col = 5
	case ")":
		col = 6
	case "$":
		col = 7
	}

	return table[row][col]
}
