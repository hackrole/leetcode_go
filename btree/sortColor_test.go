package btree

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func sortColors(nums []int){
  dd := make(map[int]int, 3)

  for _, v := range nums {
    if _, ok := dd[v]; ok{
      dd[v]++
    }else{
      dd[v] = 1
    }
  }

  index := 0
  for i := 0; i <= 2; i++{
    cnt, ok := dd[i]
    if !ok {
      continue
    }

    for j := index; j < index + cnt; j++{
      nums[j] = i
    }
    index += cnt
  }
}

func TestSortColors(t *testing.T){
  Convey("test color", t, func(){
    data := []int{2, 0, 2, 1, 1, 0}
    sortColors(data)
    So(data[0], ShouldEqual, 0)
    So(data[3], ShouldEqual, 1)
  })
}
