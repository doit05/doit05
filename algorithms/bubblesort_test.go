package algorithms
import (
  "testing"
)

func TestBubbleSort1(t *testing.T)  {
  values := []int {1, 3, 2, 5, 7, 6, 4}
  BubbleSort(values)
  for i, value := range values{
    if(i != value -1){
      t.Error("BulleSort() failed. Got", values, "Expected 1 2 3 4 5 6 7")
    }
  }
}

func TestBubbleSort2(t *testing.T)  {
  values := []int {1, 3, 2, 6, 4, 6, 4}
  BubbleSort(values)
  for i, value := range values{
    if(i != value -1 && i != value){
      t.Error(i,":", value, "BulleSort() failed. Got", values, "Expected 1 2 3 4 4 6 6")
      break
    }
  }
}

func TestBubbleSort3(t *testing.T)  {
  values := []int {5}
  BubbleSort(values)
  if(values[0] != 5){
      t.Error("BulleSort() failed. Got", values, "Expected 5")
  }
}
