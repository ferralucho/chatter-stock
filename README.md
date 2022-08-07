
## Stock chat

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


## Run locally
### Producer
Follow these steps to run the producer service locally,
- At the repository root, execute the following command to open the producer directory.
  `cd .\producer\`
- Get dependencies.
  `go get .`
- Set the environment variables - use the sample .env file provided in the repository.
- Execute the `go run main.go` command to start the service.
### Consumer
Follow these steps to run the consumer service locally,
- At the repository root, execute the following command to open the consumer directory.
  `cd .\consumer\`
- Get dependencies.
  `go get .`
- Set the environment variables - use the sample .env file provided in the repository.
- Execute the `go run main.go` command to start the service.
### RabbitMQ
- There is a docker-compose.yaml file in the repository.
- Uncomment the `- "5000:5673"` to expose the RabbitMQ instance outside docker.
- [Optionally] Comment out the `producer` and `consumer` service specs so that they are not deployed in the docker.

## Running in Docker
- Running in Docker is very simple, the repo includes a `docker-compose.yaml` file.
- At the repository root, execute `docker-compose up` command to deploy the RabbitMQ instance, producer, and consumer services.
- The producer REST API will be available at `http://localhost:5050/v1/publish/example` if the default configuration is used.