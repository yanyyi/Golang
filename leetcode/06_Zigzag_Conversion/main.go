package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	down := false
	row := 0
	result := make(map[int]string)
	for i := 0; i < len(s); i++ {
		result[row] += string(s[i])
		if row == 0 || row == numRows-1 {
			down = !down
		}
		if down {
			row += 1
		} else {
			row -= 1
		}

	}
	for i := 1; i < numRows; i++ {
		result[0] += result[i]
	}
	return result[0]
}

//func main() {
//	s := "PAYPALISHIRING"
//	n := 3
//	fmt.Println(convert(s, n))
//}
