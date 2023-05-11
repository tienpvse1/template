-- migrate:up
ALTER TABLE public."user" ADD COLUMN "address" TEXT

-- migrate:down
ALTER TABLE public."user" DROP COLUMN "address"

