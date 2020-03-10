#Dicegame
package main
import (
	"fmt"
	"log"
)
func main(){
	m := 1
	for m > 0 {
		var a,b int //input string
		if _,err := fmt.Scan(&a, &b); err != nil {
			log.Print("  Scan for i, j & k failed, due to ", err)
		}
		innerlen := a+b+1
		outcomes := make([]int, innerlen)
		for i := 0 ; i < innerlen ; i++{
			outcomes[i] = 0
		}
		for i := 1 ; i <= a ; i++  {
			for c := 1; c <= b; c++ {
				if outcomes[i+c] == 0 {
					outcomes[i+c] = 1
				}else{
					outcomes[i+c]++
				}
			}
		}
		print := make([]int,innerlen)
		max := 0
		index := 0
		for i := 0 ; i< innerlen;i++ {
			if outcomes[i] > max {
				max = outcomes[i]
				index = 0
				print[index] = i
			}else if outcomes[i] == max {
				index++
				print[index] = i
			}
		}
		for i := 0 ; i <= index; i++ {
			fmt.Println(print[i])
		}	
	}
return 
}
