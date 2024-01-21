# mytotp

# how to run golang app (http://localhost:8080)

cd golang

go mod tidy

go run main.go

# how to run react app (http://localhost:5173)

cd react

npm install

npm run dev

# how to run mysql server as docker in your notebook

docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password --name mysqltest mysql:8.3.0

# how to run mysql command line to create database and table

$ docker ps

CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES
9af429b3cacd mysql:8.3.0 "docker-entrypoint.s…" 8 minutes ago Up 8 minutes 33060/tcp, 0.0.0.0:3307->3306/tcp mysqltest

docker exec -it mysqltest sh

sh-4.4# mysql -u root -p

Enter password: password

mysql>create database myotpdb;

mysql> use myotpdb;

Database changed

mysql> create table otpuser ( id int unsigned not null PRIMARY KEY auto_increment, username varchar(255) not null default '', userpassword varchar(255) not null default '', usersecret varchar(255) not null default '',userotpurl varchar(300) not null default '');
