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
           'platypus:species,water,2:age',
           'Alice Smith',
           'https://animals.sandiegozoo.org/sites/default/files/2019-04/animals_hero-platypus.jpg'
       );

INSERT INTO a379.animals (name, tags, owner, image)
VALUES (
           'Garry',
           'snail:species,4month:age',
           'Jane Doe',
           'https://avatars.dzeninfra.ru/get-zen_doc/5212789/pub_62b410d2ffe61d7b2371fbc3_62b413bf6447426295699960/scale_1200'
       );

INSERT INTO a379.animals (name, tags, owner, image)
VALUES (
           'Xi',
           'dog:species,pet,friendly,6:age',
           'Zhang Wei',
           'https://miro.medium.com/v2/resize:fit:1400/1*rIkmavUeqyRySwlQdA9kKg.jpeg'
       );

INSERT INTO a379.animals (name, tags, owner, image)
VALUES (
           'William',
           'seal:species,water,3:age',
           'Harry Cantrell',
           'https://m.media-amazon.com/images/I/61hCRkXrA5L._AC_UF894,1000_QL80_.jpg'
       );

INSERT INTO a379.animals (name, tags, owner, image)
VALUES (
           'Gleb',
           'cat:species,pet,12:age',
           'Maria Smith',
           'https://goood.pw/assets/cats/pwgood.png'
       );
