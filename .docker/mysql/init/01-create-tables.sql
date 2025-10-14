-- MySQL Initialization Script for go-utils
-- go-utils를 위한 MySQL 초기화 스크립트

-- Use testdb database / testdb 데이터베이스 사용
USE testdb;

-- Create users table / users 테이블 생성
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    age INT NOT NULL,
    city VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_email (email),
    INDEX idx_city (city),
    INDEX idx_age (age)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert sample data / 샘플 데이터 삽입
INSERT INTO users (name, email, age, city) VALUES
    ('John Doe', 'john@example.com', 30, 'Seoul'),
    ('Jane Smith', 'jane@example.com', 25, 'Seoul'),
    ('Bob Johnson', 'bob@example.com', 35, 'Seoul'),
    ('Alice Williams', 'alice@example.com', 28, 'Incheon'),
    ('Emily Park', 'emily@example.com', 27, 'Gwangju'),
    ('Frank Lee', 'frank@example.com', 32, 'Daegu'),
    ('Grace Kim', 'grace@example.com', 29, 'Busan'),
    ('Henry Choi', 'henry@example.com', 31, 'Ulsan'),
    ('Iris Jung', 'iris@example.com', 26, 'Daejeon'),
    ('Jack Yoon', 'jack@example.com', 33, 'Gwangju')
ON DUPLICATE KEY UPDATE name=VALUES(name);

-- Verify data / 데이터 확인
SELECT COUNT(*) as total_users FROM users;
