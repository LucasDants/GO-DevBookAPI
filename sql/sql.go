package main

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {

	db, err := sql.Open("mysql", config.ConnectionDatabaseString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	exec(db, "create database if not exists devbook")
	exec(db, "use devbook")
	exec(db, "drop table if exists publications")
	exec(db, "drop table if exists followers")
	exec(db, "drop table if exists users")
	exec(db, `create table users (
		id integer auto_increment primary key,
		name varchar(50) not null,
    	nick varchar(50) not null unique,
    	email varchar(20) not null unique,
		password varchar(100) not null,
    	createdAt timestamp default current_timestamp()
	)`)

	exec(db, `CREATE TABLE followers(
    user_id int not null,
    FOREIGN KEY (user_id)
    REFERENCES  users(id)
    ON DELETE CASCADE,

    follower_id int not null,
    FOREIGN KEY (follower_id)
    REFERENCES  users(id)
    ON DELETE CASCADE,

    primary key(user_id, follower_id)
	)`)

	exec(db, `CREATE TABLE publications(
    id int auto_increment primary key,
    title varchar(50) not null,
    content varchar(300) not null,
    author_id int not null,
    FOREIGN KEY(author_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    likes int default 0,
    createdAt timestamp default current_timestamp()
	)`)

	// exec(db, `insert into users (name, nick, email, password)
	// values
	// ("Lucas", "lucas", "lucas@gmail.com", "$2a$10$oLElq5t7ZiFCAzU6.JIFUuMg5z1reoudjf9GVN.Ntyo6ZDeJn2Fna"),
	// ("Gabriel", "gabriel", "gabriel@gmail.com", "$2a$10$oLElq5t7ZiFCAzU6.JIFUuMg5z1reoudjf9GVN.Ntyo6ZDeJn2Fna");
	// insert into followers (userID, followerID)
	// values
	// (1, 2),
	// (2, 1);`)

}
