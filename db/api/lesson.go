package DbApi

const LessonCreate = `
CREATE OR REPLACE FUNCTION lesson_create(lessonName text, departmentId int,
                                         is_post_degree boolean, courseCodeLesson text,
										 cdCode text, OUT alreadyExists boolean) AS $$
DECLARE
	lessonId int;
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

	INSERT INTO LessonCreateReportFormEntry1
	(lesson) VALUES (lessonId);

	INSERT INTO LessonCreateReportFormEntry2
	(lesson) VALUES (lessonId);

	INSERT INTO LessonCreateReportFormEntry3
	(lesson) VALUES (lessonId);

	INSERT INTO LessonCreateReportFormEntry4
	(lesson) VALUES (lessonId);

	INSERT INTO LessonCreateReportFormEntry5
	(lesson) VALUES (lessonId);

	INSERT INTO LessonCreateReportFormEntry11
	(lesson) VALUES (lessonId);

	INSERT INTO LessonCreateReportFormEntry12
	(lesson) VALUES (lessonId);

	INSERT INTO LessonCreateReportFormEntry13
	(lesson) VALUES (lessonId);

	INSERT INTO LessonCreateReportFormEntry14
	(lesson) VALUES (lessonId);


	INSERT INTO LessonCreateReportFormEntry15
	(lesson) VALUES (lessonId);

	INSERT INTO LessonCreateReportFormEntry16
	(lesson) VALUES (lessonId);
END;
$$ LANGUAGE plpgsql;
`

