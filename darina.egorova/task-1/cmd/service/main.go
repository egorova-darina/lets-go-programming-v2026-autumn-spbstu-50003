
    operation, ok := readLine(scanner)
    if !ok || len(operation) != 1 {
        fmt.Println("Invalid operation")
        return
    }

    switch operation {
    case "+":
        fmt.Println(a + b)
    case "-":
        fmt.Println(a - b)
    case "*":
        fmt.Println(a * b)
    case "/":
        if b == 0 {
            fmt.Println("Division by zero")
            return
        }
        fmt.Println(a / b)
    default:
        fmt.Println("Invalid operation")
    }
}

func readLine(scanner *bufio.Scanner) (string, bool) {
    if !scanner.Scan() {
        return "", false
    }
    return strings.TrimSpace(scanner.Text()), true
}