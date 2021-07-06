DROP TABLE IF EXISTS video;
CREATE TABLE IF NOT EXISTS video(
    user VARCHAR(255),
    videoId VARCHAR(255),
    published_at TIMESTAMP NOT NULL,
    channelID VARCHAR(255),
    title VARCHAR(255),
    description VARCHAR(255),
    channelTitle VARCHAR(255)
);