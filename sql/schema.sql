CREATE TABLE categories (
    id   INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    parent_category_id INTEGER,
    FOREIGN KEY(parent_category_id) REFERENCES categories(id)
);

CREATE TABLE products (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    category_id INT,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
