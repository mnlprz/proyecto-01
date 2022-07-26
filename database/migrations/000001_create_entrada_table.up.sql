CREATE TABLE "entrada" (
  "id" bigserial PRIMARY KEY,
  "campo1" VARCHAR NOT NULL,
  "campo2" VARCHAR NOT NULL,
  "campo3" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);