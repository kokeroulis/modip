package DbSchema

const Department = `
CREATE SEQUENCE seq_departmentIds;

CREATE TABLE department
(
    id   int  primary key default nextval('seq_departmentIds'),
    name text not null
);
`
