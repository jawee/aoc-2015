package day1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func B() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/day1/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()


    scanner := bufio.NewScanner(file)
    currentFloor := 0
    for scanner.Scan() {
        line := scanner.Text()

        for i, r := range strings.Split(line, "") {
            if r == "(" {
                currentFloor++
            }

            if r == ")" {
                currentFloor--
            }
            
            if currentFloor == -1 {
                fmt.Printf("%d\n",i+1)
                break;
            }
        }
    }
}
