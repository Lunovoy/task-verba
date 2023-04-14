CREATE TABLE IF NOT EXISTS users (
    gender VARCHAR(10),
    title VARCHAR(10),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    street_number INT,
    street_name VARCHAR(50),
    city VARCHAR(50),
    state VARCHAR(50),
    country VARCHAR(50),
    postcode INT,
    latitude VARCHAR(20),
    longitude VARCHAR(20),
    timezone_offset VARCHAR(10),
    timezone_description VARCHAR(50),
    email VARCHAR(100),
    username VARCHAR(50),
    password VARCHAR(50),
    salt VARCHAR(50),
    md5 VARCHAR(50),
    sha1 VARCHAR(50),
    sha256 VARCHAR(100),
    dob_date TIME,
    dob_age INT,
    registered_date TIME,
    registered_age INT,
    phone VARCHAR(20),
    cell VARCHAR(20),
    id_name VARCHAR(20),
    id_value VARCHAR(40),
    picture_large VARCHAR(200),
    picture_medium VARCHAR(200),
    picture_thumbnail VARCHAR(200),
    nat VARCHAR(5)
);