DO $$
    BEGIN 
        IF NOT EXISTS (SELECT * FROM information_schema.tables WHERE
            table_schema = 'public' AND table_name = 'history_item') THEN

        CREATE TABLE history_item (
            id uuid primary key
            "timestamp" timestamp,
            "value" varchar(101),
            info_id uuid,

            CONSTRAINT fk_info FOREIGN KEY (info_id) REFERENCES info(id)
        );
        END IF;
    END;
$$