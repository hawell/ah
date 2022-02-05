DELETE FROM Provider;

INSERT INTO Provider VALUES (NULL, 'provider1', ST_GeomFromText('POINT(-26.66119 40.95858)'), 10.0, 3.5, 1, 1, 1);
INSERT INTO Provider VALUES (NULL, 'provider2', ST_GeomFromText('POINT(-26.66120 40.95858)'), 10.0, 4.5, 0, 1, 1);
INSERT INTO Provider VALUES (NULL, 'provider3', ST_GeomFromText('POINT(-26.66116 40.95858)'), 10.0, 4.5, 1, 0, 0);
INSERT INTO Provider VALUES (NULL, 'provider4', ST_GeomFromText('POINT(-26.66117 40.95858)'), 10.0, 4.7, 1, 1, 0);
INSERT INTO Provider VALUES (NULL, 'provider5', ST_GeomFromText('POINT(-26.66115 40.95858)'), 10.0, 4.5, 1, 0, 0);
INSERT INTO Provider VALUES (NULL, 'provider6', ST_GeomFromText('POINT(-26.66118 40.95858)'), 2.0, 4.1, 1, 0, 1);
INSERT INTO Provider VALUES (NULL, 'provider7', ST_GeomFromText('POINT(-26.66116 40.95858)'), 10.0, 4.8, 1, 0, 0);
