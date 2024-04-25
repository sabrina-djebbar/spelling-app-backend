-- spellingWords Table
CREATE TABLE IF NOT EXISTS spelling_words (
	id text PRIMARY KEY,
    spelling text not null,
    user_id text,
    difficulty double, 
	totalAvailablePoints int
);

-- spellingSet Table
CREATE TABLE IF NOT EXISTS spelling_sets (
	id text PRIMARY KEY,
    name text,
	recommended_age text,
);

-- spellingSetWords Table
CREATE TABLE IF NOT EXISTS spelling_word_sets (
    PRIMARY KEY (user_id, session_id, set_id)
    word_id text,
    set_id text,    
	FOREIGN KEY (word_id) REFERENCES spelling_words(id),
    FOREIGN KEY (set_id) REFERENCES spelling_sets(id)

);

-- spellingExercise Table
CREATE TABLE IF NOT EXISTS spelling_exercise (
    PRIMARY KEY (user_id, session_id, set_id)
    session_id text,
    user_id text,
    word_id text,
    set_id text,
    score double,
    TotalAttempts INT,
    LastAttemptDate timestamp DEFAULT NOW(),
    FOREIGN KEY (word_id) REFERENCES spelling_words(id),
    FOREIGN KEY (set_id) REFERENCES spelling_sets(id)
);
