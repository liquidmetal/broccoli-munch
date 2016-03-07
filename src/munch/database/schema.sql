CREATE TABLE IF NOT EXISTS `broccoli_stories` (
    id INTEGER PRIMARY KEY,
    url TEXT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    pubdate INTEGER NOT NULL
);

DROP TABLE IF EXISTS `broccoli_sources`;
CREATE TABLE IF NOT EXISTS `broccoli_sources` (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    type INT NOT NULL,
    lastcrawled INTEGER NOT NULL
);

INSERT INTO broccoli_sources(name, url, type, lastcrawled) VALUES ("Techcrunch", "http://feeds.feedburner.com/TechCrunch/?fmt=xml", 0, -1);
