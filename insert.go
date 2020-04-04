package main;
func insert(x []int,y int, i int) []int{
l:=len(x)+1;
b:=make([]int,l,l);
k:=0;
for j:=0; j<l; j++{
if j==i{
b[j]=y;
b[j+1]=x[k];
j=j+1;
}else
{
b[j]=x[k];
}
k++;
}
return b;
}

