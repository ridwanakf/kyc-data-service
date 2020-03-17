-- noinspection SqlNoDataSourceInspectionForFile

-- Create Table

CREATE TABLE public.kyc_mst_ktp (
	no_ktp varchar(50) NOT NULL,
	name varchar(250) NULL DEFAULT '',
	birth_place varchar(250) NULL DEFAULT '',
	birth_date date NULL,
	gender varchar(50) NULL DEFAULT '',
	address varchar(250) NULL DEFAULT '',
	religion varchar(50) NULL DEFAULT '',
	marital_status varchar(50) NULL DEFAULT '',
	job varchar(50) NULL DEFAULT '',
	nationality varchar(50) NULL DEFAULT ''
);

-- Permissions

ALTER TABLE public.kyc_mst_ktp OWNER TO postgres;
GRANT ALL ON TABLE public.kyc_mst_ktp TO postgres;