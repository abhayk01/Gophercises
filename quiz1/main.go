package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	//Reading the command line arguements for this
	args := os.Args
	var inputTimeInt int
	inputTime := args[1:]

	if inputTime == nil {
		inputTimeInt = 2
	} else {
		inputTimeInt, _ = strconv.Atoi(inputTime[0])

	}

	//Read the input file
	file, err := os.Open("quiz.csv")
	var correct, wronganswer int

	if err != nil {
		log.Fatal("Not able to open the file")

	}
	csvread := csv.NewReader(file)

	//Call a function go through all the question and answers
	quizchan := make(chan int)
	go quiztime(csvread, &correct, &wronganswer, quizchan)

	c := make(chan int)
	//Now we will start the second goroutine to check for the timer
	go checkforTime(c, inputTimeInt)

	select {
	case res := <-c:
		fmt.Printf("Time up:- Correct answer is %v and wrong answer is %v \n", correct, wronganswer)
		correct = res
	case res := <-quizchan:
		fmt.Printf("Hurrah all answered:- Correct answer is %v and wrong answer is %v \n", correct, wronganswer)
		correct = res
	}
	go fmt.Printf("Correct answer is %d and wrong answer is %d \n", correct, wronganswer)

}

/*This function waits for time specified*/
func checkforTime(c chan int, inputTime int) {
	//inputTime = 123
	time.Sleep(time.Second * time.Duration(inputTime))
	c <- inputTime
}

func quiztime(csvread *csv.Reader, correct *int, wronganswer *int, c chan int) {
	var input string

	//fmt.Println("Coming here")
	for {
		//fmt.Println("Coming here1")
		record, err := csvread.Read()

		if err == io.EOF {
			break
		}
		fmt.Printf("Provide the sum of %v and %v \n", record[0], record[1])
		fmt.Scanln(&input)

		if input == record[2] {
			*correct++
		} else {
			*wronganswer++
		}
	}
	c <- 2
}
