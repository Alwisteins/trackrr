CREATE TABLE project_notes (
    id BIGSERIAL PRIMARY KEY,
    note TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    project_id BIGINT,
    user_id BIGINT,
    CONSTRAINT fk_project
      FOREIGN KEY (project_id) REFERENCES projects(id),
    CONSTRAINT fk_user
      FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_project_notes_project_id ON project_notes(project_id);
CREATE INDEX idx_project_notes_user_id ON project_notes(user_id);