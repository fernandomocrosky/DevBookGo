USE devbook;

INSERT INTO users 
    (name, nick, email, password)
VALUES
    ("user1", "user1", "user1@email.com", "$2a$10$EFE0HHkfqdMvCColjDVgbOo2OYtjkD4wqfvrnBKH/E8eHFRRVpPWG"),
    ("user2", "user2", "user2@email.com", "$2a$10$EFE0HHkfqdMvCColjDVgbOo2OYtjkD4wqfvrnBKH/E8eHFRRVpPWG"),
    ("user3", "user3", "user3@email.com", "$2a$10$EFE0HHkfqdMvCColjDVgbOo2OYtjkD4wqfvrnBKH/E8eHFRRVpPWG"),
    ("user4", "user4", "user4@email.com", "$2a$10$EFE0HHkfqdMvCColjDVgbOo2OYtjkD4wqfvrnBKH/E8eHFRRVpPWG")
;

INSERT INTO followers 
    (user_id, follower_id)
VALUES
    (1, 2),
    (3, 2),
    (4, 1),
    (1, 3)
;