/*for normal tables */
GRANT SELECT, INSERT, UPDATE, DELETE, REFERENCES ON TABLE cats                  to demo_user;
GRANT SELECT, INSERT, UPDATE, DELETE, REFERENCES ON TABLE users                 to demo_user;

GRANT SELECT ON TABLE cats                  to demo_readonly;
GRANT SELECT ON TABLE users                 to demo_readonly;

