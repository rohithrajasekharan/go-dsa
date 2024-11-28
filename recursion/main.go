package main

import "fmt"

func factorial(n int) int{
  if n<=1 {
    return n
  }else{
    return n*factorial(n-1)
  }
}

func main(){
  fmt.Printf("factorial of 5 is %v",factorial(5))
}
