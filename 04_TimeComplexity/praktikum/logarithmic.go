package main(){
	func logarithmic(n init) init{
		var result int = 0
		for n > 1 {
			n /= 2
			result += 1 
		}
		return result
	}
}

func main1(){
	result := logarithmic (10,5)
	fmt.Println(result)
}