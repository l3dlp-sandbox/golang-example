****************
    ARRAY
****************
Arrays are fixed length sequence of items
Arrays in GO are values not like c/c++(where arrays act like pointers) unlike java(where arrays are object reference)
so
1) Assigning one array to another copies all of the elements 
2) if pass array to function it will receive copy of all array

ex :=[N]Type
     [N]string{"dd","aa"}
     [...]int{1,2,3}
****************
     SLICE
****************
slices are reference types. can be passed to function without having a new copy
slice size can be extended using inbuilt function append
ex
make([]type.length,capacity)
make([]type,length)
[]type{}
[]type{value1,value2}
     
 