package DbApi

const BookAdd = `
CREATE OR REPLACE FUNCTION book_add(bookTitle text, teacherId int, OUT id int,
                                                                   OUT title text,
                                                                   OUT alreadyExists boolean) AS $$
DECLARE
    bookRecord record;
BEGIN
	id := 0;
	title:= '';

    PERFORM * FROM book AS p WHERE bookTitle = p.title;

    IF FOUND THEN
        alreadyExists := TRUE;
        RETURN;
    ELSE
        alreadyExists := FALSE;
    END IF;

    INSERT INTO book (title, teacher) VALUES (bookTitle, teacherId)
    RETURNING * INTO bookRecord;

    id := bookRecord.id;
    title := bookRecord.title;
END;
$$ LANGUAGE plpgsql;
`
