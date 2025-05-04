package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i, terbesar, x, j int
	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	terbesar = A[0]
	for i = 0; i < n; i++{
		x = A[i]
		for j = i+1; j < n; j++{
			x+= A[j]
			if terbesar < x {
				terbesar = x
			}
		}
	}
	fmt.Print(terbesar)
}