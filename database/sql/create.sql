use  db1;

create table user(
     id bigint(20) primary key not null auto_increment,
     name varchar(20) default '',
     age int(11) default 0
)ENGINE=INNODB;

insert into user (name, age) values("高启强",40);
insert into user (name, age) values("高启盛",35);
insert into user (name, age) values("高启兰",33);
insert into user (name, age) values("安欣",37);
insert into user (name, age) values("李响",38);
insert into user (name, age) values("孟德海",63);
insert into user (name, age) values("安长林",62);
insert into user (name, age) values("孟钰",35);
insert into user (name, age) values("杨健",37);
insert into user (name, age) values("赵立冬",67);