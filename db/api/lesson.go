package DbApi

const LessonCreate = `
CREATE OR REPLACE FUNCTION lesson_create(lessonName text, departmentId int,
                                         is_post_degree boolean, courseCodeLesson text,
										 cdCode text, OUT alreadyExists boolean) AS $$
DECLARE
	lessonId int;
	akademic_year_id int;
BEGIN
	PERFORM * FROM lesson
	WHERE name = lessonName AND department = departmentId;

	IF FOUND THEN
		alreadyExists := TRUE;
		RETURN;
	ELSE
		alreadyExists := FALSE;
	END IF;

	INSERT INTO lesson
	(name, department, isPostDegree, courseCode, cardisoftCode)
	VALUES (lessonName, departmentId, is_post_degree, courseCodeLesson, cdCode)
	RETURNING id INTO lessonId;

	FOR akademic_year_id IN SELECT id FROM akademic_year LOOP
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
END;
$$ LANGUAGE plpgsql;
`

