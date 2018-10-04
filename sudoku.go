package sudoku

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Cell uint8

type Grid struct {
	Size         int // Grid is squared, where width == height.
	RegionWidth  int // Each grid is subdivided in n regions.
	RegionHeight int
	Cells        []Cell
}

func (grid Grid) String() string {
	var str strings.Builder

	for row := 0; row < grid.Size; row++ {
		if row != 0 && row%3 == 0 {
			str.WriteString("---------------------\n")
		}

		for col := 0; col < grid.Size; col++ {
			if col != 0 && col%3 == 0 {
				str.WriteString("| ")
			}

			i := col + row*grid.Size
			num := strconv.Itoa(int(grid.Cells[i]))

			str.WriteString(num)
			str.WriteRune(' ')
		}

		str.WriteRune('\n')
	}

	return str.String()
}

func createGridFromString(src string) Grid {

	tokens := strings.FieldsFunc(src, func(c rune) bool {
		return !unicode.IsNumber(c) && c != '|' && c != '-' && c != '.'
	})

	nbHorizontalRegionSeparator := 0
	regionWidth := 0
	numbers := make([]Cell, 0, len(tokens))

	for _, t := range tokens {

		if t == "|" {
			regionWidth = 0

		} else if strings.HasSuffix(t, "-") {
			nbHorizontalRegionSeparator++

		} else if t == "." {
			numbers = append(numbers, Cell(0))

		} else {
			num, err := strconv.Atoi(t)
			if err != nil {
				panic(fmt.Sprintf("Cannot parse %d (%v)", num, err))
			}

			numbers = append(numbers, Cell(num))
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

func FirstMissingNumber(cells []Cell) Cell {
	presentNumbers := make([]bool, len(cells)+1)
	presentNumbers[0] = true

	for _, cell := range cells {
		if cell > 0 {
			presentNumbers[cell] = true
		}
	}

	for i, presentNumber := range presentNumbers {
		if !presentNumber {
			return Cell(i)
		}
	}

	return 0
}
