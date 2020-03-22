CREATE table users (id int NOT NULL  AUTO_INCREMENT, first_name varchar(255) NOT NULL , last_name varchar(255)
                     ,email varchar(255), phone_number varchar(255), dob date NOT NULL,
                     password varchar(255), CONSTRAINT PRIMARY KEY (id));
