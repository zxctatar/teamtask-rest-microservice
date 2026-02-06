CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    description VARCHAR(255) NOT NULL,
    deadline TIMESTAMPTZ
);

CREATE INDEX idx_project_id ON tasks(project_id);