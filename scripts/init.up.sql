CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        email VARCHAR(255) unique NOT NULL,
        password VARCHAR(255) NOT NULL,
        user_type VARCHAR(20) NOT NULL
);

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
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_by UUID DEFAULT NULL,
        FOREIGN KEY (house_id) REFERENCES house (id)
);