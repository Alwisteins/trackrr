CREATE TYPE project_status AS ENUM (
    'not started',
    'planning',
    'in progress',
    'review',
    'completed',
    'archived'
);

CREATE TABLE projects (
    id BIGSERIAL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    status project_status DEFAULT 'not started',
    progress INT DEFAULT 0,
    deadline DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    client_id BIGINT,
    pic_id BIGINT,
    CONSTRAINT fk_client
      FOREIGN KEY (client_id) REFERENCES clients(id),
    CONSTRAINT fk_user
      FOREIGN KEY (pic_id) REFERENCES users(id)
);