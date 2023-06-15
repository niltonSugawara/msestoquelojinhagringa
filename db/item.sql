CREATE TABLE item (
	id serial primary key,
	nome varchar (50) not null,
	data_criacao date not null default now(),
	valor real,
	descricao text not null
)

CREATE TABLE item (
	id uuid not null default gen_random_uuid(),
	nome varchar (50) not null,
	marca varchar (50) not null,
	data_criacao date not null default now(),
	valor real,
	descricao text not null,
	primary key(id)
)

INSERT INTO item (nome,marca,valor,descricao) VALUES ('notebook','sony',2500.00,'notebook do bom')

