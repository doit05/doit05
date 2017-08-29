package algorithms
import (
  "fmt"
)
// func quickSort(values []int, right, left int)  {
//   temp := values[left]
//   p := left
//   i, j := left, right
//   for i <= j{
//     for j >= p && values[j] >= temp{
//       j--
//     }
//     if j >= p{
//       values[p] = values[j]
//       p = j
//     }
//     if values[i] <= temp && i <= p{
//       i ++
//     }
//     if i <= p{
//       values[p] = values[i]
//       p = i
//     }
//   }
//   values[p] = temp
//   if p - left > 1{
//     quickSort(values, left, p - 1)
//   }
//   if right - p > 1{
//     quickSort(values, p + 1, right)
//   }
// }


func QuickSort(values []int)  {
  fmt.Println(values)
  recursionSort(values, 0, len(values) - 1)

}

func recursionSort(nums []int, left int, right int) {
    if left < right {
        pivot := partition(nums, left, right)
        recursionSort(nums, left, pivot-1)
        recursionSort(nums, pivot+1, right)
    }
}

func partition(nums []int, left int, right int) int {
    for left < right {
        for left < right && nums[left] <= nums[right] {
            right--
        }
        if left < right {
            nums[left], nums[right] = nums[right], nums[left]
            left++
        }

        for left < right && nums[left] <= nums[right] {
            left++
        }
        if left < right {
            nums[left], nums[right] = nums[right], nums[left]
            right--
        }
    }
    return left
}
