package service

/*
   @author: vikram.choudhary 5/26/18
*/

func IsPrime(number int) bool  {

	if number <2 {
		return false
	}
	if number == 2 {
		return true
	}

	if number%2 ==0 {
		return false
	}

	for i:=3;i*i<= number ;i+=2  {
		if number%i == 0 {
			return false
		}
	}

	return true


}