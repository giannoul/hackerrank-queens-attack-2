package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/* Helper function for absolute value*/
func abs(i1, i2 int32) int32 {
	res := i1 - i2
	if res > 0 {
		return res
	}
	return -1 * res
}

/* Diagonal positions for diagonal "\" */
func diag1(qx, qy, s int32) [][]int32 {
	pos := make([][]int32, 0)
	for i := int32(1); i <= s-qx; i++ {
		pos = append(pos, []int32{qx + i, qy - i})
	}
	for i := int32(1); i <= s-qy; i++ {
		pos = append(pos, []int32{qx - i, qy + i})
	}
	return pos
}

/* Diagonal positions for diagonal "/" */
func diag2(qx, qy, s int32) [][]int32 {
	pos := make([][]int32, 0)
	for i := int32(1); i <= s-qy; i++ {
		pos = append(pos, []int32{qx + i, qy + i})
	}
	for i := int32(1); i < qx; i++ {
		pos = append(pos, []int32{qx - i, qy - i})
	}
	return pos
}

/*
 * Complete the 'queensAttack' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER r_q
 *  4. INTEGER c_q
 *  5. 2D_INTEGER_ARRAY obstacles
 */
func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	remx, remy := n-1, n-1
	d1pos := diag1(r_q, c_q, n)
	mind1 := int32(len(d1pos))
	d2pos := diag2(r_q, c_q, n)
	mind2 := int32(len(d2pos))
	// Write your code here
	for i := range obstacles {
		if obstacles[i][0] == r_q {
			val := int32(0)
			if obstacles[i][0] < c_q {
				val = obstacles[i][1]
			} else {
				val = n - obstacles[i][1]
			}
			remx -= val
			continue
		}
		if obstacles[i][1] == c_q {
			val := int32(0)
			if obstacles[i][1] > r_q {
				val = obstacles[i][0]
			} else {
				val = n - obstacles[i][0]
			}
			remy -= val

			continue
		}

		for d1 := range d1pos {
			if (obstacles[i][0] == d1pos[d1][0]) && (obstacles[i][1] == d1pos[d1][1]) {
				val := abs(obstacles[i][0], d1pos[d1][0]) - 1
				if val < mind1 {
					mind1 = val
				}
			}
		}

		for d2 := range d2pos {
			if (obstacles[i][0] == d2pos[d2][0]) && (obstacles[i][1] == d2pos[d2][1]) {
				val := abs(obstacles[i][0], d2pos[d2][0]) - 1
				if val < mind2 {
					mind2 = val
				}
			}
		}
	}

	return remx + remy + mind1 + mind2
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	r_qTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != 2 {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
