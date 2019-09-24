CREATE TABLE
    passwords
    (
        id CHARACTER VARYING(36) NOT NULL,
        title CHARACTER VARYING(256) NOT NULL,
        url CHARACTER VARYING(2048),
        username CHARACTER VARYING(64) NOT NULL,
        password CHARACTER VARYING(1024) NOT NULL,
        notes CHARACTER VARYING(2048),
        tags CHARACTER VARYING(256),
        CONSTRAINT pk_id PRIMARY KEY (id)
    ); 
