CREATE TABLE client_project_access (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id BIGINT,
    project_id BIGINT,
    invited_by BIGINT,
    CONSTRAINT fk_user
      FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_project
      FOREIGN KEY (project_id) REFERENCES projects(id),
    CONSTRAINT fk_invited_by
      FOREIGN KEY (invited_by) REFERENCES users(id)
);