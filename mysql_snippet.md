show databases;
use test;
create table user(id int, name varchar(255));
show tables;
show columns from user;

#CRUD
insert into user (id, name) values (1,'yale');
select * from user;
update user set id=2, name='node' where id=1;
delete from user where id=2;



