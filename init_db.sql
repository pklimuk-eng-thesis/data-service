CREATE SCHEMA IF NOT EXISTS smart_home;

SET search_path TO smart_home;

CREATE TABLE IF NOT EXISTS presence_sensor(
	id SERIAL PRIMARY KEY,
	created_at timestamp NOT NULL DEFAULT now(),
	is_enabled BOOLEAN,
	detected BOOLEAN
);

CREATE INDEX idx_presence_sensor_created_at ON presence_sensor(created_at);
-- INSERT INTO presence_sensor(is_enabled, detected) VALUES(TRUE, FALSE);

CREATE TABLE IF NOT EXISTS gas_sensor(
	id SERIAL PRIMARY KEY,
	created_at timestamp NOT NULL DEFAULT now(),
	is_enabled BOOLEAN,
	detected BOOLEAN
);

CREATE INDEX idx_gas_sensor_created_at ON gas_sensor(created_at);

CREATE TABLE IF NOT EXISTS doors_sensor(
	id SERIAL PRIMARY KEY,
	created_at timestamp NOT NULL DEFAULT now(),
	is_enabled BOOLEAN,
	detected BOOLEAN
);

CREATE INDEX idx_doors_sensor_created_at ON doors_sensor(created_at);

CREATE TABLE IF NOT EXISTS smart_bulb(
	id SERIAL PRIMARY KEY,
	created_at timestamp NOT NULL DEFAULT now(),
	is_enabled BOOLEAN
);

CREATE INDEX idx_smart_bulb_created_at ON smart_bulb(created_at);

CREATE TABLE IF NOT EXISTS smart_plug(
	id SERIAL PRIMARY KEY,
	created_at timestamp NOT NULL DEFAULT now(),
	is_enabled BOOLEAN
);

CREATE INDEX idx_smart_plug_created_at ON smart_plug(created_at);