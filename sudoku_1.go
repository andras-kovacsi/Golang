package main

import "fmt"

var sudoku = [9][9]int{
    {0, 0, 0, 2, 6, 0, 7, 0, 1},
    {6, 8, 0, 0, 7, 0, 0, 9, 0},
    {1, 9, 0, 0, 0, 4, 5, 0, 0},
    {8, 2, 0, 1, 0, 0, 0, 4, 0},
    {0, 0, 4, 6, 0, 2, 9, 0, 0},
    {0, 5, 0, 0, 0, 3, 0, 2, 8},
    {0, 0, 9, 3, 0, 0, 0, 7, 4},
    {0, 4, 0, 0, 5, 0, 0, 3, 6},
    {7, 0, 3, 0, 1, 8, 0, 0, 0},
  }

func main() {
    eValid(&sudoku, 0)
    Display(sudoku)
}

func Display(sudoku [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	for x := 0; x < 9; x++ {
		fmt.Print("| ")
		for y := 0; y < 9; y++ {
			if y == 3 || y == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", sudoku[x][y])
			if y == 8 {
				fmt.Print("|")
			}
		}
		if x == 2 || x == 5 || x == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func AbsentPeLinie(k int, sudoku [9][9]int, x int) bool {
  var y int
    for y=0; y < 9; y++ {
        if (sudoku[x][y] == k){
            return false
          }
        }
    return true
}

func AbsentPeColoana(k int, sudoku [9][9]int, y int) bool {
  var x int
  for x=0; x < 9; x++{
       if (sudoku[x][y] == k){
           return false;
         }
       }
   return true;
}

func AbsentPePatrat(k int, sudoku [9][9]int, x int, y int) bool {
  var firstX, firstY int;
  firstX =  x-(x%3)
  firstY =  y-(y%3)
  for x = firstX; x < firstX+3; x++ {
        for y = firstY; y < firstY+3; y++ {
            if (sudoku[x][y] == k){
                return false;
              }
        }
      }
    return true;

}

func eValid(sudoku *[9][9]int, position int) bool {

  if (position == 9*9){
        return true;
      }

      var x, y, k int

    x = position/9
    y = position%9

    if (sudoku[x][y] != 0){
        return eValid(sudoku, position+1);
      }

    for k=1; k <= 9; k++ {
        if AbsentPeLinie(k, *sudoku, x) && AbsentPeColoana(k, *sudoku, y) && AbsentPePatrat(k, *sudoku, x, y) {
            sudoku[x][y] = k;

            if (eValid(sudoku, position+1)){
                return true;
              }
        }
    }
    sudoku[x][y] = 0;
    return false;

}