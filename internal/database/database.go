package database

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"net/url"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	_ "modernc.org/sqlite"
)

// migrations contains versioned SQL files shipped with the application.
//
//go:embed migrations/*.sql
var migrations embed.FS

// Open initializes SQLite, applies migrations, and returns a database handle.
func Open(dataDir string) (*sql.DB, error) {
	dbPath := filepath.Join(dataDir, "doubao.db")
	dsn := (&url.URL{Scheme: "file", Path: dbPath}).String()

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", dbPath, err)
	}

	// This is a single-user application. A single connection keeps SQLite locking
	// behavior predictable while WAL still makes reads friendly to future features.
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping %s: %w", dbPath, err)
	}
	for _, pragma := range []string{
		"PRAGMA busy_timeout = 5000",
		"PRAGMA foreign_keys = ON",
		"PRAGMA journal_mode = WAL",
	} {
		if _, err := db.Exec(pragma); err != nil {
			db.Close()
			return nil, fmt.Errorf("configure sqlite with %q: %w", pragma, err)
		}
	}
	if err := migrate(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("migrate database: %w", err)
	}
	return db, nil
}

func migrate(db *sql.DB) error {
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			applied_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`); err != nil {
		return err
	}

	entries, err := fs.ReadDir(migrations, "migrations")
	if err != nil {
		return err
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}
		version, err := migrationVersion(entry.Name())
		if err != nil {
			return err
		}

		var applied int
		if err := db.QueryRow("SELECT COUNT(1) FROM schema_migrations WHERE version = ?", version).Scan(&applied); err != nil {
			return err
		}
		if applied > 0 {
			continue
		}

		script, err := fs.ReadFile(migrations, path.Join("migrations", entry.Name()))
		if err != nil {
			return err
		}
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		if _, err = tx.Exec(string(script)); err != nil {
			tx.Rollback()
			return fmt.Errorf("apply %s: %w", entry.Name(), err)
		}
		if _, err = tx.Exec("INSERT INTO schema_migrations (version, name) VALUES (?, ?)", version, entry.Name()); err != nil {
			tx.Rollback()
			return fmt.Errorf("record %s: %w", entry.Name(), err)
		}
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("commit %s: %w", entry.Name(), err)
		}
	}
	return nil
}

func migrationVersion(name string) (int, error) {
	underscore := strings.IndexByte(name, '_')
	if underscore <= 0 {
		return 0, fmt.Errorf("migration filename must start with a numeric version: %s", name)
	}
	version, err := strconv.Atoi(name[:underscore])
	if err != nil {
		return 0, fmt.Errorf("parse migration version %s: %w", name, err)
	}
	return version, nil
}
