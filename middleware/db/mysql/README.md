# Learning Relational Database in Go using MySQL

Using standard library via [database/sql](https://pkg.go.dev/database/sql) with the driver [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

# Login
```
docker exec -it cf4185400ab6 mysql -u root -p -h localhost
```

# Select database
```
USE shoes
```

# Create schema
```sql
CREATE TABLE shoes(
    number VARCHAR(10) NOT NULL,
    name VARCHAR(10)
);
```

# Insert
```sql
INSERT INTO shoes(number, name) VALUES("01", "AF1");
```
# Enable Bluk Insert
```sql
SHOW GLOBAL VARIABLES LIKE 'local_infile';
SET GLOBAL local_infile = 1;
```