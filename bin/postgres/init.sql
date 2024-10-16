-- init.sql

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    json_data JSONB
);

INSERT INTO projects (name, description, json_data) VALUES
('Project A', 'Description for Project A', '{"key1": "value2"}'),
('Project B', 'Description for Project B', '{"key1": "value1"}'),
('Project C', 'Description for Project C', '{"key1": "value3"}');
