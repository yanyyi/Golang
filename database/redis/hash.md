
`hset user:1001 id 1`   

`hset user:1002 name Mike`  

`hget user:1001 name`  
"Mike"

`hmset user:1002 id 2 name Chen age 17 `  

`hget user:1002 age`  
"17"

`hexists user:1002 name`  
(integer)1

`hexists user:1002 gender`  
(integer)0  

`hkeys user:1002`  
1)"id"  
2)"name"  
3)"age"  

`hvals user:1002`  
1)"2"  
2)"Chen"  
3)17"

`hincrby user:1002 age 2`  
(integer)19  

`hsetnx user:1002 gender 1`  
(integer)1  


