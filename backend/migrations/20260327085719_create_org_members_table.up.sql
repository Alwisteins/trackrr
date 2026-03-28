CREATE TYPE org_role AS ENUM (
    'admin',
    'member',
    'owner'
);

CREATE TABLE org_members (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    org_id BIGINT,
    user_id BIGINT,
    position_id BIGINT,
    role org_role NOT NULL DEFAULT 'member',
    CONSTRAINT fk_org
      FOREIGN KEY (org_id) REFERENCES organisations(id),
    CONSTRAINT fk_user
      FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_position
      FOREIGN KEY (position_id) REFERENCES positions(id)
);