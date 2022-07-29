CREATE TABLE "ofertas" (
  "id" bigserial PRIMARY KEY,
  "num_persona" VARCHAR NOT NULL,
  "oferta" VARCHAR,
  "nombre" VARCHAR,
  "apellido" VARCHAR,
  "sit_iva" VARCHAR,
  "digital" VARCHAR,
  "fecha" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);