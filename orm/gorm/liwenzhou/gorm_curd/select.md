# 查询  
参考博客: https://liwenzhou.com/posts/Go/gorm-crud/
## 一般查询  
根据主键查询第一条记录  
user := new(User)
db.First(&user)  
sql语句: SELECT * FROM users ORDER BY id LIMIT 1;

随机获取一条记录  
db.Take(&user)  
sql语句: SELECT * FROM users LIMIT 1;

根据主键查询最后一条记录   
db.Last(&user)  
sql语句: SELECT * FROM users ORDER BY id DESC LIMIT 1;

查询所有的记录  
var users []User   
db.Find(&users)  
sql语句: SELECT * FROM users;

查询指定的某条记录(仅当主键为整型时可用)  
db.First(&user, 10)  
sql语句: SELECT * FROM users WHERE id = 10;

## Where条件
### 普通sql查询  
Get first matched record   
db.Where("name = ?", "jinzhu").First(&user)  
sql语句: SELECT * FROM users WHERE name = 'jinzhu' limit 1;

Get all matched records   
db.Where("name = ?", "jinzhu").Find(&users)   
sql语句: SELECT * FROM users WHERE name = 'jinzhu';

<>  
db.Where("name <> ?", "jinzhu").Find(&users)   
sql语句: SELECT * FROM users WHERE name <> 'jinzhu';

IN   
db.Where("name IN (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)  
sql语句: SELECT * FROM users WHERE name in ('jinzhu','jinzhu 2');

LIKE   
db.Where("name LIKE ?", "%jin%").Find(&users)  
sql语句: SELECT * FROM users WHERE name LIKE '%jin%';

AND   
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)   
sql语句: SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

Time   
db.Where("updated_at > ?", lastWeek).Find(&users)   
sql语句: SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

BETWEEN   
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)   
sql语句: SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

### Struct & Map查询
Struct   
db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)  
sql语句: SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;

Map   
db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)  
sql语句: SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

主键的切片  
db.Where([]int64{20, 21, 22}).Find(&users)  
sql语句:  SELECT * FROM users WHERE id IN (20, 21, 22);


