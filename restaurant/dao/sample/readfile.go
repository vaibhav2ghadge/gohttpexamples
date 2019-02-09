package main

import (  
    "bufio"
    "fmt"
    "log"
    "os"
	dbrepo "../dbrepository"
     mongoutils "../utils"

)

func main() {  
	file, err := os.Open("../restaurant.json")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
var data Restaurant
scanner := bufio.NewScanner(file)
for scanner.Scan() {
	err := json.Unmarshal(scanner.Text(), &data)
    fmt.Println(data)
	break
}

if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
}
