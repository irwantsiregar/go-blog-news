CREATE TABLE IF NOT EXISTS TABLE contents (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    excerpt VARCHAR(250) NOT NULL,
    description TEXT NOT NULL,
    image TEXT NOT NULL,
    created_by_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES categories(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    upadated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


CREATE INDEX idx_categories_created_by_id ON categories(created_by_id);
CREATE INDEX idx_contents_category_id ON contents(category_id);