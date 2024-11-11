package main
import "fmt"
func main(){
	lru:=[]int{2,3,2,0,4,0,3,1,6}
	lrucap:=4
	s:=[]int{}
	for i,val :=range lru{
		//fmt.Println(i,lrucap,s,val)
		if !check(val,s){
			if len(s)<lrucap{
				//append at end
				s=append(s,val)
			}else{
				//remove 1sst index element and append at end
				s=s[1:]
				s=append(s,val)
			}
		}else{
			//remove the val from that index and add at end
			index:=find(val,s)
			if index!=-1{
				fmt.Println(s[0:index],s[index+1:])
				s=append(s[0:index],s[index+1:]...)
				s=append(s,val)
			}
		}
		fmt.Println("after %d run %v",i,s)
	}

}

func check(val int,arr []int)bool{
	for _,x:=range arr{
		if x==val{
			return true
		}
	}
	return false
}
func find(val int,arr []int)int{
	for i,x:=range arr{
		if x==val{
			return i
		}
	}
	return -1
}