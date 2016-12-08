package main 

import (
	"ethos/syscall"
	//"ethos/ethos"
	"ethos/efmt"
	ethosLog "ethos/log"
)
/*
Function to check leap year.
*/
func checkLeapYear(year int) int{
	if (year%4 == 0){
  		if (year%100 == 0){
			if (year%400 == 0){
				year = 0
			}else{
				year = 1 
			}
		}else{
			year = 0
		}
	}else{
		year = 1
	}
	return year
}
/*
Function to print time by setting appropriate factors like the am or pm and time zones calculation.
*/
func checkAndPrint(timeD string,d string,zone string,month int64,year int64,day int64,h int64,m int64,s int64) {
	amPm := "am"
	if (h < 0 ){
		h = 24+h
	}else if(h == 24){
		h = 00
	}

	if (h>24){
		h = h-24
		day=day+1
	}

	if (m >= 60){
		m = m - 60
		h = h+1
	}

	if (h>=12 && h<24){
		amPm = "pm"
	}else{
		amPm = "am"
	}

	efmt.Print(zone,month,d,day,d,year," ",h,timeD,m,timeD,s," ",amPm,"\n")
	efmt.Print(zone,day,d,month,d,year," ",h,timeD,m,timeD,s," ",amPm,"\n")
}

/*
Function to format the time in various formats and zones.
*/
func printTime(timeStamp MyTime) {
	var delimiters [4]string
	delimiters[0]="/"
	delimiters[1]=":"
	delimiters[2]="-"
	d:="/"
	timeD:=delimiters[1]
	for i:=0; i<3; i+=1{
		d = delimiters[i]
		efmt.Println("Timestamp format:mm",d,"dd",d,"yy","hh:mm:s")
		checkAndPrint(timeD,d,"GMT(London):",timeStamp.month,timeStamp.year,timeStamp.day,timeStamp.hours,timeStamp.minutes,timeStamp.seconds)
		checkAndPrint(timeD,d,"Central European Time(Paris):",timeStamp.month,timeStamp.year,timeStamp.day,timeStamp.hours+1,timeStamp.minutes,timeStamp.seconds)
		checkAndPrint(timeD,d,"South African Time Zone(Johannesburg):",timeStamp.month,timeStamp.year,timeStamp.day,timeStamp.hours+2,timeStamp.minutes,timeStamp.seconds)
		checkAndPrint(timeD,d,"Indian Standard Time(New Delhi):",timeStamp.month,timeStamp.year,timeStamp.day,timeStamp.hours+5,timeStamp.minutes+30,timeStamp.seconds)
		checkAndPrint(timeD,d,"Moscow Time Zone(Moscow):",timeStamp.month,timeStamp.year,timeStamp.day,timeStamp.hours+2,timeStamp.minutes,timeStamp.seconds)
		checkAndPrint(timeD,d,"China Time Zone(Beijing):",timeStamp.month,timeStamp.year,timeStamp.day,timeStamp.hours+8,timeStamp.minutes,timeStamp.seconds)
		checkAndPrint(timeD,d,"Australian Eastern Time Zone(Sydney):",timeStamp.month,timeStamp.year,timeStamp.day,timeStamp.hours+10,timeStamp.minutes,timeStamp.seconds)
		checkAndPrint(timeD,d,"Central Time Zone(Chicago):",timeStamp.month,timeStamp.year,timeStamp.day,timeStamp.hours-6,timeStamp.minutes,timeStamp.seconds)
		checkAndPrint(timeD,d,"Eastern Time Zone(New York):",timeStamp.month,timeStamp.year,timeStamp.day,timeStamp.hours-5,timeStamp.minutes,timeStamp.seconds)
		checkAndPrint(timeD,d,"Pacific Time Zone(Seattle):",timeStamp.month,timeStamp.year,timeStamp.day,timeStamp.hours-8,timeStamp.minutes,timeStamp.seconds)
	efmt.Println("/******************************/")
	}
}

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	status := ethosLog.RedirectToLog("myProgram")
	if status != syscall.StatusOk {
        	efmt.Fprintf(syscall.Stderr,"Error opening %v: %v\n",path,status)
		syscall.Exit(syscall.StatusOk)
	}
	//Fetch epoch time (19 digits)
	efmt.Println(syscall.GetTime())
	//Parse the epoch time and convert to readable format(seconds)
	seconds := syscall.GetTime()/1000000000
	minutes := seconds/60
	hours := minutes/60
	days := hours/24
	efmt.Println("Seconds: ",seconds)
	efmt.Println("Minutes",minutes)
	efmt.Println("Hours",hours)
	efmt.Println("Days",days)

	/*Code to find the current year by parsing the seconds passed since 
	  1970 Jan 1st 0:00:00*/
	initYear := 1970
	yearFlag := 0
	noOfDays := days
	//Check for leap years and accordingly subtract 366 or 365 days from the time elapsed
	for (yearFlag == 0){
		if (checkLeapYear(initYear) == 0){
			if (noOfDays - 366 < 365){
				yearFlag = 1
			}
			noOfDays = noOfDays-366
		}else if (checkLeapYear(initYear) == 1){
			if (noOfDays - 365 < 365){
				yearFlag = 1
			}  
			noOfDays = noOfDays-365
		}
		initYear = initYear+1
	}
	efmt.Println("We are in",initYear,"with",noOfDays,"days up.")
/*Now calculate the month by running through epoch time and calculating the days passed using 31 or 30 or 28 or 29 days month"
*/
	currMonth := 1
	noOfDaysFlag := 0
	for (noOfDaysFlag != 1){
		if (currMonth == 4 || currMonth == 6 || currMonth == 9 || currMonth == 11){
		noOfDays = noOfDays - 30
		if (noOfDays<=31){
			noOfDaysFlag = 1
		}
		}else if (currMonth == 2){
			if (checkLeapYear(initYear) == 0){
				noOfDays = noOfDays - 29
			}else{
				noOfDays = noOfDays - 28
			}
			if (noOfDays<=31){
				noOfDaysFlag = 1
			}
		}else{
			noOfDays = noOfDays - 31
			if (currMonth+1 != 2){
				if (noOfDays <= 30){
					noOfDaysFlag = 1
				}
			}else{
				if(checkLeapYear(initYear) == 0){
					if (noOfDays<=29){
					noOfDaysFlag = 1
					}	
				}else{
					if(noOfDays<=28){
						noOfDaysFlag = 1
					}	
				}
			}
		}
		currMonth = currMonth + 1
	}
	efmt.Println("We are in the", currMonth,"month")
	efmt.Println("Date is: ",currMonth,"/",noOfDays+1,"/",initYear)

	/*Now calculate the amount of hours passed and the current hour along with the minutes and seconds.
	*/
	noOfHoursFlag := 0
	noOfMinutesFlag := 0
	noOfSecondsFlag := 0
	for (noOfHoursFlag != 1){
		if (hours<24){
			noOfHoursFlag = 1
		}else{
			hours = hours-24
		}
	}

	for (noOfMinutesFlag != 1){
		if (minutes<60){
			noOfMinutesFlag = 1
		}else{
			minutes = minutes - 60
		}	
	}
	
	for (noOfSecondsFlag != 1){
		if (seconds<60){
			noOfSecondsFlag = 1
		}else{
			seconds = seconds - 60
		}
	}
	
	efmt.Println("Timestamp:",currMonth,"/",noOfDays+1,"/",initYear,":",hours,":",minutes,":",seconds)
	timeStamp := MyTime{}
	timeStamp.hours = int64(hours)
	timeStamp.minutes = int64(minutes)
	timeStamp.seconds = int64(seconds)
	timeStamp.year = int64(initYear)
	timeStamp.month = int64(currMonth)
	timeStamp.day = int64(noOfDays+1)
	
	//Printing Different time zones in different time formats
	printTime(timeStamp)
}
