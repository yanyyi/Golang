单键多值  
redis列表是简单的字符串列表,按照插入顺序排序,可以添加一个元素到列表的头部(左边)或者尾部(右边).  
它的底层是一个双向链表,对两端的操作性能很高,通过索引下标操作中间的结点性能会较差  

`lpush/rpush (key1) (value1)(value2) ....`  

`lpop/rpop (key)`  值在键在,键光值亡  

`rpoplpush (key1)(key2)`  从key1列表右边吐出一个值,插入到key2列表左边  

`lrange (key) (start) (stop)`

`lindex (key) (index)`

`llen (key)`

`linsert (key) before (value) (newvalue)` 在value的后面插入newvalue插入值

`lrem (key) (n) (value)` 从左边删掉n个value(从左到右)  
`lpush k1 I`  
`lpush k1 hello`  
`rpush k1 am`  
`rpush k1 golang`  
`lpush k1 hello`  
`lrange k1 0 -1`  
1)"hello"  
2)"I"  
3)"am"  
4)"golang"  
5)"hello"  
`lrem k1 2 hello`  
(integer) 2  
`lrange k1 0 -1`  
1)"I"  
2)"am"  
3)"golang"

`lset (key) (index) (value)`   
`lset k1 2 Python`   
`lrange k1 0 -1`   
1)"I"  
2)"am"  
3)"Python"

