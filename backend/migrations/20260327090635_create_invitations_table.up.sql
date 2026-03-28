CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'declined', 'expired');

CREATE TABLE invitations (
    id BIGSERIAL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    email VARCHAR(150) NOT NULL,
    token VARCHAR(64) NOT NULL UNIQUE,
    status invitation_status DEFAULT 'pending',
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    org_id BIGINT,
    invited_by BIGINT,
    CONSTRAINT fk_org
      FOREIGN KEY (org_id) REFERENCES organisations(id),
    CONSTRAINT fk_invited_by
      FOREIGN KEY (invited_by) REFERENCES users(id)
);