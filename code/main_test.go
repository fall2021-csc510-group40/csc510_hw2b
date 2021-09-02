package main
import ( 
"testing"
)


func TestisPalindrome(t *testing.T){
    word := "abcba"
    if !isPalindrome(word){
	  t.Errorf("%v is a palindrome",word)
}
}
