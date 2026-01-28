CREATE TABLE project_notes {
    id BIGSERIAL PRIMARY KEY,
    note TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,,
    CONSTRAINT fk_project
      FOREIGN KEY (project_id) REFERENCES projects(id),
    CONSTRAINT fk_user
      FOREIGN KEY (user_id) REFERENCES users(id)
}