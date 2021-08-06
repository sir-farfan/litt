CREATE DATABASE IF NOT EXISTS litt;
USE litt;

CREATE TABLE IF NOT EXISTS  tag(
    id INT NOT NULL AUTO_INCREMENT,
    tag VARCHAR(40) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS member(
    id INT NOT NULL AUTO_INCREMENT,
    full_name VARCHAR(200),
    PRIMARY KEY (id)
    /* type */
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS member_tags(
    member_id INT,
    tag_id INT,
    FOREIGN KEY (member_id)
        REFERENCES member(id),
    FOREIGN KEY (tag_id)
        REFERENCES tag(id)
) ENGINE=INNODB;

