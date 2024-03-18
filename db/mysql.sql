-- Create the database
CREATE DATABASE gojwtdb
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

-- Create user
CREATE USER 'gojwtuser'@'localhost' IDENTIFIED BY 'gojwtpass';

-- Grant all privileges on the database to the user
GRANT ALL PRIVILEGES ON gojwtdb.* TO 'gojwtuser'@'localhost';

-- Flush privileges to apply changes
FLUSH PRIVILEGES;

-- Logout as root and log in as gojwtuser
-- mysql -u gojwtuser -p -D gojwtdb
-- Create required tables

CREATE TABLE t_user (
  id            BIGINT AUTO_INCREMENT PRIMARY KEY,
  email         VARCHAR(255) NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  firstname     VARCHAR(100),
  lastname      VARCHAR(100),
  created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
