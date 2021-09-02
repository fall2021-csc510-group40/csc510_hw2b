package main
import ( 
"testing"
)

func IsPalindrome(word string) bool {
     end := len(word)-1
     for i:=0;i<len(word)/2 ;i++{
       if word[i] != word[end-i]  {
              return false
}
}
return true
}

func TestIsPalindrome(t *testing.T){
    word := "abcba"
    if IsPalindrome(word){
	  t.Errorf("%v is a palindrome",word)
}
}
