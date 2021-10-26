create table users (
	user_id serial not null primary key,
	email varchar(128) not null,
	password varchar(60) not null,
	created_at timestamptz default current_timestamp
);

insert into users (email, password) values ('lumencoder@gmail.com', 'xyz');
