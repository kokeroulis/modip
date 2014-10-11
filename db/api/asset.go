package DbApi

const AssetAdd = `
CREATE OR REPLACE FUNCTION asset_add(assetContent text, teacherId int, assetTypeId int, OUT id int,
                                                                                        OUT content text,
                                                                                        OUT alreadyExists boolean) AS $$
DECLARE
    assetRecord record;
BEGIN
	id := 0;
	content := '';

    PERFORM * FROM asset AS a WHERE assetContent = content;

    IF FOUND THEN
        alreadyExists := TRUE;
        RETURN;
    ELSE
        alreadyExists := FALSE;
    END IF;

    INSERT INTO asset (content, teacher, assettype)
	VALUES (assetContent, teacherId, assetTypeId)
    RETURNING * INTO assetRecord;

    id := assetRecord.id;
    content := assetRecord.content;
END;
$$ LANGUAGE plpgsql;
`
const AssetRemove = `
CREATE OR REPLACE FUNCTION asset_remove(assetId int, OUT id int,
                                                     OUT content text,
													 OUT assetTypeId int,
                                                     OUT isvalid boolean) AS $$
DECLARE
    assetRecord record;
BEGIN
	id := 0;
	content := '';
	assetTypeId := 0;

    PERFORM * FROM asset AS a WHERE assetId = id;

    IF FOUND THEN
    	isvalid := TRUE;
    ELSE
        isvalid := FALSE;
		RETURN;
    END IF;

	DELETE FROM asset WHERE id = assetId
    RETURNING * INTO assetRecord;

    id := assetRecord.id;
    content := assetRecord.content;
	assetTypeId := assetRecord.assettype;
END;
$$ LANGUAGE plpgsql;
`

const AssetMove = `
CREATE OR REPLACE FUNCTION asset_move(assetId int, newAssetTypeId int, OUT id int,
                                                                       OUT content text,
												                       OUT assetTypeId int,
                                                                       OUT isvalid boolean) AS $$
DECLARE
    assetRecord record;
BEGIN
	id := 0;
	content := '';
	assetTypeId := 0;

    PERFORM * FROM asset AS a WHERE assetId = id;

    IF FOUND THEN
    	isvalid := TRUE;
    ELSE
        isvalid := FALSE;
		RETURN;
    END IF;

	UPDATE asset SET assettype = newAssetTypeId
	WHERE id = assetId
    RETURNING * INTO assetRecord;

    id := assetRecord.id;
    content := assetRecord.content;
	assetTypeId := assetRecord.assettype;
END;
$$ LANGUAGE plpgsql;
`
const AssetModify = `
CREATE OR REPLACE FUNCTION asset_remove(assetId int, newContent text, OUT id int,
                                                                      OUT content text,
												                      OUT assetTypeId int,
                                                                      OUT isvalid boolean) AS $$
DECLARE
    assetRecord record;
BEGIN
	id := 0;
	content := '';
	assetTypeId := 0;

    PERFORM * FROM asset AS a WHERE assetId = id;

    IF FOUND THEN
    	isvalid := TRUE;
    ELSE
        isvalid := FALSE;
		RETURN;
    END IF;

	UPDATE asset SET content = newContent
	WHERE id = assetId
    RETURNING * INTO assetRecord;

    id := assetRecord.id;
    content := assetRecord.content;
	assetTypeId := assetRecord.assettype;
END;
$$ LANGUAGE plpgsql;
`
