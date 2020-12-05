package main
import (
	"os"
	"bufio"
	"fmt"
	"sort"
)

func Split(min int, max int, character string) (int, int) {
	if character == "F" || character == "L" {
		return min, min + (max-min)/2 
	}

	return min+(max-min+2-1)/2, max
}

func ProcessInputLine(line string) (int, int, int){
	var minRow = 0
	var maxRow = 127

	for i:=0; i<7; i++{
		minRow, maxRow = Split(minRow, maxRow, string(line[i]))
	}

	var minColumn = 0
	var maxColumn = 7
	for j:=7; j<10; j++{
		minColumn, maxColumn = Split(minColumn, maxColumn, string(line[j]))
	}

	if minRow != maxRow || minColumn != maxColumn {
		panic(fmt.Sprintf(`something went wrong,
		our binary search didn't converge %d %d %d %d`,
		minRow, maxRow, minColumn, maxColumn))
	}

	return minRow, minColumn, minRow*8+minColumn 
}

func main(){
	file, err := os.Open("day_five.txt")
	if err != nil {
		panic("Can't read the file")
	}

	scanner := bufio.NewScanner(file)

	maxId := 0
	var ids []int
	for scanner.Scan() {
		_, _, id := ProcessInputLine(scanner.Text())
		if id > maxId {
			maxId = id
		}
		ids = append(ids,id)
	}

	sort.Ints(ids)
	previousId :=-1
	for _,id := range ids {	
		if id != (previousId + 1) {
			fmt.Println("non-consequtive id:",id)	
		} else {
			fmt.Println("valid id:",id)
		}
		previousId = id
	}
	fmt.Println("maxId: ", maxId)
}
