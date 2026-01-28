CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    company_name VARCHAR(255) NOT NULL UNIQUE,
    contact_name VARCHAR(255) NOT NULL,
    contact_email VARCHAR(255) NOT NULL,
    contact_phone VARCHAR(20) NOT NULL,
    notes TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);
