CREATE TABLE tipe_sensor (
    id SERIAL NOT NULL PRIMARY KEY,
    tipe VARCHAR(50) NOT NULL
);

CREATE TABLE informasi_sensor (
    id SERIAL NOT NULL PRIMARY KEY,
    status BOOLEAN NOT NULL,
    identity VARCHAR(30) NOT NULL
);

CREATE TABLE monitoring_location (
    id SERIAL NOT NULL PRIMARY KEY,
    nama VARCHAR(50) NOT NULL UNIQUE,
    provinsi VARCHAR(50) NOT NULL,
    kecamatan VARCHAR(50) NOT NULL,
    desa VARCHAR(50) NOT NULL
);

CREATE TABLE sensors (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    inf_sensor_id INT NOT NULL,
    tipe_sensor_id INT NOT NULL,
    mon_loc_id INT NOT NULL,
    ditempatkan_pada TIMESTAMP NOT NULL,

    CONSTRAINT tipe_sensor
        FOREIGN KEY (tipe_sensor_id)
        REFERENCES tipe_sensor(id),
    
    CONSTRAINT informasi_sensor
        FOREIGN KEY (inf_sensor_id)
        REFERENCES informasi_sensor(id),

    CONSTRAINT monitoring_location
        FOREIGN KEY (mon_loc_id)
        REFERENCES monitoring_location(id)
);

CREATE TABLE value_sensor (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    sensor_id INT NOT NULL,
    data FLOAT NOT NULL,
    dibuat_pada TIMESTAMP NOT NULL,

    CONSTRAINT sensors
        FOREIGN KEY (sensor_id)
        REFERENCES sensors(id)
);