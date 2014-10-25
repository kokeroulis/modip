package DbSchema

const Category = `
CREATE TABLE category
(
	id          int   primary key not null,
	name        text  not null,
	parent      int   default 0,
	data        text  default '',
	authactions json  not null
);
`

const CategoryGroup = `
CREATE SEQUENCE seq_categorygroupIds;

CREATE TABLE categorygroup
(
	id       int   primary key default nextval('seq_categorygroupIds') ,
	name     text  not null,
	category int   not null references category(id) on delete cascade
);
`

