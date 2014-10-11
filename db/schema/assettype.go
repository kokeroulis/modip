package DbSchema

const AssetType = `
CREATE SEQUENCE seq_assetTypeIds;

CREATE TABLE assetType
(
    id   int  primary key default nextval('seq_assetTypeIds'),
    name text unique
);
`
