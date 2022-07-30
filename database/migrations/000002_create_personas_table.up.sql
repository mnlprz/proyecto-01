CREATE TABLE "personas" (
  "id" bigserial PRIMARY KEY,
  "num_persona" VARCHAR NOT NULL,
  "cod_entidad" VARCHAR NOT NULL,
  "cod_centro" VARCHAR NOT NULL,
  "num_contrato" VARCHAR NOT NULL,
  "cod_prod" VARCHAR NOT NULL,
  "cod_subprod" VARCHAR NOT NULL,
  "moneda" VARCHAR NOT NULL,
  "saldo" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);