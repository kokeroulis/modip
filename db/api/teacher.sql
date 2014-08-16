--CREATE OR REPLACE FUNCTION teacher_auth(teacherEmail text, passwordCandidate text) RETURNS RECORD AS $$
--DECLARE
--    teacherPassword text;
--    teacherRecord record;
--BEGIN
--    SELECT INTO teacherPassword password FROM teacher WHERE teacherEmail = email;

--    IF NOT FOUND THEN
--        RETURN teacherRecord;
--    END IF;

--    IF teacherPassword != passwordCandidate THEN
--        RETURN teacherRecord;
--    END IF;
--
--    SELECT INTO teacherRecord id, name, email, department.id AS departmentId, department.name as departmentName
 --  - FROM teacher INNER JOIN department ON teacher.email = teacherEmail AND
      --                                    teacher.department = department.id;
   -- RETURN teacherRecord;
--END;
--$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION teacher_auth(teacherEmail text, passwordCandidate text, OUT id int,
                                                                                   OUT name text,
                                                                                   OUT email text,
                                                                                   OUT departmentName text,
                                                                                   out departmentId int) AS $$
DECLARE
    teacherPassword text;
    teacherRecord record;
BEGIN
    SELECT INTO teacherPassword password FROM teacher AS t WHERE teacherEmail = t.email;

    IF NOT FOUND THEN
        RETURN;
    END IF;

    IF teacherPassword != passwordCandidate THEN
        RETURN;
    END IF;

    SELECT INTO teacherRecord t.id, t.name, t.email, department.id AS departmentId, department.name as departmentName
    FROM teacher AS t INNER JOIN department ON t.email = teacherEmail AND
                                               t.department = department.id;

    id := teacherRecord.id;
    name := teacherRecord.name;
    email := teacherRecord.email;
    departmentName := teacherRecord.departmentName;
    departmentId := teacherRecord.departmentId;
END;
$$ LANGUAGE plpgsql;

