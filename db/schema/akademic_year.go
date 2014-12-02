package DbSchema

const AkademicYear = `
CREATE SEQUENCE seq_akademic_yearIds;

CREATE TABLE akademic_year
(
    id         int  primary key default nextval('seq_akademic_yearIds'),
    name       text not null
);
`
