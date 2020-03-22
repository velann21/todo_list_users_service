CREATE TABLE users_roles (user_id int, role_id int, CONSTRAINT FOREIGN KEY (user_id)
              REFERENCES users(id), CONSTRAINT FOREIGN KEY (role_id)
              REFERENCES roles(id), PRIMARY KEY (user_id,role_id));