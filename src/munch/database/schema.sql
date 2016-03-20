DROP TABLE IF EXISTS `broccoli_sources`;
CREATE TABLE IF NOT EXISTS `broccoli_sources` (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    type INT NOT NULL,
    lastcrawled INTEGER NOT NULL
);
INSERT INTO broccoli_sources(name, url, type, lastcrawled)
    VALUES ("Techcrunch", "http://feeds.feedburner.com/TechCrunch/?fmt=xml", 0, -1),
           ("VentureBeat", "http://feeds.venturebeat.com/VentureBeat", 0, -1),
           ("MKBHD", "http://youtube.com/user/marquesbrownlee/", 2, -1);

DROP TABLE IF EXISTS `broccoli_stories`;
CREATE TABLE IF NOT EXISTS `broccoli_stories` (
    id INTEGER PRIMARY KEY,
    url TEXT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    pubdate INTEGER NOT NULL,
    source_id INTEGER NOT NULL,
    FOREIGN KEY(source_id) REFERENCES broccoli_sources(id)
);

DROP TABLE IF EXISTS `broccoli_users`;
CREATE TABLE IF NOT EXISTS `broccoli_users` (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL
);

DROP TABLE IF EXISTS `broccoli_newsletters`;
CREATE TABLE IF NOT EXISTS `broccoli_newsletters` (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    pubdate INTEGER NOT NULL
);
INSERT INTO broccoli_newsletters(title, pubdate) VALUES ("Utk's AI Newsletter", -1);

DROP TABLE IF EXISTS `broccoli_newsletters_sources`;
CREATE TABLE IF NOT EXISTS `broccoli_newsletters_sources` (
    id INTEGER PRIMARY KEY,
    newsletter_id INTEGER NOT NULL,
    source_id INTEGER NOT NULL,
    source_lastchecked INTEGER NOT NULL,
    FOREIGN KEY(newsletter_id) REFERENCES broccoli_newsletters(id),
    FOREIGN KEY(source_id) REFERENCES broccoli_sources(id)
);
INSERT INTO broccoli_newsletters_sources(newsletter_id, source_id, source_lastchecked)
    VALUES (1, 1, -1);

DROP TABLE IF EXISTS `broccoli_users_newsletters`;
CREATE TABLE IF NOT EXISTS `broccoli_users_newsletters` (
    user_id INTEGER NOT NULL,
    newsletter_id INTEGER NOT NULL,
    FOREIGN KEY(user_id) REFERENCES broccoli_users(id),
    FOREIGN KEY(newsletter_id) REFERENCES broccoli_newsletters(id)
);
INSERT INTO broccoli_users_newsletters(user_id,newsletter_id) VALUES (1, 1);
