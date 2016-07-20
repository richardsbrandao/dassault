CREATE EXTENSION pgcrypto;

SELECT gen_random_uuid();

create table orders (
id serial primary key,
order_id UUID default gen_random_uuid(),
status varchar(10),
created_at timestamp,
updated_at timestamp,
price bigint,
item_id int,
transaction_id int null
)

create table items (
id serial primary key,
sku UUID default gen_random_uuid(),
unit_price bigint,
quantity smallint
)

create table transactions (
id serial primary key,
transaction_id UUID default gen_random_uuid(),
external_id varchar(40),
amount bigint,
type varchar(7),
authorization_code varchar(15),
card_brand varchar(4),
card_bin varchar(6),
card_last varchar(4)
)

ALTER TABLE orders ADD CONSTRAINT fk_items FOREIGN KEY (item_id) REFERENCES items;
ALTER TABLE orders ADD CONSTRAINT fk_transactions FOREIGN KEY (transaction_id) REFERENCES transactions;
