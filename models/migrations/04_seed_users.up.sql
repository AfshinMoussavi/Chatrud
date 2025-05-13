-- models/migrations/000002_seed_users.up.sql
INSERT INTO users (id, name, email, phone, password, active) VALUES
    ('1', 'test1', 'test1@gmail.com', '09111111112', 'hashed_password', TRUE),
    ('2', 'test2', 'test2@gmail.com', '09222222221', 'hashed_password', TRUE)
;