package main

import (
        "fmt"
        "math/rand"
        "time"
        "strconv"
        "strings"
)

func main() {
	start := time.Now()

        var userinput string
        fmt.Println("Enter 'number of sides' d 'how many dice'")
        fmt.Scan(&userinput)
//Creating string variable which scans and stores user input
        var dcheck bool
        dcheck = strings.Contains(userinput, "d")
//Creating boolean variable to check for a d in userinput using strings.Contains
        if dcheck {
                var diceArray []string
                diceArray = strings.Split(userinput, "d")
//Creating an array which splits the string (userinput) when there is a d
                fmt.Println(diceArray)
                test, err := strconv.Atoi(diceArray[0])
//Converting variable called test and converting these contents to an integer where th$
                if err != nil {
                        fmt.Println("EROORORORORORORO", err)
                }

                test2, err := strconv.Atoi(diceArray[1])
//Converting a second variable called test 2 to an integer where value is equal to the$
                if err != nil {
                        fmt.Println("ERRROOORRRR NUMBER 2", err)

                }

                for i := 0; i < test; i++ {
                        fmt.Println(randNum(test2))
                }

// if i := 0 ; and i is less that test value -- add 1+
//			print random number from function below of test2 value
	}
	elapsed := time.Since(start)
	fmt.Println("Page took: ", elapsed)
}

func randNum(limitNum int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(limitNum)
}
