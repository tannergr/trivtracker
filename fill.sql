DROP Table places;
CREATE TABLE Places (
  ID    SERIAL UNIQUE,
  MapsID varchar(255) NOT NULL,
  Barname   varchar(255) NOT NULL,
  eType  varchar(255) NOT NULL,
  DayOfWeek varchar(255) NOT NULL,
  Comments varchar(255) NOT NULL,
  Lat   Numeric(8,5) NOT NULL,
  Long  Numeric(8,5) NOT NULL,
  CHECK (Barname <> ''),
  CHECK (eType <> ''),
  CHECK (DayOfWeek <> '')
);

INSERT INTO Places (MapsID, Barname, eType, DayOfWeek, Comments, Lat, Long) VALUES
('6c9a3577b23b9d4eec9dfb3b1054c819e5f71883', 'Kitchen Table', 'Trivia', 'Thursday', 'Hello world!', 49.271417, -123.15483);

ALTER TABLE Places ADD PRIMARY KEY (ID);
