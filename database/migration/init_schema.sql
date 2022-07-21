CREATE TABLE "entrada" (
  "id" bigserial PRIMARY KEY,
  "campo1" bigint NOT NULL,
  "campo2" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);
