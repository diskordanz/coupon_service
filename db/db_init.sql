CREATE DATABASE coupons_db owner postgres;

\connect coupons_db

CREATE TYPE coupon_type AS ENUM ('PERCENT', 'FLAT');
CREATE TYPE day_of_week AS ENUM ('MONDAY','TUESDAY','WEDNESDAY','THURSDAY','FRIDAY','SATURDAY','SUNDAY');

CREATE TABLE coupons(
	id		SERIAL PRIMARY KEY,
	name		TEXT NOT NULL,
	code		TEXT NOT NULL,
	description 	TEXT,
	type		coupon_type,
	status 		BOOLEAN,
	time_from 	TIME,
	time_to 	TIME,
	date_from 	DATE,
	date_to 	DATE,
	days		day_of_week[],
	value 		FLOAT,
    	franchise_id 	INTEGER,
	FOREIGN KEY (franchise_id) REFERENCES franchise(id)
);

INSERT INTO coupons(name,code,type,status,value,franchise_id) 
	VALUES
	('Coupon 1','CODE1',0,true,10,1),
	('Coupon 2','CODE2',1,false,20,1),
	('Coupon 3','CODE3',1,true,15,1);