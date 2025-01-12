# Library Management System API

This is a Library Management System API written in Go, using Fiber for routing and GORM for database interaction. The API provides functionality to manage books in a library.

## Table of Contents
- [Installation]
- [Usage]
- [API Endpoints]
- [Database Setup]
    1. -- Authors (ตารางผู้แต่ง)
    CREATE TABLE Authors (
        author_id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        biography TEXT
    );

    2. -- Categories (ตารางหมวดหมู่)
    CREATE TABLE Categories (
        category_id INT AUTO_INCREMENT PRIMARY KEY,
        category_name VARCHAR(255) NOT NULL
    );

    3. -- Members (ตารางสมาชิก)
    CREATE TABLE Members (
        member_id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255),
        phone_number VARCHAR(15),
        address VARCHAR(255)
    );

    4. -- Books (ตารางหนังสือ)
    CREATE TABLE Books (
        book_id INT AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        author_id INT,
        category_id INT,
        publish_year YEAR,
        isbn VARCHAR(20) UNIQUE,
        description TEXT,
        available_qty INT DEFAULT 0,
        FOREIGN KEY (author_id) REFERENCES Authors(author_id),
        FOREIGN KEY (category_id) REFERENCES Categories(category_id)
    );

    5. -- BorrowingRecords (ตารางการยืม-คืนหนังสือ)
    CREATE TABLE BorrowingRecords (
        record_id INT AUTO_INCREMENT PRIMARY KEY,
        book_id INT,
        member_id INT,
        borrow_date DATE NOT NULL,
        return_date DATE,
        status ENUM('borrowed', 'returned') DEFAULT 'borrowed',
        FOREIGN KEY (book_id) REFERENCES Books(book_id),
        FOREIGN KEY (member_id) REFERENCES Members(member_id)
    ); 

- [Swagger UI]
    http://localhost:3000/swagger/index.html
- [License]

## Installation

To get started with this project, you need to have Go installed on your machine. You also need a running MySQL (or other SQL databases) instance.

### Prerequisites

- Go (1.18 or later)
- MySQL database (or another supported database)
- Docker (optional for running MySQL container)

### Steps to Install:

1. Clone the repository:
   ```bash
   git clone https://github.com/benjarwanhomta/library-management-system.git
   cd library-management-system
