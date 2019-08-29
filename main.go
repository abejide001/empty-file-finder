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
  var names []byte
  for _, file := range files {
	  if file.Size() == 0 { // check if file size is 0
		  name := file.Name()
		  names = append(names, name...) // name... converts the string to byte
		  names = append(names, '\n')
	  }
  }
  if err := ioutil.WriteFile("out.ext", names, 0355); err != nil {
	  fmt.Println(err)
	  return
  }
  fmt.Printf("%s", names)
}