package main

import (
  "fmt"
  "time"
  "os"
)

type Move struct {
  x int
  y int
  state bool
}

func InitFields(size int) [][]bool {
  fields := make([][]bool, size)

  for x := range fields {
    fields[x] = make([]bool, size)
    for y := range fields[x] {
      fields[x][y] = false
    }
  }

  return fields
}

func SetCell(fields [][]bool, x, y int) {
  fields[x][y] = !fields[x][y]
}

func UpdateCells(fields [][]bool) {
  var moves []Move

  for x, column := range fields {
    for y, value := range column {
      aliveNeighbors := 0
      if x > 0 {
        if y > 0 {
          if fields[x - 1][y - 1] == true {
            aliveNeighbors = aliveNeighbors + 1
            // fmt.Printf("(%d, %d) Neighbor found at: (%d, %d)\n",x,y,x-1,y-1)
          }
        }

        if fields[x - 1][y] == true {
          aliveNeighbors = aliveNeighbors + 1
          // fmt.Printf("(%d, %d) Neighbor found at: (%d, %d)\n",x,y,x-1,y)
        }

        if y < len(fields) - 1 {
          if fields[x - 1][y + 1] {
            aliveNeighbors = aliveNeighbors + 1
            // fmt.Printf("(%d, %d) Neighbor found at: (%d, %d)\n",x,y,x-1,y+1)
          }
        }
      }

      if x < len(fields) - 1 {
        if(y > 0) {
          if fields[x + 1][y - 1] == true {
            aliveNeighbors = aliveNeighbors + 1
            // fmt.Printf("(%d, %d) Neighbor found at: (%d, %d)\n",x,y,x+1,y-1)
          }
        }

        if fields[x + 1][y] == true {
          aliveNeighbors = aliveNeighbors + 1
          // fmt.Printf("(%d, %d) Neighbor found at: (%d, %d)\n",x,y,x+1,y)
        }

        if y < len(fields) - 1 {
          if fields[x + 1][y + 1] == true {
            aliveNeighbors = aliveNeighbors + 1
            // fmt.Printf("(%d, %d) Neighbor found at: (%d, %d)\n",x,y,x+1,y+1)
          }
        }
      }

      if y > 0 {
        if fields[x][y - 1] == true {
          aliveNeighbors = aliveNeighbors + 1
          // fmt.Printf("(%d, %d) Neighbor found at: (%d, %d)\n",x,y,x,y-1)
        }
      }

      if y < len(fields) - 1 {
        if fields[x][y + 1] == true {
          aliveNeighbors = aliveNeighbors + 1
          // fmt.Printf("(%d, %d) Neighbor found at: (%d, %d)\n",x,y,x,y+1)
        }
      }

      if aliveNeighbors < 2 && value == true {
        moves = append(moves, Move{x : x, y : y, state : false})
      } else if aliveNeighbors > 3 && value == true {
        moves = append(moves, Move{x : x, y : y, state : false})
      } else if aliveNeighbors == 3 && value == false {
        moves = append(moves, Move{x : x, y : y, state : true})
      }
    }
  }

  for i := range moves {
    fields[moves[i].x][moves[i].y] = moves[i].state
  }
}

func ShowField(fields [][]bool) {
  fmt.Printf("\n")
  for _, column := range fields {
    for _, value := range column {
      if value == true {
        fmt.Printf("X")
        } else {
          fmt.Printf("-")
        }
      }
      fmt.Printf("\n")
    }
    fmt.Printf("\n")
  }

  func main() {
    size := 0
    sleepTime := 0

    // Read input
    fmt.Printf("Enter the size of the grid: ")
    _, err := fmt.Scanf("%d", &size)
    if err != nil {
            fmt.Println(err)
    }

    fmt.Printf("Enter the wait interval (in 100 ms): ")
    _, err = fmt.Scanf("%d", &sleepTime)
    if err != nil {
        fmt.Println(err)
    }


    if(size < 1) {
      fmt.Printf("The size may not be less than 1.")
      os.Exit(1)
    }
    if(sleepTime < 1) {
      fmt.Printf("The wait interval may not be less than 1.")
      os.Exit(1)
    }

    duration := time.Duration(sleepTime * 100000000)

    fields := InitFields(size)

    // Add a glider to the field
    SetCell(fields, 1, 1)
    SetCell(fields, 2, 2)
    SetCell(fields, 2, 3)
    SetCell(fields, 3, 1)
    SetCell(fields, 3, 2)

    for {
      ShowField(fields)
      UpdateCells(fields)
      time.Sleep(duration)
    }
  }
