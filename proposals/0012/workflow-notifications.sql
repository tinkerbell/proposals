SET ROLE tinkerbell;

CREATE TABLE IF NOT EXISTS workflow (
	id UUID UNIQUE NOT NULL
	, template UUID NOT NULL
	, devices JSONB NOT NULL
	, created_at TIMESTAMPTZ
	, updated_at TIMESTAMPTZ
	, deleted_at TIMESTAMPTZ
);

CREATE OR REPLACE FUNCTION notify_workflow_changes()
RETURNS trigger AS $$
DECLARE 
  data json;
  notification json;
BEGIN
  IF (TG_OP = 'DELETE') THEN
    data = row_to_json(OLD);
  ELSE
    data = row_to_json(NEW);
  END IF;

  notification = json_build_object(
    'table', TG_TABLE_NAME,
    'action', TG_OP,
    'data', data
  );

  PERFORM pg_notify('workflow_changed', notification::text);
  RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER workflow_changed
AFTER INSERT OR UPDATE OR DELETE
ON workflow
FOR EACH ROW
EXECUTE PROCEDURE notify_workflow_changes()
