redis有序集合zset与普通集合set非常相似,是一个没有重复元素的字符串字符串集合。  
不同之处是有序集合的每个成员都关联了一个评分(score),这个评分(score)被用来按照从最低分到最高分的方式排序集合中的成员。集合的成员是唯一的,但是评分可以是重复的。  
因为元素是有序的，所以也可以很快的根据评分(score)或者次序(position)来获取一个范围的元素。  
访问有序集合的中间元素也是非常快的,因此能够使用有序集合作为一个没有重复成员的智能列表。  

`zadd topn 200 java 300 c++ 400 mysql 200 php`  
(integer) 4  

`zrange topn 0 -1`  
1)"java"  
2)"php"  
3)"c++"  
4)"mysql"  

`zrange topn 0 -1 withscores`  
1)"java"  
2)"200"  
3)"php"  
4)"200"
5)"c++"  
6)"300"  
7)"mysql"  
8)"400"  

`zrangebyscore topn 300 500`  
1)"c++"  
2)"mysql"  

`zincrby topn 300 php`  
"500"  

`zrem topn mysql`  
(integer) 1  

`zcount topn 200 400`  
(integer) 2  


`zrank topn java`  
(integer) 0  



