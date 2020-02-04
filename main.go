package main

import (
  "bufio"
  "fmt"
  "os"
)

// NEWLINE \n
const NEWLINE = 10

// YES y
const YES = 121

// NO n
const NO = 110

type FilesToGenerate struct {

}

func main() {
  err := _main(os.Args[1:])
  if err != nil {
    fmt.Printf("failed: %+v\n", err)
    return
  }
}

func _main(args []string) error {
  errChan := make(chan error)

  root := "./"
  if len(args) >= 1 {
    if args[0] != "" {
      root = fmt.Sprintf("%s/", args[0])
    }
  }
  fmt.Sprintf("root: %s", root)

  fmt.Printf("Generate Files ?\n")
  generate, err := getResponse(NEWLINE)
  if err != nil {
    errChan <- fmt.Errorf("generate question: %w", err)
  }
  if generate[0] == YES {
    fmt.Printf("")
  }

  select {
    case err := <- errChan:
      return fmt.Errorf("_main: %w", err)
    default:
      return nil
  }
}

func getResponse(d byte) ([]byte, error) {
  reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadBytes(d)
  if err != nil {
    return []byte{}, fmt.Errorf("getResponse: %w", err)
  }

  return input[0 : len(input) -1], nil
}
