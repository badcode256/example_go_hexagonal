create database Business
go

use Business
go

create table Users(
id int identity,
username varchar(100),
email varchar(100),
password varchar(100),
createdAt datetime,
updatedAt datetime
)

create proc create_user
@username varchar(100),
@email varchar(100),
@password varchar(100)
as
insert into Users (username,email,password,createdAt)
values(@username,@email,@password,GETDATE())
go

alter proc edit_user
@id int,
@username varchar(100),
@email varchar(100),
@password varchar(100)
as
update Users set username=@username,email=@email,password=@password,updatedAt=GETDATE() where id=@id
go

create proc delete_user
@id int
as
delete from Users where id=@id
go

create proc list_user
as
select id,username,email,isnull(createdAt,'') as createdAt,isnull(updatedAt,'') as updatedAt from Users
go