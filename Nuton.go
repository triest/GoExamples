package main

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Sqrt(x float64) float64 {
	//var f float64;
	var z float64  =30;
	for i := 0; i < 10; i++ {
		z=z-(z*z-x)/(2*z)
	}

	return  z;
}

func fibonacci( n int)  float64 {
	var x float64 =1;
	var fib float64=0;

	 for int(x)<n{
	 	fib=fib+float64(Sqrt(5.0*fib*fib +4*(-1))/x);
	 	x=x+1;
	 }

	return  fib;
}

func main() {

//	print(Sqrt(4))
	print(fibonacci(2))

}
