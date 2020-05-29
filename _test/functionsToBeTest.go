package _test

import "io"

func avg(nos ...int)int{
	sum:=0
	for _,n:=range nos{
		sum+=n
	}
	if sum ==0{
		return 0
	}
	return sum/len(nos)
}

func readN(r io.Reader, n int)(string,error){
	buf:=make([]byte,n)
	m,err:=r.Read(buf)
	if err!=nil{
		return "",err
	}
	return string(buf[:m]),nil
}