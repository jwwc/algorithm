package main
import(
	"fmt"
	"time"
)
var fun[32][32]int
func main(){
	t:= time.Now()
	fun[1][1] = 1
	cover(0,0,1,1,32,1)
	e:= time.Since(t)
	fmt.Println(e)
	time.Sleep(time.Second/1000)
	fmt.Println(fun)
}
func cover(x int ,y int ,d int,w int,size int,flag int){
	if(size == 1){
	return 
	}	
	size = size/2
	if(d<x+size && w<y+size){
		fun[x+size-1][y+size] = flag
		fun[x+size][y+size] = flag
		fun[x+size][y+size-1]= flag
		flag = flag+1
                go  cover(x,y,w,d,size,flag)
		go  cover(x,y+size,x+size,y+size-1,size,flag)
		go  cover(x+size,y+size,x+size,y+size,size,flag)
		go  cover(x+size,y,x+size-1,y+size,size,flag)
	}
	if(d<x+size && w>=y+size){
                fun[x+size][y+size] = flag
                fun[x+size][y+size-1] = flag
                fun[x+size-1][y+size-1]= flag
		flag = flag +1
                go cover(x,y+size,w,d,size,flag)
                go cover(x+size,y+size,x+size,y+size,size,flag)
                go cover(x+size,y,x+size,y+size-1,size,flag)
                go cover(x,y,x+size-1,y+size-1,size,flag)

        }
	if(d>=x+size && w>=y+size){
                fun[x+size][y+size-1] = flag
                fun[x+size-1][y+size-1] = flag
                fun[x+size-1][y+size]=flag 
		flag = flag +1
                go cover(x+size,y+size,w,d,size,flag)
                go cover(x+size,y,x+size,y+size-1,size,flag)
                go cover(x,y,x+size-1,y+size-1,size,flag)
                go cover(x,y+size,x+size-1,y+size,size,flag)

        }
	if(d>=x+size && w<y+size){
                fun[x+size-1][y+size-1] = flag
                fun[x+size-1][y+size] = flag
                fun[x+size][y+size]= flag
		flag = flag +1
                go cover(x+size,y,w,d,size,flag)
                go cover(x,y,x+size-1,y+size-1,size,flag)
                go cover(x,y+size,x+size-1,y+size,size,flag)
                go cover(x+size,y+size,x+size,y+size,size,flag)

        }

}
