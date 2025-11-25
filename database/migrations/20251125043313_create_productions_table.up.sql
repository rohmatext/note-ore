CREATE TABLE IF NOT EXISTS productions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    ore_id SMALLINT NOT NULL,
    source_id INTEGER NOT NULL,
    weight DECIMAL(10, 2) NOT NULL,
    notes VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_productions_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id),

    CONSTRAINT fk_productions_ore_id
        FOREIGN KEY (ore_id)
        REFERENCES ores(id),

    CONSTRAINT fk_productions_source_id
        FOREIGN KEY (source_id)
        REFERENCES sources(id)
 )