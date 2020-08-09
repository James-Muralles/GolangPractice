
INSERT INTO dogs(name, breed) 
VALUES
('Hank', 'Boxer'),
('Spot', 'Dalmation'),
('Julie', 'Pitbull')
RETURNING id;

INSERT INTO dogs(name, breed) 
VALUES
('Bob', 'Bull Dog'),
('Ed', 'Beagle'),
('July', 'Pitbull')
RETURNING id;

