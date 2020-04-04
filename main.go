package main;
import "log"
func main(){
a:=[] int{1,2,3,10,20,12,45,4,5};
log.Println(a);
a=append(a,100);
log.Println("Added an element 100 to [1,2,3,10,20,12,45,4,5]",  a);
a=insert(a,11,5);
log.Println("Inserted ea element 11 at index 5 to [1,2,3,10,20,12,45,4,5,100]",  a);
a=delete(a);
log.Println("delete an element to [1,2,3,10,20,11,12,45,4,5,100]",  a);
a=deleteAt(a,3);
log.Println("delete an element at index 3 to [1,2,3,10,20,11,12,45,4,5]",  a);
}

