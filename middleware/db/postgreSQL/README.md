### Connnet to db
```shell
docker exec -it ec8334588b0b psql -U jun -d shoe
```

### Display tables
```
\dt
```

### Create tables
```
CREATE TABLE shoes(
number VARCHAR NOT NULL,
name VARCHAR,
eur NUMERIC);
```

### Insert rows
```
INSERT INTO shoes(number, name, eur) VALUES('01', 'AF1', 35);
```

### Select from tables
```
SELECT * FROM shoes;
```
