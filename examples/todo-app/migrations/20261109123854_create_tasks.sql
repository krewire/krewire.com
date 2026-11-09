-- Migration: create_tasks
-- Created at: 2026-11-09T12:38:54Z

-- +up
CREATE TABLE IF NOT EXISTS tasks (
  id         INTEGER PRIMARY KEY AUTOINCREMENT,
  title      TEXT     NOT NULL,
  done       INTEGER  NOT NULL DEFAULT 0,
  created_at DATETIME NOT NULL
);

-- +down
DROP TABLE IF EXISTS tasks;
