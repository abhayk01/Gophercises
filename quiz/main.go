package main

//This is the first version of the question .. Second one is with the name quiz1

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	correct := make(chan int)
	wrong := make(chan int)
	ch := make(chan int)
	go quizprogramme(correct, wrong, ch)

	go Counter(ch)

	for {
		val, _ := <-ch
		fmt.Println(val)
		if val == 1 {
			fmt.Println("Total Number of correct answer is", <-correct)
			fmt.Println("Total Number of Wrong answer is ", <-wrong)
			os.Exit(100)
		} else if val == 2 {
			fmt.Println("Total Number of correct answer is", <-correct)
			fmt.Println("Total Number of Wrong answer is ", <-wrong)
			fmt.Println("You have answered all of it")
		}
	}

}

func Counter(c chan int) {
	//Wait for 5 Seconds and then take the value

	time.Sleep(5 * time.Second)
	c <- 1
	//close(c)
	return
}

func quizprogramme(c chan int, w chan int, ch chan int) {

	//First we are going to read the file
	qf, err := os.Open("quiz.csv")

	if err != nil {
		log.Fatal("Not able to open the file")
		return
	}

	//Read the csv file line by line and ask for answer from the user
	csvreader := csv.NewReader(qf)
	defer qf.Close()

	var rightcounter, wrongcounter int

	for {
		record, err := csvreader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record[0])
		var input string
		fmt.Scanln(&input)

		if record[1] == input {
			rightcounter++
		} else {
			wrongcounter++
		}

	}
	c <- rightcounter
	w <- wrongcounter
	ch <- 2

	//fmt.Println("Total Number of correct answer is", rightcounter)
	//fmt.Println("Total Number of Wrong answer is ", wrongcounter)
}
