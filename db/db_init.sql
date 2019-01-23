CREATE DATABASE coupons3_db owner postgres;

\connect coupons3_db

--CREATE TYPE coupon_type AS ENUM ('PERCENT', 'FLAT');
--DROP TABLE coupons FROM coupons3_db;

CREATE TABLE coupons(
	id				SERIAL PRIMARY KEY,
	name 			TEXT NOT NULL,
	code			TEXT NOT NULL,
	description 	TEXT,
	type			INT,
	status 			BOOLEAN,
	time_from 		TIME,
	time_to 		TIME,
	date_from 		DATE,
	date_to 		DATE,
	value 			FLOAT,
    franchise_id 	INT NOT NULL,
	days			TEXT);
	--FOREIGN KEY (franchise_id) REFERENCES franchise(id)
--);

INSERT INTO coupons(name,code,type,status,value,franchise_id) VALUES ('Coupon 1','CODE1',0,true,10,1),
	('Coupon 2','CODE2',1,false,20,1),
	('Coupon 3','CODE3',1,true,15,1);