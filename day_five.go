package main
import (
	"os"
	"bufio"
	"fmt"
)

func Split(min int, max int, character string) (int, int) {
	if character == "F" || character == "L" {
		return min, min + (max-min)/2 
	}

	return min+(max-min+2-1)/2, max
}

func main(){
	file, err := os.Open("day_five.txt")
	if err != nil {
		panic("Can't read the file")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
