package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type tile struct {
	isFloor    bool
	isOccupied bool
}

type floorRow []*tile

type floorPlan []*floorRow

func main() {
	var floorPlan floorPlan
	var iter int
	scanner := bufio.NewScanner(openStdinOrFile())

	for scanner.Scan() {
		floorPlan = append(floorPlan, parseRow(scanner.Text()))
	}

	for floorPlan.update() {
		iter++
		fmt.Printf("%v----------------%v\n", floorPlan.toString(), iter)
	}
	fmt.Printf("After %v iterations %v seats are occupied", iter, floorPlan.countOccupied())
}

func parseRow(rowString string) *floorRow {
	var isFloor, isOccupied bool
	row := make(floorRow, len(rowString))
	for i, s := range rowString {
		if s == 'L' {
			isFloor = false
			isOccupied = false
		} else if s == '#' {
			isFloor = false
			isOccupied = true
		} else {
			isFloor = true
			isOccupied = false
		}
		row[i] = &tile{isFloor, isOccupied}
	}

	return &row
}

func (fp floorPlan) countNeighbours(x int, y int) int {
	neighbours := fp[y].countNeighbours(x, false)
	if y-1 >= 0 {
		neighbours += fp[y-1].countNeighbours(x, true)
	}
	if y+1 < len(fp) {
		neighbours += fp[y+1].countNeighbours(x, true)
	}
	return neighbours
}

func (fp floorPlan) countVisibleNeighbours(x int, y int) int {
	var neighbours int
	var foundNW, foundN, foundNE, foundW, foundE, foundSW, foundS, foundSE bool
	for i := 1; y-i >= 0; i++ {
		currentRow := *fp[y-i]
		if foundNW && foundN && foundNE {
			break
		}
		if !foundNE {
			if x-i < 0 {
				foundNE = true
			} else if currentRow[x-i].isOccupied {
				foundNE = true
				neighbours++
			} else if !currentRow[x-i].isFloor && !currentRow[x-i].isOccupied {
				foundNE = true
			}
		}

		if !foundNW {
			if x+i >= len(currentRow) {
				foundNW = true
			} else if currentRow[x+i].isOccupied {
				foundNW = true
				neighbours++
			} else if !currentRow[x+i].isFloor && !currentRow[x+i].isOccupied {
				foundNW = true
			}
		}

		if !foundN {
			if currentRow[x].isOccupied {
				foundN = true
				neighbours++
			} else if !currentRow[x].isFloor && !currentRow[x].isOccupied {
				foundN = true
			}
		}
	}
	for i := 1; x-i >= 0 || x+i < len(*fp[y]); i++ {
		currentRow := *fp[y]
		if foundE && foundW {
			break
		}
		if x-i < 0 {
			foundE = true
		} else if !foundE {
			if currentRow[x-i].isOccupied {
				foundE = true
				neighbours++
			} else if !currentRow[x-i].isFloor {
				foundE = true
			}
		}

		if !foundW && x+i >= len(currentRow) {
			foundW = true
		} else if !foundW {
			if currentRow[x+i].isOccupied {
				foundW = true
				neighbours++
			} else if !currentRow[x+i].isFloor {
				foundW = true
			}
		}
	}
	for i := 1; y+i < len(fp); i++ {
		currentRow := *fp[y+i]
		if foundSW && foundS && foundSE {
			break
		}
		if !foundSE {
			if x-i < 0 {
				foundSE = true
			} else if currentRow[x-i].isOccupied {
				foundSE = true
				neighbours++
			} else if !currentRow[x-i].isFloor && !currentRow[x-i].isOccupied {
				foundSE = true
			}
		}

		if !foundSW {
			if x+i >= len(currentRow) {
				foundSW = true
			} else if currentRow[x+i].isOccupied {
				foundSW = true
				neighbours++
			} else if !currentRow[x+i].isFloor && !currentRow[x+i].isOccupied {
				foundSW = true
			}
		}

		if !foundS {
			if currentRow[x].isOccupied {
				foundS = true
				neighbours++
			} else if !currentRow[x].isFloor && !currentRow[x].isOccupied {
				foundS = true
			}
		}
	}
	return neighbours
}

func (fr floorRow) countNeighbours(x int, countSelf bool) int {
	var neighbours int
	if countSelf && fr[x].isOccupied {
		neighbours++
	}
	if x-1 >= 0 && fr[x-1].isOccupied {
		neighbours++
	}
	if x+1 < len(fr) && fr[x+1].isOccupied {
		neighbours++
	}
	return neighbours
}

func (fp floorPlan) update() bool {
	var totalUpdates int
	updatePlan := make([][]bool, len(fp))
	for row, fr := range fp {
		updatePlan[row] = make([]bool, len(*fr))
		for column, seat := range *fr {
			if seat.isFloor {
				continue
			}
			neighbours := fp.countVisibleNeighbours(column, row)
			if neighbours > 4 && seat.isOccupied ||
				neighbours == 0 && !seat.isOccupied {
				updatePlan[row][column] = true
				totalUpdates++
			}
		}
	}
	if totalUpdates == 0 {
		return false
	}
	for row, arr := range updatePlan {
		for column, shouldUpdate := range arr {
			if shouldUpdate {
				fr := *fp[row]
				fr[column].isOccupied = !fr[column].isOccupied
			}
		}
	}
	return true
}

func (fp floorPlan) countOccupied() int {
	var occupied int

	for _, row := range fp {
		occupied += row.countOccupied()
	}

	return occupied
}

func (fr floorRow) countOccupied() int {
	var occupied int
	for _, seat := range fr {
		if seat.isOccupied {
			occupied++
		}
	}
	return occupied
}

func (fr floorRow) toString() string {
	var sb strings.Builder
	for _, seat := range fr {
		if seat.isOccupied {
			sb.WriteRune('#')
		} else if seat.isFloor {
			sb.WriteRune('.')
		} else {
			sb.WriteRune('L')
		}
	}
	return sb.String()
}

func (fp floorPlan) toString() string {
	var sb strings.Builder
	for _, row := range fp {
		sb.WriteString(row.toString())
		sb.WriteRune('\n')
	}
	return sb.String()
}

func openStdinOrFile() io.Reader {
	var err error
	r := os.Stdin
	if len(os.Args) > 1 {
		r, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	return r
}
