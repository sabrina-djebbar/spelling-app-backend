CREATE TABLE IF NOT EXISTS spelling_word (
    id text PRIMARY KEY,
    spelling text NOT NULL,
    definition text,
    class text NOT NULL,
    tags text,
    difficulty double precision NOT NULL,
    total_available_points int NOT NULL,
    created timestamp DEFAULT NOW()
);

ALTER TABLE spelling_word OWNER TO postgres;

CREATE TABLE IF NOT EXISTS spelling_set (
	id text PRIMARY KEY,
	name text NOT NULL,
	recommended_age text NOT NULL,
    description text,
    tags text,
    creator text NOT NULL,
    created timestamp DEFAULT NOW()
);

ALTER TABLE spelling_set OWNER TO postgres;

CREATE TABLE IF NOT EXISTS spelling_set_words (
    set_id text NOT NULL,
    word_id text NOT NULL,
    FOREIGN KEY (word_id) REFERENCES spelling_word(id),
    FOREIGN KEY (set_id) REFERENCES spelling_set(id)
);

ALTER TABLE spelling_set_words OWNER TO postgres;

CREATE TABLE IF NOT EXISTS spelling_exercise (
    id text,
    user_id text NOT NULL,
    set_id text NOT NULL,
    word_id text NOT NULL,
    spelling text NOT NULL,
    score double precision NOT NULL,
    num_of_attempts int NOT NULL,
    last_attempt date NOT NULL,
    FOREIGN KEY (word_id) REFERENCES spelling_word(id),
    FOREIGN KEY (set_id) REFERENCES spelling_set(id),
    PRIMARY KEY (word_id, id)
);

ALTER TABLE spelling_exercise OWNER TO postgres;
