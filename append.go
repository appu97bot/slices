package main;
func append(x []int,y int)[] int{
k:=len(x)+1;
var z []int;
if k<=cap(x){
z=x[:len(x)];
}else
{
l:=len(x)*2;
z=make([] int,k,l);
copy(z,x);
}
z[len(x)]=y;
return z;
}

