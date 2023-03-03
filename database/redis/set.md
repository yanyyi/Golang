redis的set对外提供的功能与list是类似的,特殊之处在于set是可以自动排重的,当需要存储一个列表数据,又不希望出现重复数据时,set是一个很好的选择,并且set提供了判断某个成员是否在一个set集合内的重要接口,这也是list所不能提供的  


`sadd (key) (value1) (value2)`  
`sadd k1 v1 v2 v2`  
(integer) 2  
`smembers k1`  
1)"v2"  
2)"v1" 

`sismember k1 v2`  
(integer) 1  
`sismember k1 v3`  
(integer) 0  

`scard (key)`  返回该集合的元素个数  

`srem (key) (value1) (value2)`  删除集合中的某个元素  

`spop (key)`  随机从集合中吐出一个值并删除这个数

`srandmember (key) (n)`   随机从集合中取出一个值,但是不会删除这个数  

`smove (source) (destination) value`  

`sinter (key1) (key2)`  返回两个集合的交集元素  

`sunion (key1) (key2)`  返回两个集合的并集元素

`sdiff (key1) (key2)`  返回两个集合的差集元素(key1中包含的key2中没有的)