package main
import (
	"os"
	"fmt"
	"io/ioutil"
)
func main() {
  args := os.Args[1:]
    if len(args) == 0 {
	  fmt.Println("provide a directory")
	  return
  }
  files, err := ioutil.ReadDir(args[0])
  if err != nil {
	  fmt.Println("error", err)
	  return
  }
  for _, file := range files {
	  if file.Size() == 0 { // check if file size is 0
		  fmt.Println(file.Name())
	  }
  }
}