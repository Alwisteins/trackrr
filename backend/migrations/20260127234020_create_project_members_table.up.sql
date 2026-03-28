CREATE TYPE project_role AS ENUM (
    'admin',
    'member'
);

CREATE TABLE project_members (
    id BIGSERIAL PRIMARY KEY,
    role project_role NOT NULL DEFAULT 'member',
    is_pic BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    project_id BIGINT,
    user_id BIGINT,
    CONSTRAINT fk_project
      FOREIGN KEY (project_id) REFERENCES projects(id),
    CONSTRAINT fk_user
      FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_project_members_project_id ON project_members(project_id);
CREATE INDEX idx_project_members_user_id ON project_members(user_id);