package main

import "fmt"

func isPalindrome(word string) bool {
     end := len(word)-1
     for i:=0;i<len(word)/2 ;i++{
       if word[i] != word[end-i]  {
             fmt.Println("%s is not a palindrome",word)
	      return false
}
}
fmt.Println("%s is a palindrome",word)
return true
}

func main(){

fmt.Println("Palindrome Test")

}
