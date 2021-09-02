package main

import "fmt"

func isPalindrome(word string) bool {
     end := len(word)-1
     for i:=0;i<len(word)/2 ;i++{
       if word[i] != word[end-i]  {
	      return false
}
}
return true
}
