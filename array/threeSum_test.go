package array
//Given an array nums of n integers,
// are there elements a, b, c in nums
// such that a + b + c = 0?
// Find all unique triplets in the array
// which gives the sum of zero.
// 从头遍历, 然后对剩余数组双指针寻找和为第一个元素负数的一对数.
// 注意对相同的数做跳过, 避免重复

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
  "sort"
  "testing"
)

func threeSum(nums []int) [][]int {
  var result [][]int

  sort.Ints(nums)

  for i := 0; i < len(nums) - 2; i++{
    if i > 0 && nums[i] == nums[i-1]{
      continue
    }
    target := - nums[i]

    start := i + 1
    end := len(nums) - 1
    for start < end {
      if nums[start] + nums[end] == target{
        result = append(result, []int{nums[i], nums[start], nums[end]})
        for start < end && nums[start] == nums[start+1]{
          start++
        }
        start++
        for end > start && nums[end] == nums[end-1]{
          end--
        }
        end--
      }else if nums[start] + nums[end] > target{
        end--
      }else{
        start++
      }
    }
  }
  return result
}


func TestThreeSum(t *testing.T)  {
  Convey("test three sum", t, func(){
    data := []int{-1, 0, 1, 2, -1, -4}
    dut := threeSum(data)
    fmt.Println(dut)
    So(dut[0], ShouldEqual, []int{-1, -1, 2})
    So(dut[1], ShouldEqual, []int{-1, 0, 1})
  })
}
