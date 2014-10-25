package DbApi

const CategoryAdd = `
CREATE OR REPLACE FUNCTION category_add(categoryId int, categoryName text, parentId int,
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
		(id, name, parent, authactions)
		VALUES (categoryId, categoryName, parentId, categoryAuthActions);
	ELSE
		INSERT INTO category
		(id, name, authactions)
		VALUES (categoryId, categoryName, categoryAuthActions);
	END IF;

END;
$$ LANGUAGE plpgsql;
`

const CategorySave = `
CREATE OR REPLACE FUNCTION category_save(categoryId int, categoryData text, OUT notExists boolean) AS $$
DECLARE
BEGIN
	PERFORM * FROM category
	WHERE id = categoryId;

	IF FOUND THEN
		notExists := FALSE;
	ELSE
		notExists := TRUE;
		RETURN;
	END IF;

	UPDATE category SET data = categoryData WHERE id = categoryId;

END;
$$ LANGUAGE plpgsql;
`

const CategoryGroupAdd = `
CREATE OR REPLACE FUNCTION category_group_add(categoryId int, groupName text, OUT alreadyExists boolean) AS $$
DECLARE
BEGIN
	PERFORM * FROM categorygroup
	WHERE category = categoryId;

	IF FOUND THEN
		alreadyExists := TRUE;
		RETURN;
	ELSE
		alreadyExists := FALSE;
	END IF;

	INSERT INTO categorygroup
	(name, category)
	VALUES (groupName, categoryId);

END;
$$ LANGUAGE plpgsql;
`

