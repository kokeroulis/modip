package DbApi

const AkademicYear = `
CREATE OR REPLACE FUNCTION akademic_year_create(akademic_year_name text,
										 out alreadyExists boolean) AS $$
DECLARE
BEGIN
	PERFORM * FROM akademic_year
	WHERE name = akademic_year_name;

	IF FOUND THEN
		alreadyExists := TRUE;
		RETURN;
	ELSE
		alreadyExists := FALSE;
	END IF;

	INSERT INTO akademic_year
	(name)
	VALUES (akademic_year_name);
END;
$$ LANGUAGE plpgsql;
`
