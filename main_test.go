package main
import ( 
"testing"
)


func TestisPalindrome(t *testing.T){
    word := "abcba"
    if IsPalindrome(word){
	  t.Errorf("%v is a palindrome",word)
}
}
