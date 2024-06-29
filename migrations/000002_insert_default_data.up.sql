INSERT INTO users (username, email, password_hash, created_at) VALUES
('johndoe', 'john.doe@email.com', 'hashed_password_1', NOW()),
('jane_smith', 'jane.smith@email.com', 'hashed_password_2', NOW()),
('mike_johnson', 'mike.johnson@email.com', 'hashed_password_3', NOW()),
('sarah_williams', 'sarah.williams@email.com', 'hashed_password_4', NOW()),
('david_brown', 'david.brown@email.com', 'hashed_password_5', NOW()),
('emily_davis', 'emily.davis@email.com', 'hashed_password_6', NOW()),
('chris_wilson', 'chris.wilson@email.com', 'hashed_password_7', NOW()),
('lisa_taylor', 'lisa.taylor@email.com', 'hashed_password_8', NOW()),
('robert_anderson', 'robert.anderson@email.com', 'hashed_password_9', NOW()),
('amanda_thomas', 'amanda.thomas@email.com', 'hashed_password_10', NOW());

INSERT INTO user_profiles (user_id, full_name, bio, role, location, avatar_image, website) VALUES
(1, 'John Doe', 'Passionate guitarist and songwriter', 'musician', 'New York, NY', NULL, 'www.johndoemusic.com'),
(2, 'Jane Smith', 'Avid music enthusiast and concert-goer', 'listener', 'Los Angeles, CA', NULL, NULL),
(3, 'Mike Johnson', 'Electronic music producer and DJ', 'producer', 'Miami, FL', NULL, 'www.mikejohnsonmusic.com'),
(4, 'Sarah Williams', 'Classical pianist and music teacher', 'musician', 'Boston, MA', NULL, 'www.sarahwilliamspiano.com'),
(5, 'David Brown', 'Hip-hop artist and beatmaker', 'musician', 'Atlanta, GA', NULL, 'www.davidbrownmusic.com'),
(6, 'Emily Davis', 'Indie folk singer-songwriter', 'musician', 'Portland, OR', NULL, 'www.emilydavismusic.com'),
(7, 'Chris Wilson', 'Music blog writer and critic', 'listener', 'Chicago, IL', NULL, 'www.chriswilsonreviews.com'),
(8, 'Lisa Taylor', 'Jazz vocalist and composer', 'musician', 'New Orleans, LA', NULL, 'www.lisataylorjazz.com'),
(9, 'Robert Anderson', 'Record producer and studio owner', 'producer', 'Nashville, TN', NULL, 'www.andersonstudios.com'),
(10, 'Amanda Thomas', 'Pop music fan and playlist curator', 'listener', 'Seattle, WA', NULL, NULL);