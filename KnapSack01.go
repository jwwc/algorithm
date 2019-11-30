//01背包是一种动态规划的思想
//动态规划思想一般用于最优解和计数当中的最值问题
//动态规划的核心就是找出状态转移方程
//状态转移方程的思想就是把当前问题转化为和当前问题一样但是数据的规模更小的问题
package main
import(
	"fmt"
	"math"
)
func main(){
	var value = []int{2,4,9,5,7,}
	var weight = []int{1,3,7,6,8,}
	var volume  = 12
	var fun [20][20]int 
	//var j = volume
	var i int
	for i=1;i<5;i++{
		for j:=volume; j>=1; j--{
		if(j>=weight[i]){
			fun[i][j] = int(math.Max(float64(fun[i-1][j]),float64(fun[i-1][j-weight[i]]+value[i])))
		//	fmt.Println(fun[i][j],i,j)
		}else{
			fun[i][j] = fun[i-1][j]
		}
		//fmt.Println(fun[i][j])
	}
	//fmt.Println(fun[i-1][0])
	}
	fmt.Println(fun[i-1][volume])
}
/*func main(){
	var value = []int{2,4,5,7,9,}
        var weight = []int{1,3,5,8,7,}
        var volume  = 12
        var fun [20][20]int
        var j = volume
}
func dfs(value [5]int,weight [5]int,i int,j int)int{
	if(j<value[i] | i<= 0){
		return 
	}
	return int(math.Max(float64(dfs(value,weight,i-1,j)),float64(dfs(value,weight,i-1,j-weight[i])+value[i])))
}*/
