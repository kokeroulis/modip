package DbSchema

const ResearchProgram = `
CREATE SEQUENCE seq_researchProgramIds;

CREATE TABLE researchProgram
(
	id            int     primary key default nextval('seq_researchProgramIds'),
    teacher       int     not null references teacher(id) on delete cascade,
	analutika_stoixeia_programmatos_titlos_programmatos   text    not null,
    analutika_stoixeia_programmatos_ylopoihsh             boolean not null,
    analutika_stoixeia_programmatos_foreas_xrimatodotisis text    not null,
    analutika_stoixeia_programmatos_proupologismos        int     not null,
    analutika_stoixeia_programmatos_diarkia               int     not null,
    analutika_stoixeia_programmatos_hmerominia            text    not null,
    analutika_stoixeia_programmatos_sxolia                text    not null
);
`
