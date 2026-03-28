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
    org_id BIGINT,
    created_by BIGINT,
    CONSTRAINT fk_org
      FOREIGN KEY (org_id) REFERENCES organisations(id),
    CONSTRAINT fk_created_by
      FOREIGN KEY (created_by) REFERENCES users(id)
);