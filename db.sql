CREATE TABLE public.account (
	email varchar NULL,
	"password" varchar NULL
);

create table product (
	id bigserial,
	name varchar,
	amount int8,
	image varchar
);

create table info_order (
	id_order int8,
	id_product int8,
	amount int8
)