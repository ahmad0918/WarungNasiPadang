## Project Warung Nasi Padang

### Architecture

1. PostgreSQL
2. Gin Framework

### Database

```sql
CREATE DATABASE db_warung_nasi_pandang;

CREATE TABLE menu (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(100),
  price INTEGER,
  CONSTRAINT name_unique UNIQUE(name)
);

CREATE TABLE transaksi (
    id VARCHAR(100) PRIMARY KEY,
    menu VARCHAR(100),
    quantity INTEGER,
    date DATE,
    CONSTRAINT fk_transaksi_menu FOREIGN KEY(menu) 
    REFERENCES menu(name)
);
```

### Make sure this project able to run

### Windows

```
SET DB_HOST=localhost
SET DB_PORT=5432
SET DB_NAME=db_warung_nasi_padang
SET DB_USER=yourdbuser
SET DB_PASSWORD=yourdbpassword
SET API_PORT=8012
SET API_HOST=localhost
go run .
```

### Linux / Mac

```
DB_HOST=localhost DB_PORT=5432 DB_NAME=db_warung_nasi_padang DB_USER=yourdbuser DB_PASSWORD=yourdbpassword API_PORT=8012 API_HOST=localhost go run .
```

### what will be tested

1. Repository
2. Usecase