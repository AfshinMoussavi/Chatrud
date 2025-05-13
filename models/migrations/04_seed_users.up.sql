-- models/migrations/000002_seed_users.up.sql
INSERT INTO users (id, name, email, phone, password, active) VALUES
    ('1', 'Afshin', 'afshin@gmail.com', '09028657218', 'hashed_password', TRUE),
    ('2', 'Reza', 'reza@gmail.com', '09031305593', 'hashed_password', TRUE)
;