
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    password VARCHAR(255),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

CREATE TABLE profile(
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(60),
    last_name VARCHAR(60),
    phone_number VARCHAR(13),
    image VARCHAR(255),
    user_id INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

CREATE TABLE movies(
    id SERIAL PRIMARY KEY,
    tittle VARCHAR(60),
    genre VARCHAR(100),
    images VARCHAR(255),
    synopsis VARCHAR(255),
    author VARCHAR(60),
    actors VARCHAR(255),
    release_date DATE,
    duration TIME,
    tag VARCHAR(30),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

INSERT INTO movies (tittle, genre, images, synopsis, author, actors, release_date, duration, tag) VALUES
('Avatar', 'Action, Adventure, Science Fiction', 'avatar.jpg', 'Di dunia Pandora, manusia berusaha mengeksploitasi sumber daya alam, sementara kelompok Navi berjuang untuk melindungi rumah mereka.', 'James Cameron', 'Sam Worthington, Zoe Saldana', '2009-12-18', '02:42:00', 'epic'),
('The Dark Knight', 'Action, Crime, Drama', 'dark_knight.jpg', 'Batman berusaha menghentikan Joker yang berencana menciptakan kekacauan di Gotham City.', 'Christopher Nolan', 'Christian Bale, Heath Ledger', '2008-07-18', '02:32:00', 'thriller'),
('Inception', 'Action, Adventure, Science Fiction', 'inception.jpg', 'Seorang pencuri yang dapat memasuki mimpi orang lain diupah untuk menanamkan ide ke pikiran seseorang.', 'Christopher Nolan', 'Leonardo DiCaprio, Joseph Gordon-Levitt', '2010-07-16', '02:28:00', 'mind-bending'),
('Titanic', 'Drama, Romance', 'titanic.jpg', 'Kisah cinta antara Jack dan Rose yang terjebak dalam tragedi tenggelamnya kapal Titanic.', 'James Cameron', 'Leonardo DiCaprio, Kate Winslet', '1997-12-19', '03:14:00', 'classic'),
('The Matrix', 'Action, Science Fiction', 'matrix.jpg', 'Seorang programmer mengetahui bahwa dunia tempat dia tinggal adalah simulasi buatan yang diciptakan oleh mesin.', 'Lana Wachowski, Lilly Wachowski', 'Keanu Reeves, Laurence Fishburne', '1999-03-31', '02:16:00', 'cyberpunk'),
('Forrest Gump', 'Drama, Romance', 'forrest_gump.jpg', 'Seorang pria sederhana mengalami berbagai peristiwa besar dalam sejarah Amerika, dari Perang Vietnam hingga Watergate.', 'Eric Roth', 'Tom Hanks, Robin Wright', '1994-07-06', '02:22:00', 'inspirational'),
('Avengers: Endgame', 'Action, Adventure, Science Fiction', 'avengers_endgame.jpg', 'Para Avengers berusaha untuk mengalahkan Thanos setelah kehancuran yang dia bawa di Avengers: Infinity War.', 'Anthony Russo, Joe Russo', 'Robert Downey Jr., Chris Evans', '2019-04-26', '03:02:00', 'superhero'),
('The Shawshank Redemption', 'Drama', 'shawshank.jpg', 'Seorang pria yang dihukum untuk kejahatan yang tidak dia lakukan berteman dengan sesama tahanan dan merencanakan pelarian.', 'Frank Darabont', 'Tim Robbins, Morgan Freeman', '1994-09-22', '02:22:00', 'hope'),
('Jurassic Park', 'Action, Adventure, Science Fiction', 'jurassic_park.jpg', 'Seorang ilmuwan menciptakan taman safari yang dihuni oleh dinosaurus, tetapi segalanya menjadi kacau ketika dinosaurus lepas.', 'Michael Crichton', 'Sam Neill, Laura Dern', '1993-06-11', '02:06:00', 'thriller'),
('The Godfather', 'Crime, Drama', 'godfather.jpg', 'Kisah keluarga mafia Corleone dan perjuangan Michael Corleone untuk mempertahankan kekuatan dan kehormatan keluarga.', 'Mario Puzo', 'Marlon Brando, Al Pacino', '1972-03-24', '02:55:00', 'classic'),
('The Lion King', 'Animation, Adventure, Drama', 'lion_king.jpg', 'Seorang anak singa berjuang untuk merebut kembali tahtanya setelah ayahnya dibunuh.', 'Roger Allers, Rob Minkoff', 'Matthew Broderick, Jeremy Irons', '1994-06-24', '01:28:00', 'animated'),
('Gladiator', 'Action, Adventure, Drama', 'gladiator.jpg', 'Seorang prajurit yang dikhianati berjuang untuk membalas dendam di arena gladiator Roma.', 'Ridley Scott', 'Russell Crowe, Joaquin Phoenix', '2000-05-05', '02:35:00', 'epic'),
('Star Wars: Episode V - The Empire Strikes Back', 'Action, Adventure, Fantasy', 'empire_strikes_back.jpg', 'Rebel Alliance berjuang melawan Kekaisaran Galaksi, sementara Luke Skywalker melanjutkan latihannya dengan Yoda.', 'Irvin Kershner', 'Mark Hamill, Harrison Ford', '1980-05-21', '02:04:00', 'sci-fi'),
('Back to the Future', 'Adventure, Comedy, Science Fiction', 'back_to_the_future.jpg', 'Seorang remaja bepergian ke masa lalu menggunakan mesin waktu dan bertemu orang tuanya yang muda.', 'Robert Zemeckis', 'Michael J. Fox, Christopher Lloyd', '1985-07-03', '01:56:00', 'time-travel'),
('The Avengers', 'Action, Adventure, Science Fiction', 'avengers.jpg', 'Tim superhero berusaha untuk bekerja sama dan melawan ancaman yang mengancam dunia.', 'Joss Whedon', 'Robert Downey Jr., Chris Hemsworth', '2012-05-04', '02:23:00', 'teamwork'),
('The Lord of the Rings: The Return of the King', 'Action, Adventure, Drama', 'lotr_return_of_king.jpg', 'Frodo dan Sam berusaha untuk menghancurkan cincin kekuasaan di Gunung Doom, sementara pasukan Rohan dan Gondor bersatu melawan Sauron.', 'Peter Jackson', 'Elijah Wood, Ian McKellen', '2003-12-17', '03:21:00', 'epic'),
('Pulp Fiction', 'Crime, Drama', 'pulp_fiction.jpg', 'Kisah beberapa karakter yang terhubung melalui serangkaian kejadian kriminal di Los Angeles.', 'Quentin Tarantino', 'John Travolta, Uma Thurman', '1994-10-14', '02:34:00', 'cult'),
('Forrest Gump', 'Drama, Romance', 'forrest_gump.jpg', 'Seorang pria sederhana mengalami berbagai peristiwa besar dalam sejarah Amerika, dari Perang Vietnam hingga Watergate.', 'Eric Roth', 'Tom Hanks, Robin Wright', '1994-07-06', '02:22:00', 'inspirational'),
('Schindler List', 'Biography, Drama, History', 'schindlers_list.jpg', 'Seorang pengusaha Jerman menyelamatkan lebih dari seribu orang Yahudi selama Holocaust.', 'Steven Zaillian', 'Liam Neeson, Ben Kingsley', '1993-12-15', '03:15:00', 'historical'),
('The Departed', 'Crime, Drama, Thriller', 'departed.jpg', 'Seorang polisi menyamar sebagai gangster, sementara seorang gangster menyamar sebagai polisi dalam perjuangan untuk melindungi kota Boston.', 'William Monahan', 'Leonardo DiCaprio, Matt Damon', '2006-10-06', '02:31:00', 'thriller'),
('The Silence of the Lambs', 'Crime, Drama, Thriller', 'silence_of_the_lambs.jpg', 'Seorang agen FBI bekerja sama dengan seorang pembunuh berantai untuk menangkap pembunuh lainnya.', 'Thomas Harris', 'Jodie Foster, Anthony Hopkins', '1991-02-14', '01:58:00', 'psychological'),
('The Prestige', 'Drama, Mystery, Science Fiction', 'prestige.jpg', 'Dua pesulap terlibat dalam persaingan sengit untuk menciptakan trik sulap yang lebih hebat.', 'Jonathan Nolan', 'Christian Bale, Hugh Jackman', '2006-10-20', '02:10:00', 'mystery'),
('The Dark Knight Rises', 'Action, Drama', 'dark_knight_rises.jpg', 'Batman kembali untuk melawan Bane yang berusaha menghancurkan Gotham City.', 'Christopher Nolan', 'Christian Bale, Tom Hardy', '2012-07-20', '02:44:00', 'superhero'),
('The Green Mile', 'Crime, Drama, Fantasy', 'green_mile.jpg', 'Seorang sipir penjara menjalin hubungan dengan seorang terpidana mati yang memiliki kekuatan penyembuhan luar biasa.', 'Frank Darabont', 'Tom Hanks, Michael Clarke Duncan', '1999-12-10', '03:09:00', 'emotional'),
('Interstellar', 'Adventure, Drama, Science Fiction', 'interstellar.jpg', 'Seorang ilmuwan dan timnya melakukan perjalanan antar bintang untuk mencari tempat tinggal baru bagi umat manusia.', 'Jonathan Nolan', 'Matthew McConaughey, Anne Hathaway', '2014-11-07', '02:49:00', 'space'),
('The Wizard of Oz', 'Adventure, Family, Fantasy', 'wizard_of_oz.jpg', 'Seorang gadis muda bertualang ke dunia fantasi untuk menemukan jalan pulang, bertemu teman-teman baru sepanjang perjalanan.', 'L. Frank Baum', 'Judy Garland, Frank Morgan', '1939-08-15', '01:42:00', 'classic'),
('Casablanca', 'Drama, Romance, War', 'casablanca.jpg', 'Seorang pemilik kafe di Casablanca menjadi terlibat dalam perjuangan cinta dan perjuangan melawan Nazi selama Perang Dunia II.', 'Julius J. Epstein', 'Humphrey Bogart, Ingrid Bergman', '1942-11-26', '01:42:00', 'classic'),
('The Shining', 'Horror, Mystery, Thriller', 'shining.jpg', 'Seorang penulis yang sedang mengalami masalah kejiwaan diisolasi bersama keluarganya di hotel terpencil yang penuh dengan kekuatan jahat.', 'Stephen King', 'Jack Nicholson, Shelley Duvall', '1980-05-23', '02:26:00', 'psychological'),
('Fight Club', 'Drama', 'fight_club.jpg', 'Seorang pria yang bosan dengan hidupnya memulai sebuah klub perkelahian bawah tanah yang mengarah pada peristiwa-peristiwa berbahaya.', 'Chuck Palahniuk', 'Brad Pitt, Edward Norton', '1999-10-15', '02:19:00', 'cult'),
('The Godfather: Part II', 'Crime, Drama', 'godfather_part2.jpg', 'Melanjutkan kisah keluarga Corleone, film ini memperlihatkan awal mula perjalanan Michael Corleone menjadi seorang pemimpin mafia.', 'Francis Ford Coppola', 'Al Pacino, Robert De Niro', '1974-12-20', '03:22:00', 'epic'),
('The Princess Bride', 'Adventure, Family, Fantasy', 'princess_bride.jpg', 'Seorang pemuda menceritakan kisah petualangan menakjubkan dengan aksi heroik, cinta, dan persahabatan kepada kakeknya.', 'William Goldman', 'Cary Elwes, Robin Wright', '1987-09-25', '01:38:00', 'classic'),
('The Godfather: Part III', 'Crime, Drama', 'godfather_part3.jpg', 'Michael Corleone berusaha untuk melepaskan diri dari dunia kejahatan dan menghadapi masa lalu keluarganya.', 'Mario Puzo', 'Al Pacino, Diane Keaton', '1990-12-25', '02:42:00', 'mafia'),
('The Revenant', 'Action, Drama, Adventure', 'revenant.jpg', 'Seorang pemburu berjuang untuk bertahan hidup setelah dikhianati dan diserang oleh beruang.', 'Alejandro González Iñárritu', 'Leonardo DiCaprio, Tom Hardy', '2015-12-25', '02:36:00', 'survival'),
('The Social Network', 'Biography, Drama', 'social_network.jpg', 'Kisah pendirian Facebook dan perjuangan Mark Zuckerberg untuk mempertahankan kontrol terhadap perusahaannya.', 'Aaron Sorkin', 'Jesse Eisenberg, Andrew Garfield', '2010-10-01', '02:00:00', 'biography'),
('Mad Max: Fury Road', 'Action, Adventure, Science Fiction', 'mad_max.jpg', 'Seorang wanita berusaha melarikan diri dari tiran di dunia pasca-apokaliptik, dengan bantuan seorang pembawa jalan.', 'George Miller', 'Tom Hardy, Charlize Theron', '2015-05-15', '02:00:00', 'post-apocalyptic'),
('The Usual Suspects', 'Crime, Drama, Mystery', 'usual_suspects.jpg', 'Seorang detektif mencoba mengungkap siapa yang bertanggung jawab atas kejahatan yang melibatkan lima penjahat yang telah berkumpul bersama.', 'Christopher McQuarrie', 'Kevin Spacey, Gabriel Byrne', '1995-08-16', '01:46:00', 'thriller'),
('The Terminator', 'Action, Science Fiction, Thriller', 'terminator.jpg', 'Seorang pembunuh cyborg dikirim kembali ke masa lalu untuk membunuh ibu pemimpin perlawanan manusia.', 'James Cameron', 'Arnold Schwarzenegger, Linda Hamilton', '1984-10-26', '01:47:00', 'action'),
('The Great Gatsby', 'Drama, Romance', 'great_gatsby.jpg', 'Kisah seorang pria misterius yang berusaha meraih kembali cinta lamanya dalam dunia penuh kemewahan dan dekadensi.', 'Baz Luhrmann', 'Leonardo DiCaprio, Carey Mulligan', '2013-05-10', '02:23:00', 'romantic'),
('The Grand Budapest Hotel', 'Comedy, Drama', 'grand_budapest_hotel.jpg', 'Seorang penjaga hotel yang terhormat berjuang untuk membuktikan bahwa dia tidak bersalah atas tuduhan pencurian sebuah lukisan langka.', 'Wes Anderson', 'Ralph Fiennes, F. Murray Abraham', '2014-03-28', '01:39:00', 'quirky'),
('The Martian', 'Adventure, Drama, Science Fiction', 'martian.jpg', 'Seorang astronaut yang terdampar di Mars berjuang untuk bertahan hidup dan mencoba untuk kembali ke Bumi.', 'Drew Goddard', 'Matt Damon, Jessica Chastain', '2015-10-02', '02:24:00', 'survival'),
('Guardians of the Galaxy', 'Action, Adventure, Comedy', 'guardians_of_the_galaxy.jpg', 'Sekelompok individu yang tidak biasa menjadi pahlawan galaksi dengan melawan ancaman besar terhadap alam semesta.', 'James Gunn', 'Chris Pratt, Zoe Saldana', '2014-08-01', '02:01:00', 'superhero');
CREATE Table cinema(
    id serial PRIMARY KEY,
    name VARCHAR(50),
    price INT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

INSERT INTO cinema (name, price) VALUES
('Cinepolis Mega Bekasi', 15000),
('CGV Grand Indonesia',15000),
('Cinemaxx Lippo Mall Puri',15000),
('XXI Plaza Indonesia',15000),
('Cineworld Mall Kelapa Gading',15000),
('CGV Pacific Place',15000),
('Cinema XXI Tunjungan Plaza',15000),
('Cinemaxx Surabaya',15000),
('CGV Cinere Mall',15000),
('Cinepolis Palembang',15000),
('XXI Mall of Indonesia',15000),
('CGV Blok M Square',15000),
('Cinemaxx Transmart Carrefour',15000),
('Cinema XXI Semanggi',15000),
('CGV Central Park',15000),
('Cinepolis Jakarta Garden City',15000),
('XXI Grand City Mall Surabaya',15000),
('Cinemaxx Park 23 Bali',15000),
('CGV Bogor Trade Mall',15000),
('Cinepolis Cilandak Town Square',15000),
('Cinemaxx Festival City Link Bandung',15000),
('CGV Epiwalk Mall',15000),
('Cinepolis Dago Plaza Bandung',15000),
('XXI Supermall Karawaci',15000),
('CGV Mall Kota Kasablanka',15000),
('Cinemaxx Mall Ciputra Cibubur',15000),
('Cinepolis Trans Studio Mall Makassar',15000),
('XXI Paris Van Java Bandung',15000),
('Cinepolis Bandung Supermall',15000),
('CGV La Piazza Mall',15000),
('Cinemaxx Living World Alam Sutera',15000),
('CGV Grand Indonesia',15000),
('Cinemaxx Lippo Mall Puri',15000),
('XXI Plaza Indonesia',15000),
('Cineworld Mall Kelapa Gading',15000),
('CGV Pacific Place',15000),
('Cinema XXI Tunjungan Plaza',15000),
('Cinemaxx Surabaya',15000),
('CGV Cinere Mall',15000);


CREATE TABLE cinema_date(
    id SERIAL PRIMARY KEY,
    name_date DATE,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

INSERT INTO cinema_date (name_date)
VALUES
    ('2025-01-01'),
    ('2025-01-02'),
    ('2025-01-03'),
    ('2025-01-04'),
    ('2025-01-05');

CREATE TABLE cinema_time(
    id SERIAL PRIMARY KEY,
    name_time TIME,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP 
);


INSERT INTO cinema_time (name_time)
VALUES
    ('10:00:00'),
    ('12:00:00'),
    ('14:00:00'),
    ('16:00:00'),
    ('18:00:00'),
    ('20:00:00'),
    ('22:00:00');

CREATE TABLE cinema_location(
    id SERIAL PRIMARY KEY,
    name_location VARCHAR(50),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);


INSERT INTO cinema_location (name_location) VALUES
('Jakarta'),
('Surabaya'),
('Bandung'),
('Medan'),
('Yogyakarta'),
('Makassar'),
('Semarang'),
('Palembang'),
('Bali'),
('Malang'),
('Bogor'),
('Tangerang'),
('Depok'),
('Bekasi');

CREATE TABLE movie_schedules (
    id serial primary key,
    movie_id int REFERENCES movies(id),
    cinema_id int REFERENCES cinema(id),
    date_id int REFERENCES cinema_date(id),
    time_id int REFERENCES cinema_time(id),
    location_id int REFERENCES cinema_location(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at timestamp
);

INSERT INTO movie_schedules (movie_id, cinema_id, date_id, time_id,
location_id) VALUES 
(1,1,1,1,1),
(1,2,1,1,1),
(1,3,1,1,1),
(1,4,1,1,1),
(1,5,1,1,1),
(1,1,2,1,1),
(1,1,3,1,1),
(1,1,4,1,1),
(1,1,5,1,1),
(1,1,1,2,1),
(1,1,1,3,1),
(1,1,1,4,1),
(1,1,1,5,1),
(2,1,1,1,1),
(2,2,1,1,1),
(2,3,1,1,1),
(2,4,1,1,1),
(2,5,1,1,1),
(2,1,2,1,1),
(2,1,3,1,1),
(2,1,4,1,1),
(2,1,5,1,1),
(2,1,1,2,1),
(2,1,1,3,1),
(2,1,1,4,1),
(2,1,1,5,1),
(3,1,1,1,1),
(3,2,1,1,1),
(3,3,1,1,1),
(3,4,1,1,1),
(3,5,1,1,1),
(3,1,2,1,1),
(3,1,3,1,1),
(3,1,4,1,1),
(3,1,5,1,1);

CREATE TABLE payment_method (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    rekening VARCHAR(20),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);



INSERT INTO payment_method (name, rekening) VALUES
('Goggle Pay', '1234567890123456'), ('Visa','2345678901234567'), 
('Gopay', '3456789012345678'), ('Pay Pal', '4567890123456789'),
('Dana', '3456789012345678'), ('Bank BCA', '2345678901234567'), 
('Bank BRI', '4567890123456789'), ('OVO', '1234567890123456');

CREATE TABLE status_payment(
    id SERIAL PRIMARY KEY,
    name VARCHAR(10),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);

INSERT INTO status_payment (name) VALUES
('Paid'), ('Not Paid');

CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    profile_id INT REFERENCES profile(id),
    movie_id INT REFERENCES movies(id),
    cinema_id INT REFERENCES cinema(id),
    payment_id INT REFERENCES payment_method(id),
    seat VARCHAR[],
    date_order DATE,
    qty INT,
    total_price INT,
    expired_payment DATE,
    status_id INT REFERENCES status_payment(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP
);