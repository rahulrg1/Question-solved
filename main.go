package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Enter number of seconds and enter")
	num := 0
	fmt.Scan(&num)
	ch := make(chan struct{})
	go func(v int) {
		time.Sleep(time.Duration(v) * time.Second)
		ch <- struct{}{}
	}(num)

	f, err := os.Open("/home/rahulgarai/Desktop/New Folder 1/problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	file := csv.NewReader(f)

	val := ""
	rightCount := 0
	incorrectCount := 0

	for {
		go func() {
			<-ch
			fmt.Println("Right Answers: ", rightCount)
			fmt.Println("Total Questions", incorrectCount+rightCount)
			os.Exit(0)
		}()
		record, err := file.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record[0])
		fmt.Scan(&val)
		if val == "-999" {
			var name string
			fmt.Scan(&name)
			os.Rename("/home/rahulgarai/Desktop/New Folder 1/problems.csv", "/home/rahulgarai/Desktop/New Folder 1/"+name+".csv")
		}
		if val == record[1] {
			rightCount++
		} else {
			incorrectCount++
		}
	}
	fmt.Println("Right Answers: ", rightCount)
	fmt.Println("Total Questions", incorrectCount+rightCount)
}
