package DbApi

const CategoryAdd = `
CREATE OR REPLACE FUNCTION category_add(categoryId int, categoryName text, parentId int, categoryData json,
                                        categoryAuthActions json, OUT alreadyExists boolean) AS $$
DECLARE
BEGIN
	PERFORM * FROM category
	WHERE id = categoryId;

	IF FOUND THEN
		alreadyExists := TRUE;
		RETURN;
	ELSE
		alreadyExists := FALSE;
	END IF;

	IF parentId > 0 THEN
		INSERT INTO category
		(id, name, parent, data, authactions)
		VALUES (categoryId, categoryName, parentId, categoryData, categoryAuthActions);
	ELSE
		INSERT INTO category
		(id, name, data, authactions)
		VALUES (categoryId, categoryName, categoryData, categoryAuthActions);
	END IF;

END;
$$ LANGUAGE plpgsql;
`
