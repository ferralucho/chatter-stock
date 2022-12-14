
## Stock chat

#### Docker
    docker build -t chatter-stock .

    docker run -it --rm -p 5051:5050 chatter-stock

To see all running containers we can use the command inside another terminal.

    docker ps

#### Set environment variables:

mysql_users_username youruser

mysql_users_password yourpassword

mysql_users_host 127.0.0.1

mysql_users_schema users_db

		CREATE TABLE `users_db`.`users` (
		  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
		  `first_name` VARCHAR(45) NULL,
		  `last_name` VARCHAR(45) NULL,
		  `email` VARCHAR(45) NOT NULL,
		  `date_created` VARCHAR(45) NULL,
		  `status` VARCHAR(45) NULL,
		  `password` VARCHAR(45) NULL,
		  PRIMARY KEY (`id`),
		  UNIQUE INDEX `email_UNIQUE` (`email` ASC));


		SELECT * FROM users_db.users;

Postman Rest:

		CREATE USER

		curl --location --request POST 'localhost:8080/users' \
		--header 'Content-Type: application/json' \
		--data-raw '{
			"first_name": "Luciano",
			"last_name": "Ferrari",
			"email": "lucianodarioferrari@gmail.com",
			"password": "password1"
		}'

To connect to the database if you are using visual studio code, add this in launch.json 

		{
		    "name": "Launch",
		    "type": "go",
		    "request": "launch",
		    "mode": "auto",
		    "program": "${workspaceFolder}/main.go",
		    "env": {"mysql_users_username":"root","mysql_users_password":"","mysql_users_host":"127.0.0.1:3306","mysql_users_schema":"users_db"},
		    "args": []
		}
