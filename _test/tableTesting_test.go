package _test

import "testing"

func TestAvg(t *testing.T)  {
	for _,tt:=range []struct{
		Nos []int
		Result int
  }{
  	{Nos:[]int{2,4},Result:3},
  	{Nos:[]int{1,2,5},Result:2},
  }{
  	if avg:=avg(tt.Nos...); avg!=tt.Result{
  		t.Fatalf("expected average of %v is %d, got %d\n",tt.Nos,tt.Result,avg)
	}
	}
}
