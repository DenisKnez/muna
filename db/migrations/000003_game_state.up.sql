DO $$
    BEGIN
        IF NOT EXISTS (SELECT * FROM pg_catalog.pg_type typ INNER JOIN pg_catalog.pg_namespace nsp ON nsp."oid" = typ.typnamespace
            WHERE nsp.nspname  = current_schema() AND typ.typname = 'state') THEN
            CREATE TYPE "state" AS ENUM('1', '2');
            END IF;
    END ;
$$
