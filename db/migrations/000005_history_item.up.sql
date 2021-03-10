CREATE TABLE IF NOT EXISTS history_item (
    id uuid primary key,
    "timestamp" timestamp,
    "value" varchar(101),
    info_id uuid,

    CONSTRAINT fk_info FOREIGN KEY (info_id) REFERENCES info(id)

)