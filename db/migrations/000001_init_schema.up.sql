CREATE TABLE IF NOT EXISTS owners (
                                      id SERIAL PRIMARY KEY,
                                      name VARCHAR(100) NOT NULL,
                                      surname VARCHAR(100) NOT NULL,
                                      patronymic VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS cars (
                                    id SERIAL PRIMARY KEY,
                                    reg_num VARCHAR(20) UNIQUE,
                                    mark VARCHAR(100),
                                    model VARCHAR(100),
                                    year INTEGER,
                                    owner_id INTEGER REFERENCES owners(id) ON DELETE SET NULL
);