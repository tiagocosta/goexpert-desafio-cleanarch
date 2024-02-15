CREATE TABLE orders (
  id   varchar(36)  NOT NULL PRIMARY KEY,
  tax  decimal(10,2)  NOT NULL,
  price  decimal(10,2)  NOT NULL,
  final_price  decimal(10,2)  NOT NULL
);