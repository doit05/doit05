package main
import (
  "flag"
  "fmt"
  "bufio"
  "io"
  "os"
  "strconv"
  "time"
  "github.com/doit05/algorithms"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValue(infile string)(values []int, err error)  {
  file, err := os.Open(infile)
  if err != nil{
    fmt.Println("Faild to open the input file", infile)
    return
  }

  defer file.Close()
  br := bufio.NewReader(file)
  values = make([]int, 0)
  for{
    line, isPrefix, err1 := br.ReadLine()
    if err1 != nil{
      if err1 != io.EOF {
        err = err1
      }
      break
    }
    if isPrefix{
      fmt.Println("A too long line, sees unexpected.")
      return
    }
    str := string(line)
    value, err1 := strconv.Atoi(str)
    if err1 != nil{
      err = err1
      return
    }
    values = append(values, value)
  }
  return
}
func writeValues(values []int, outfile string) error  {
  file, err := os.Create(outfile)
  if err != nil{
    fmt.Println("Faild to create the output file", outfile)
    return err
  }
  defer file.Close()
  for _, value := range values{
    str := strconv.Itoa(value)
    file.WriteString(str + "\n")
  }
  return nil
}


func main() {
  flag.Parse()
  if infile != nil{
    fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm = ", *algorithm)
  }
  values, err := readValue(*infile)
  if err == nil {
    t1 := time.Now()
    switch *algorithm {
    case "qsort":
      algorithms.QuickSort(values)
    case "bubblesort":
      algorithms.BubbleSort(values)
    default:
      fmt.Println("Sorting algorithm", *algorithm, "is not found!")
    }
    t2 := time.Now()
    fmt.Println("Sort time: ", t2.Sub(t1))
    writeValues(values, *outfile)
  }else{
    fmt.Println(err)
  }
}
