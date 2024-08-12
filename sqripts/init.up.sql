CREATE TABLE IF NOT EXISTS house (
        id BIGSERIAL PRIMARY KEY,
        address VARCHAR(255) NOT NULL,
        year_built INT NOT NULL,
        developer VARCHAR(255),
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS flat (
        id SERIAL PRIMARY KEY,
        house_id INT NOT NULL,
        price INT NOT NULL,
        rooms INT NOT NULL,
        status VARCHAR(20) DEFAULT 'created',
        FOREIGN KEY (house_id) REFERENCES house (id)
);

