package _test

import (
	"errors"
	"testing"
)

type ReadMock struct {
	ReadMock func([]byte)(int,error)
}

func (m ReadMock)Read(p []byte)(int,error){
	return m.ReadMock(p)
}

func TestReadN_bufSize(t *testing.T)  {
	total:=0
	mr:=&ReadMock{
		func(b []byte) (i int, e error) {
			total=len(b)
			return 0,nil
		},
	}
	readN(mr,5)
	if total !=5{
		t.Fatalf("expected 5, got %d",total)
	}
}

func TestReadN_error(t *testing.T)  {
	expect:=errors.New("some non-nill error")
	mr:=&ReadMock{
		func(bytes []byte) (i int, e error) {
			return 0,expect
		},
	}
	_,err:=readN(mr,5)
	if err !=expect{
		t.Fatal("expected error")
	}
}