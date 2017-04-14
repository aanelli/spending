package main
import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func main() {
    var spending_map = make(map[string]float64)

    //open the CSV file, returning the log if it errors
    file, err := os.Open("february.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    //create a scanner and start scanning line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        //type, date1, date2, place, amount
        line_data := strings.Split(line, ",")
        place := strings.Trim(line_data[3], " 0123456789:# ")
        spend:= parse_float(line_data[4])
        spending_map[place] = spending_map[place] - spend
        //fmt.Println(place, "spending is now up to", spending_map[place])

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    for k := range spending_map {
        stuff := strconv.FormatFloat(spending_map[k], 'f', 2, 64)
        fmt.Println("you spent "+stuff+"dollars at "+k)
    }

}

func parse_float(item string) float64 {
    i, err := strconv.ParseFloat(item, 64)
    if err == nil {
        return i
    }
    return -1.0
}
