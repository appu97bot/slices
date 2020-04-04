package main;
func deleteAt(x []int,i int)[] int{
l:=len(x)-1;
b:=make([] int,l,l);
k:=0;
for j:=0; j<l; {
if j==i{
b[j]=x[k+1];
k=k+1;
}else
{
b[j]=x[k];
}
k++;
j++;
}
return b;
}

