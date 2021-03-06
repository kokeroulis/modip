package DbApi

const TeacherLogin = `
CREATE OR REPLACE FUNCTION teacher_auth(teacherEmail text, passwordCandidate text, OUT id int,
                                                                                   OUT name text,
                                                                                   OUT email text,
                                                                                   OUT departmentName text,
                                                                                   OUT departmentId int,
                                                                                   OUT auth boolean,
                                                                                   OUT username text,
                                                                                   OUT type int) AS $$
DECLARE
    teacherPassword text;
    teacherRecord record;
BEGIN
    id := 0;
    name := '';
    email := '';
    departmentName := '';
    departmentId := 0;
    auth := FALSE;

    SELECT INTO teacherPassword password FROM teacher AS t WHERE teacherEmail = t.email;

    IF NOT FOUND THEN
        RETURN;
    END IF;

    IF teacherPassword != passwordCandidate THEN
        RETURN;
    END IF;

    SELECT INTO teacherRecord t.id, t.name, t.email, department.id AS departmentId, department.name as departmentName,
    t.username, t.type
    FROM teacher AS t INNER JOIN department ON t.email = teacherEmail AND
                                               t.department = department.id;

    id := teacherRecord.id;
    name := teacherRecord.name;
    email := teacherRecord.email;
    departmentName := teacherRecord.departmentName;
    departmentId := teacherRecord.departmentId;
    username := teacherRecord.username;
    type := teacherRecord.type;
    auth := TRUE;
END;
$$ LANGUAGE plpgsql;
`

const TeacherCreate = `
CREATE OR REPLACE FUNCTION teacher_create(teacherName text,
										  teacherPassword text,
										  teacherEmail text,
										  departmentId int,
                                          teacherUsername text,
                                          teacherType int,
										  OUT alreadyExists boolean) AS $$
DECLARE
	teacherId int;
	akademic_year_id int;
BEGIN
	PERFORM * FROM teacher
	WHERE name = teacherName;

	IF FOUND THEN
		alreadyExists := TRUE;
		RETURN;
	ELSE
		alreadyExists := FALSE;
	END IF;

	INSERT INTO teacher
	(name, password, email, department, username, type)
	VALUES (teacherName, teacherPassword, teacherEmail, departmentId, teacherUsername, teacherType)
	RETURNING id INTO teacherId;

	FOR akademic_year_id IN SELECT id FROM akademic_year LOOP
		INSERT INTO TeacherCreateReportFormEntry2
		(teacher, akademic_year) VALUES (teacherId, akademic_year_id);

		INSERT INTO TeacherCreateReportFormEntry3
		(teacher, akademic_year) VALUES (teacherId, akademic_year_id);

		INSERT INTO TeacherCreateReportFormEntry4
		(teacher, akademic_year) VALUES (teacherId, akademic_year_id);

		INSERT INTO TeacherCreateReportFormEntry5
		(teacher, akademic_year) VALUES (teacherId, akademic_year_id);
	END LOOP;
END;
$$ LANGUAGE plpgsql;
`

