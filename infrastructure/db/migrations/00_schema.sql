CREATE SCHEMA IF NOT EXISTS track_api;

SET search_path TO track_api;

CREATE OR REPLACE FUNCTION track_api.update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;
