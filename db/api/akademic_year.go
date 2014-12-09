package DbApi

const AkademicYear = `
CREATE OR REPLACE FUNCTION akademic_year_create(akademic_year_name text,
										 out alreadyExists boolean) AS $$
DECLARE
	akademic_year_id int;
	lessonId int;
	teacherId int;
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
	VALUES (akademic_year_name)
	RETURNING id INTO akademic_year_id;

	FOR lessonId IN SELECT id FROM lesson LOOP
		INSERT INTO LessonCreateReportFormEntry1
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry2
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry3
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry4
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry5
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry6
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry7
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry8
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry9
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry10
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry11
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry12
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry13
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry14
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry15
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);

		INSERT INTO LessonCreateReportFormEntry16
		(lesson, akademic_year) VALUES (lessonId, akademic_year_id);
	END LOOP;

	FOR teacherId IN SELECT id FROM teacher LOOP
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
