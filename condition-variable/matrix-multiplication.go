package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// type matrix2d struct{
// 	x int
// 	y int
// }
const (
	matrixSize = 250
)

var (
	lock    = sync.RWMutex{}
	condVar = sync.NewCond(lock.RLocker())
	wg      = sync.WaitGroup{}
	matrixA = [matrixSize][matrixSize]int{}
	matrixB = [matrixSize][matrixSize]int{}
	result  = [matrixSize][matrixSize]float64{}
)

func generateMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func main() {

	fmt.Println("Operation started...")
	t := time.Now()

	wg.Add(matrixSize)
	for row := 0; row < matrixSize; row++ {
		go matrixmul(row)
		result = [matrixSize][matrixSize]float64{}
	}

	for j := 0; j < 100; j++ {
		wg.Wait()
		lock.Lock()

		generateMatrix(&matrixA)
		generateMatrix(&matrixB)

		wg.Add(matrixSize)
		lock.Unlock()
		condVar.Broadcast()
	}

	fmt.Println("Operation ended!")
	fmt.Println("Time elapsed: ", time.Since(t))
}

func matrixmul(row int) {
	lock.RLock()
	for {
		wg.Done()
		condVar.Wait()
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				result[row][col] += float64(matrixA[row][i]) * float64(matrixB[i][col])
			}
		}
	}
}
