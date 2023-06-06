package OneMillionPerformance

import (
	"bufio"
	"fmt"
	"os"
)

func SplitFiles(){
	file,err:=os.Open("/home/damitha/Desktop/circle documents/data/logbookpaynow.csv")
	if err !=nil{
		fmt.Println(err)
	}
	count:=lineCount(file)

	fmt.Println("line count:",count)
	defer file.Close()

	fileSplitter()
}

func fileSplitter(){
	r,e:=os.Open("/home/damitha/Desktop/circle documents/data/logbookpaynow.csv")
	if e !=nil{
		fmt.Println(r)
	}
	scanner:=bufio.NewScanner(r)
	f:=0
	l:=0
	fileName := "0"
	var file *os.File
	file,e=os.Create("/home/damitha/Desktop/circle documents/1mFolder/"+fileName)
	if e !=nil{
		fmt.Println(e)
		return
	}
	for scanner.Scan(){
		_,e=file.Write([]byte(scanner.Text()+"\n"))
		if e != nil{
			fmt.Println(e)
		}
		l++

		if l == 1000{
			f++
			e=file.Close()
			if e !=nil{
				fmt.Println(e)
			}
			l=0
			file,e=os.Create(fmt.Sprintf("/home/damitha/Desktop/circle documents/1mFolder/%v",f))
			if e!=nil{
				fmt.Println(e)
			}
		}
	}
}
