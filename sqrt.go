package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
  var z float64 = 1
  var temp float64 = 0 
  for {
	temp = z - ((z * z - x) / (2*z))
	fmt.Println(temp)
	if temp - z == 0 {
	  return temp;
	}
	z = temp;
	
  }
}

func main() {
	fmt.Println(Sqrt(2))
}
