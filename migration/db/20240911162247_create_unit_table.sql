-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "type" AS ENUM (
  'capsule',
  'cabin'
);

CREATE TYPE "status" AS ENUM (
  'Available',
  'Occupied',
  'Cleaning In Progress',
  'Maintenance Needed'
);

CREATE TABLE IF NOT EXISTS "unit" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" varchar,
  "type" type,
  "status" status,
  "lastUpdated" timestamptz
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "unit";
-- +goose StatementEnd
