package array

// Given an array nums of n integers and an integer target,
// are there elements a, b, c, and d in nums
// such that a + b + c + d = target?
// Find all unique quadruplets in the array which gives the sum of target.
// 两层从头到尾循环, 一次去数组不重复元素,之后用双指针取剩下元素中两个和为target-(nums[i] + nums[j])
// 注意跳过部分元素避免重复
// 注意第二层循环要大于第一层值后才能跳过, 避免num[i] == num[j], 导致跳过.

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"sort"
)

func fourSum(nums []int, target int) [][]int{
  var result [][]int

  sort.Ints(nums)

  for i := 0; i < len(nums) - 3; i++ {
    if i > 0 && nums[i] == nums[i-1]{
      continue
    }

    for j := i+1; j < len(nums) - 2; j++ {
      if j > i + 1 && nums[j] == nums[j-1]{
        continue
      }

      tgt := target -(nums[i] + nums[j])
      start := j + 1
      end := len(nums) -1
      for start < end {
        if nums[start] + nums[end] == tgt{
          result = append(result, []int{nums[i], nums[j], nums[start], nums[end]})
          for start < end && nums[start] == nums[start+1]{
            start++
          }
          start++
          for end > start && nums[end] == nums[end-1]{
            end--
          }
          end--
        }else if nums[start] + nums[end] > tgt{
          end--
        }else{
          start++
        }
      }
    }
  }
  return result
}


func TestFourSum(t *testing.T){
  Convey("test four sum", t, func(){
    data := []int{1, 0, -1, 0, -2, 2}
    dut := fourSum(data, 0)
    println(dut)
    So(dut[0], ShouldEqual, []int{-2, -1, 1, 2})
  })
}
