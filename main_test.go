package main
import ( 
"testing"
"fmt"
)


func TestisPalindrome(t *testing.T){
    word := "abcba"
    fmt.Println("Testing if the word %s is palindrome or not",word)
    if !isPalindrome(word){
	  t.Errorf("%v is a palindrome",word)
}
}
