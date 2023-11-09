CREATE TABLE department (
    department_id SERIAL PRIMARY KEY,
    department_name VARCHAR(255) NOT NULL
);

CREATE TABLE employee (
    employee_id SERIAL PRIMARY KEY,
    employee_name VARCHAR(255) NOT NULL,
    department_id INT,
    FOREIGN KEY (department_id) REFERENCES department (department_id)
);

CREATE TABLE project (
    project_id SERIAL PRIMARY KEY,
    project_name VARCHAR(255) NOT NULL
);

CREATE TABLE employee_project (
    employee_project_id INT,
    employee_id INT,
    project_id INT,
    PRIMARY KEY(employee_id, project_id),
    FOREIGN KEY (employee_id) REFERENCES employee (employee_id),
    FOREIGN KEY (project_id) REFERENCES project (project_id)
);

CREATE TABLE task (
    task_id SERIAL PRIMARY KEY,
    task_description TEXT NOT NULL,
    project_id INT,
    FOREIGN KEY (project_id) REFERENCES project (project_id)
);

INSERT INTO department (department_name) VALUES
('Sales Department'),
('Development Department'),
('Marketing Department'),
('Finance Department'),
('Quality Department'),
('Customer Service Department'),
('Logistics Department'),
('Resource Department'),
('Security Department'),
('Research Department');

INSERT INTO employee (employee_name, department_id) VALUES
('Ivan Ivanov', 1),
('Elena Petrova', 2),
('Alexey Smirnov', 3),
('Olga Novikova', 1),
('Dmitry Kozlov', 2),
('Natalia Ivanova', 3),
('Pavel Sidorov', 1),
('Anna Kuznetsova', 2),
('Sergey Morozov', 3),
('Ekaterina Popova', 1);

INSERT INTO project (project_name) VALUES
('Project A'),
('Project B'),
('Project C'),
('Project D'),
('Project E'),
('Project F'),
('Project G'),
('Project H'),
('Project I'),
('Project J');

INSERT INTO employee_project (employee_id, project_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 6),
(7, 7),
(8, 8),
(9, 9),
(10, 10);

INSERT INTO task (task_description, project_id) VALUES
('Development of functionality', 1),
('Product testing', 2),
('Marketing campaign', 3),
('Financial analysis', 4),
('Quality control', 5),
('Technical support', 6),
('Logistics of deliveries', 7),
('Resource management', 8),
('Security provision', 9),
('Research and development', 10);
