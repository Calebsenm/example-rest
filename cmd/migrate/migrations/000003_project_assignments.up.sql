CREATE TABLE project_assignments (
    project_id INT,
    participant_id VARCHAR(20),
    PRIMARY KEY (project_id, participant_id),
    FOREIGN KEY (project_id) REFERENCES projects(project_id),
    FOREIGN KEY (participant_id) REFERENCES participants(identification)
);
