create database gwp;

use gwp;
create table posts(
                      id  int  primary key auto_increment,
                      content varchar(255) ,
                      author varchar(255)
);