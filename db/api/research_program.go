package DbApi

const ResearchProgramCreate = `
CREATE OR REPLACE FUNCTION research_program_create(
    title                                                 text,
    analutika_stoixeia_programmatos_ylopoihsh             boolean,
    analutika_stoixeia_programmatos_foreas_xrimatodotisis text,
    analutika_stoixeia_programmatos_proupologismos        int,
    analutika_stoixeia_programmatos_diarkia               int,
    analutika_stoixeia_programmatos_hmerominia            text,
    analutika_stoixeia_programmatos_sxolia                text,
	teacherId                                             int,
    OUT alreadyExists boolean) AS $$
DECLARE
BEGIN
	PERFORM * FROM researchProgram
	WHERE analutika_stoixeia_programmatos_titlos_programmatos = title AND teacher = teacherId;

	IF FOUND THEN
		alreadyExists := TRUE;
		RETURN;
	ELSE
		alreadyExists := FALSE;
	END IF;

	INSERT INTO researchProgram
    (analutika_stoixeia_programmatos_titlos_programmatos,
    analutika_stoixeia_programmatos_ylopoihsh,
    analutika_stoixeia_programmatos_foreas_xrimatodotisis,
    analutika_stoixeia_programmatos_proupologismos,
    analutika_stoixeia_programmatos_diarkia,
    analutika_stoixeia_programmatos_hmerominia,
    analutika_stoixeia_programmatos_sxolia,
	teacher)
	VALUES
    (title,
    analutika_stoixeia_programmatos_ylopoihsh,
    analutika_stoixeia_programmatos_foreas_xrimatodotisis,
    analutika_stoixeia_programmatos_proupologismos,
    analutika_stoixeia_programmatos_diarkia,
    analutika_stoixeia_programmatos_hmerominia,
    analutika_stoixeia_programmatos_sxolia,
	teacherId);
END;
$$ LANGUAGE plpgsql;
`

