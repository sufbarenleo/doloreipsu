package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Open the file
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create a new bufio.Reader
    reader := bufio.NewReader(file)

    // Read and print lines
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err.Error() != "EOF" {
                fmt.Println("Error reading file:", err)
            }
            break
        }
        fmt.Print(line)
    }
}
