package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	"log"
	"ethos/efmt"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	fd, status := ethos.OpenDirectoryPath(path)
	if status != syscall.StatusOk {
		log.Fatalf ("Error opening %v: %v\n", path, status)
	}

	efmt.Println("Writing data into box type")
	data:=MyBox{10,00,40,20}
	xDiff := data.lx-data.ux
	yDiff:= data.ly-data.uy
	slope:=(yDiff/xDiff)
	manipData:=MyBox{data.lx*data.lx,data.ly*data.ly,data.ux*data.ux,data.uy*data.uy}
	efmt.Println("Original data cordinates:(",data.lx,",",data.ly,"),(",data.ux,",",data.uy,").")
	efmt.Println("Manipulated data cordinates:(",manipData.lx,",",manipData.ly,"),(",manipData.ux,",",manipData.uy,").")
	efmt.Println("Slope of the original data cordinates: ",slope)
	data.Write(fd)
	data.WriteVar(path +"box")
	manipData.Write(fd)
	manipData.WriteVar(path+"manipBox")
}
