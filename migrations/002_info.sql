DO $$
    BEGIN 
        IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE
            table_schema = 'public' AND table_name = 'info') THEN

            CREATE TABLE info (
                id uuid primary key,
                "state" state
            );
        END IF;
    END;
$$