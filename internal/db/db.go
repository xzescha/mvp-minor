package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DB struct {
	Conn *sql.DB
}

func New(connStr string) (*DB, error) {
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}
	return &DB{Conn: conn}, nil
}

func (db *DB) GetTasks() ([]Task, error) {
	rows, err := db.Conn.Query("SELECT id, description FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Description); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (db *DB) GetGifts() ([]Gift, error) {
	rows, err := db.Conn.Query("SELECT id, name FROM gifts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gifts []Gift
	for rows.Next() {
		var g Gift
		if err := rows.Scan(&g.ID, &g.Name); err != nil {
			return nil, err
		}
		gifts = append(gifts, g)
	}
	return gifts, nil
}

func (db *DB) MarkTaskCompleted(userID int64, taskID int) error {
	_, err := db.Conn.Exec(`
        INSERT INTO user_tasks(user_id, task_id, completed)
        VALUES ($1, $2, TRUE)
        ON CONFLICT (user_id, task_id) DO UPDATE SET completed = TRUE`,
		userID, taskID)
	return err
}

func (db *DB) AssignGiftToUser(userID int64, giftID int) error {
	_, err := db.Conn.Exec(`
        INSERT INTO user_gifts(user_id, gift_id)
        VALUES ($1, $2)
        ON CONFLICT DO NOTHING`, userID, giftID)
	return err
}

type Task struct {
	ID          int
	Description string
}

type Gift struct {
	ID   int
	Name string
}
