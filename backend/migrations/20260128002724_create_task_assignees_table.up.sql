CREATE TABLE task_assignees (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    task_id BIGINT,
    user_id BIGINT,
    CONSTRAINT fk_task
      FOREIGN KEY (task_id) REFERENCES tasks(id),
    CONSTRAINT fk_user
      FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_task_assignees_task_id ON task_assignees(task_id);
CREATE INDEX idx_task_assignees_user_id ON task_assignees(user_id);