use gwp;
create table posts (
   id      int auto_increment  primary key,
   content varchar(255),
   author  varchar(255)
);

insert into posts (content, author) values("Go语言","Google");
insert into posts (content, author) values("Windows操作系统","Microsoft");