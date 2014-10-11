package DbSchema

const Asset = `
CREATE SEQUENCE seq_assetIds;

CREATE TABLE asset
(
	id        int  primary key default nextval('seq_assetIds'),
	content   text unique,
	assetType int  not null references assettype(id) on delete cascade,
	teacher   int  not null references teacher(id) on delete cascade
);
`
