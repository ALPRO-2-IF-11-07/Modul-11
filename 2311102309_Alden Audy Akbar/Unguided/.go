package main
//2311102309
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Masukkan data suara :")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    data := strings.Split(input, " ")

    var totalSuara int
    var suaraValid []int
    hitungSuara := make(map[int]int)

    for _, item := range data {
        num, err := strconv.Atoi(item)
        if err != nil {
            continue
        }

        if num == 0 {
            break 
        }

        totalSuara++
if num >= 1 && num <= 20 {
            suaraValid = append(suaraValid, num)
            hitungSuara[num]++
        }
    }

    fmt.Printf("Suara masuk: %d\n", totalSuara)
    fmt.Printf("Suara sah: %d\n", len(suaraValid))

    for calon, jumlah := range hitungSuara {
        fmt.Printf("%d: %d\n", calon, jumlah)
    }
}


