CREATE TYPE activity_action AS ENUM (
  'status update',
  'progress update',
  'note added',
  'task created',
  'task updated',
  'comment added'
);

CREATE TABLE activity_logs (
    id BIGSERIAL PRIMARY KEY,
    action activity_action NOT NULL,
    old_value TEXT,
    new_value TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    project_id BIGINT,
    user_id BIGINT,
    CONSTRAINT fk_project
      FOREIGN KEY (project_id) REFERENCES projects(id),
    CONSTRAINT fk_user
      FOREIGN KEY (user_id) REFERENCES users(id)
);
