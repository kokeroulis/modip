CREATE OR REPLACE FUNCTION paper_add(paperTitle text, teacherId int, OUT id int,
                                                                     OUT title text,
                                                                     OUT alreadyExists boolean) AS $$
DECLARE
    paperRecord record;
BEGIN
    PERFORM * FROM paper AS p WHERE paperTitle = p.title;

    IF FOUND THEN
        alreadyExists := TRUE;
        RETURN;
    ELSE
        alreadyExists := FALSE;
    END IF;

    INSERT INTO paper (title, teacher) VALUES (paperTitle, teacherId)
    RETURNING * INTO paperRecord;

    id := paperRecord.id;
    title := paperRecord.title;
END;
$$ LANGUAGE plpgsql;

