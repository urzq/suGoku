package sudoku

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// Grid is the basic sudoku data. The most common is a 9*9 grid.
type Grid struct {
	Size         int // Grid is squared, where width == height.
	RegionWidth  int // Each grid is subdivided in n regions.
	RegionHeight int
	Cells        []int // Values from 1 to (RegionWidth * RegionHeight). 0 means the value isn't present.
}

func (grid Grid) String() string {

	nbDigit := getNumberOfDigit(grid.RegionWidth*grid.RegionHeight - 1)

	nbRegion := grid.Size / grid.RegionWidth

	lineSize := grid.Size*nbDigit +
		(grid.Size - 1) + // number of space after each digit.
		(nbRegion-1)*2 // count all the region separator: "| "

	line := strings.Repeat("-", lineSize)

	var str strings.Builder

	for row := 0; row < grid.Size; row++ {
		if row != 0 && row%grid.RegionHeight == 0 {
			str.WriteString(line)
			str.WriteString("\n")
		}

		for col := 0; col < grid.Size; col++ {
			if col != 0 && col%grid.RegionWidth == 0 {
				str.WriteString("| ")
			}

			i := col + row*grid.Size
			num := gridNumberToString(grid.Cells[i], nbDigit)
			str.WriteString(num)

			if col < grid.Size-1 {
				str.WriteRune(' ')
			}
		}

		str.WriteRune('\n')
	}

	return str.String()
}

func gridNumberToString(num, nbDigit int) string {
	var str string
	if num == 0 {
		str = "."
	} else {
		str = strconv.Itoa(num)
	}

	if len(str) < nbDigit {
		spacePrefix := strings.Repeat(" ", nbDigit-len(str))
		return spacePrefix + str
	}
	return str
}

// getNumberOfDigit(7861) returns 4
func getNumberOfDigit(num int) int {
	nb := 0
	for num > 0 {
		num /= 10
		nb++
	}
	return nb
}

func createGridFromString(src string) Grid {
	tokens := strings.FieldsFunc(src, func(c rune) bool {
		return !unicode.IsNumber(c) && c != '|' && c != '-' && c != '.'
	})

	nbHorizontalRegionSeparator := 0
	regionWidth := 0
	numbers := make([]int, 0, len(tokens))

	for _, t := range tokens {

		if t == "|" {
			regionWidth = 0

		} else if strings.HasSuffix(t, "-") {
			nbHorizontalRegionSeparator++

		} else if t == "." {
			numbers = append(numbers, 0)

		} else {
			num, err := strconv.Atoi(t)
			if err != nil {
				panic(fmt.Sprintf("Cannot parse %d (%v)", num, err))
			}

			numbers = append(numbers, num)
			regionWidth++
		}
	}

	s := int(math.Sqrt(float64(len(numbers))))

	return Grid{
		Size:         s,
		RegionWidth:  regionWidth,
		RegionHeight: s / (nbHorizontalRegionSeparator + 1),
		Cells:        numbers,
	}
}
