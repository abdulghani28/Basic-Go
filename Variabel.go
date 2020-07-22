package main

import "fmt"

func reverse(s string) string { 
    rns := []rune(s) // convert to rune 
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 { 
  
        rns[i], rns[j] = rns[j], rns[i] 
    } 
  
    return string(rns) 
} 
  
func main(){
	kakak := "merosa pesona"
	adik := reverse(kakak)


    fmt.Println("Nama Kakak : " + kakak) 
    fmt.Println("Nama Adil : " + adik) 

}
