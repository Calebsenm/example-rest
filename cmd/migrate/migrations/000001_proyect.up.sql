CREATE TABLE projects (
    project_id INT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    start_date DATE NOT NULL,
    end_date DATE,
    value DECIMAL(10, 2) NOT NULL
);