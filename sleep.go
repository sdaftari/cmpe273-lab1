package timesleep

import ("fmt";"time")

func Sleep (t time.Duration) bool{
	if(t > 0) {
		fmt.Println("Sleep Function Called at : ", time.Now())
		
		select {
			case  <- time.After(time.Millisecond*t):
				if(t > 0) {
					fmt.Println("Timeout at : ", time.Now())
					fmt.Println("Answer is : ", 5+10)
					return true
				} else {
					return false
				}					 		
		} 
	} else {
		return false
	}	
}