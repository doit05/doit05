package algorithms

import (
  "testing"
)

func TestQuickSort1(t *testing.T)  {
  values := []int {1, 3, 2, 5, 7, 6, 4}
  QuickSort(values)
  for i, value := range values{
    if(i != value -1){
      t.Error("QuickSort() failed. Got", values, "Expected 1 2 3 4 5 6 7")
      break
    }
  }
}

func TestQuickSort2(t *testing.T)  {
  values := []int {1, 3, 2, 6, 4, 6, 4}
  QuickSort(values)
  for i, value := range values{
    if(i > value){
      t.Error("QuickSort() failed. Got", values, "Expected 1 2 3 4 4 6 6")
      break
    }
  }
}

func TestQuickSort3(t *testing.T)  {
  values := []int {5}
  QuickSort(values)
  if(values[0] != 5){
      t.Error("QuickSort() failed. Got", values, "Expected 5")
  }
}
