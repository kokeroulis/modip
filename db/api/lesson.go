package DbApi

const LessonCreate = `
CREATE OR REPLACE FUNCTION lesson_create(lessonId int, lessonName text, departmentId int,
                                         boolean is_post_degree, text courseCodeLesson OUT alreadyExists boolean) AS $$
DECLARE
BEGIN
	PERFORM * FROM lesson
	WHERE id = lessonId AND department = departmentId;

	IF FOUND THEN
		alreadyExists := TRUE;
		RETURN;
	ELSE
		alreadyExists := FALSE;
	END IF;

	INSERT INTO lesson
	(id, name, department, isPostDegree, courseCode)
	VALUES (lessonId, lessonName, departmentId, is_post_degree, courseCodeLesson);

END;
$$ LANGUAGE plpgsql;
`