CREATE TABLE IF NOT EXISTS tasks (
                                     id SERIAL PRIMARY KEY,
                                     description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS gifts (
                                     id SERIAL PRIMARY KEY,
                                     name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS user_tasks (
                                          user_id BIGINT NOT NULL,
                                          task_id INT NOT NULL,
                                          completed BOOLEAN DEFAULT FALSE,
                                          PRIMARY KEY (user_id, task_id),
                                          FOREIGN KEY (task_id) REFERENCES tasks(id)
);

CREATE TABLE IF NOT EXISTS user_gifts (
                                          user_id BIGINT NOT NULL,
                                          gift_id INT NOT NULL,
                                          PRIMARY KEY (user_id, gift_id),
                                          FOREIGN KEY (gift_id) REFERENCES gifts(id)
);
