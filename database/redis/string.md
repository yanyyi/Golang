`set k1 v100` 


`get k1`

`append k1 abc`  
(integer) 7

`strlen k1`  
(integer) 7

`setnx (key) (value)`  只有在key不存在时,设置key的value  

`set k4 500`  
`incr k4`  
(integer) 501  
`decr k4`   
(integer) 500


`incrby k4 100`  
(integer) 600  
`decrby k4 200`  
(integer) 400  

`mset (key1) (value1) (key2) (value2) ....`  

`mget (key1) (key2) ....`  

`msetnx (key1) (value1) (key2) (value2) ....`

`set name lucymary`  
`getrange name 0 3`  
"lucy"  

`setrange name 3 abc`  
(integer) 8  
`get name`  
"lucabcry"


`setex (key) (过期时间) (value)`  

`getset (key) (value)`  以新换旧,设置了新值同时获得旧值  






