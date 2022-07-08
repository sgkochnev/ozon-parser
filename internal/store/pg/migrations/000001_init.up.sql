CREATE Table goods(
    id serial PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL ,
    url_img TEXT,
    price INTEGER
);
