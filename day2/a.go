package day2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func A() {
    pwd, _ := os.Getwd()
    file, err := os.Open(pwd + "/day2/input.txt")

    if err != nil {
        fmt.Println(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    totalScore := 0
    // 2*l*w + 2*w*h + 2*h*l
    for scanner.Scan() {
        l, w, h := getDimensions(scanner.Text())
        totalScore += 2*l*w + 2*w*h + 2*h*l
        totalScore += minArea(l,w,h)
        fmt.Println(totalScore)

    }
    
    fmt.Println(totalScore)
}

func minArea(l, w, h int) int {
    list := []int{l,w,h}
    sort.Slice(list, func(i, j int) bool {
        return list[i] < list[j]
    })
    
    return list[0] * list[1]
}

func getDimensions(a string) (int, int, int) {
    spl := strings.Split(a, "x")
    x, _ := strconv.Atoi(spl[0])
    y, _ := strconv.Atoi(spl[1])
    z, _ := strconv.Atoi(spl[2])

    return x, y, z
}
