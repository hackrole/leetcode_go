package btree

import (
	. "github.com/smartystreets/goconvey/convey"
  "testing"
)

func searchMatrix(matrix [][]int, target int) bool {
    sr, sc := 0, 0
    er, ec := len(matrix), len(matrix[0])
    mr, mc := er / 2, ec /2


    for  {
        if matrix[mr][mc] == target{
            return true
        }

        if matrix[mr][mc] < target{
            if er == mr && ec == mc {
                break
            }
            sr = mr
            sc = mc

            mr = (er + mr) / 2
            mc = (ec + mc) / 2
        }else{
            if sr == mr && sc == mc {
                break
            }
            er = mr
            ec = mc

            mr = (mr + sr) / 2
            mc = (mc + sc) / 2
        }
    }
    return false
}

func TestSearchMatrix(t *testing.T)  {

  Convey("test search", t, func(){
    data := [][]int{
      []int{1, 3, 5, 7},
      []int{10, 11, 16, 20},
      []int{23, 30, 34, 50},
    }
    dut := searchMatrix(data, 3)
    So(dut, ShouldEqual, true)
  })
}
