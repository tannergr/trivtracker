DROP Table suggestedplaces;
CREATE TABLE suggestedplaces (
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

ALTER TABLE suggestedplaces ADD PRIMARY KEY (ID);
