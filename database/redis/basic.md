`set (key) (value)`

`keys *`

`exists (key)` 

`type (key)`

`del (key)`

`unlink (key)`  #根据value选择非阻塞删除,效果相当于del   
 
`expire (key) (seconds)`  为给定的key设置过期时间  

`ttl (key)`  查看还有多少秒过期,-1表示永不过期,-2表示已过期  

`select`  命令切换数据库

`dbsize` 查看当前数据库的key的数量  

`flushdb` 清空当前库 

`flushall` 通杀全部库

