package DbSchema

const Category = `
CREATE TABLE category
(
	id          int   primary key not null,
	name        text  not null,
	parent      int   default 0,
	data        json,
	authactions json  not null
);
`
