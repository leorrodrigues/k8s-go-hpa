package main

import "fmt";
import "math";

func loop() string {
	x := 0.0001
	for i:=0; i<1000000000; i++ {
		x += math.Sqrt(x)
	}
	return  "CodeEducation Rocks."
}

func main() {
	fmt.Println(loop())
}