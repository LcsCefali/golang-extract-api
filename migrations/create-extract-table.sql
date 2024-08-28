CREATE TABLE IF NOT EXISTS extracts (
  "id" SERIAL PRIMARY KEY,
  "client_id" INTEGER NOT NULL,
  "amount" INTEGER NOT NULL,
  "operation" character(1) NOT NULL,
  "description" character(10) NOT NULL,
  "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);
