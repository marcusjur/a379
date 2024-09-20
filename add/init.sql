CREATE TABLE IF NOT EXISTS a379.animals (
    id INT AUTO_INCREMENT,
    name VARCHAR(255),
    tags VARCHAR(255),
    owner VARCHAR(255),
    image VARCHAR(255),
    PRIMARY KEY (id)
    );

INSERT INTO a379.animals (name, tags, owner, image)
VALUES (
           'Perry',
           'platypus,water',
           'Alice Smith',
           'https://example.com/image.jpeg'
       );

INSERT INTO a379.animals (name, tags, owner, image)
VALUES (
           'Garry',
           'snail',
           'Jane Doe',
           'https://example.com/image.jpeg'
       );

INSERT INTO a379.animals (name, tags, owner, image)
VALUES (
           'Xi',
           'dog,pet',
           'Zhang Wei',
           'https://example.com/image.jpeg'
       );
