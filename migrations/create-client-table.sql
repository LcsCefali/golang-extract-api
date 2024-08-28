CREATE TABLE IF NOT EXISTS clients (
  "id" SERIAL PRIMARY KEY,
  "name" TEXT NOT NULL,
  "credit_limit" INTEGER NOT NULL
  "credit_used" INTEGER NOT NULL DEFAULT 0 CONSTRAINT credit_used_constraint CHECK (credit_used >= (limite * -1))
);