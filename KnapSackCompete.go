package main
import(
	"fmt"
)
func main(){
	var weight= []int{1,3,5,7,9,}
	var value= []int {2,4,6,8,10,}
	var fun [6][11] int
	for i:=1;i<=4;i++{
		for j:=1; j<=10; j++{
			fun[i][j] = fun[i-1][j]
			for k:=1;k*weight[i] <= j;k++{
				  fun[i][j] = max(fun[i][j],fun[i][j-k*weight[i]]+k*value[i])
				  fmt.Println(fun[i][j])
		}
	}
	fmt.Println(fun[4][10])
	}

}
func max(a int, b int) int{
	if a>b {
		return a
	}
	return b
}
