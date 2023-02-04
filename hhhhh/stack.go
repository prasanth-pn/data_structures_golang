package main


type node struct{
	key string
	value string
	next *node
}

const size int=10
var table []*node
func main() {
	
}
func hash(key string)int{
	sum :=0
	for i:=0;i<len(key);i++{
		sum+=int(key[i])

	}
	return sum%size
}
func inser(key,value string){
	h:=hash(key)
	newnode:=new(node)
	newnode.key=key
	newnode.value=value
	newnode.next=table[h]
	table[h]=newnode

}

