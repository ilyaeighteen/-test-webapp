create schema if not exists app authorization db_user;

create table app.user (
        id SERIAL PRIMARY KEY,
		name TEXT,
		created_at TIMESTAMPTZ not null DEFAULT CURRENT_TIMESTAMP
);