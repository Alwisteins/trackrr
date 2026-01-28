CREATE TABLE task_assignees {
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_task
      FOREIGN KEY (task_id) REFERENCES tasks(id),
    CONSTRAINT fk_user
      FOREIGN KEY (user_id) REFERENCES users(id)
}