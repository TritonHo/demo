
insert into users(id, email, password_digest, first_name, last_name)
values
('eeee1df4-9fae-4e32-98c1-88f850a00001', 'Susan.Wong@abc.com', '$2a$08$I3jbICL3r6DhriF8Ildpa.uaaCLavZEV9.JLUcQRkZxfLXCTcSAwO', 'Susan', 'Wong'),
('eeee1df4-9fae-4e32-98c1-88f850a00002', 'Betty.Ho@abc.com', '$2a$08$I3jbICL3r6DhriF8Ildpa.uaaCLavZEV9.JLUcQRkZxfLXCTcSAwO', 'Susan', 'Wong');



insert into cats(id, user_id, name, gender) 
values
('ffff1df4-9fae-4e32-98c1-88f850a00001', 'eeee1df4-9fae-4e32-98c1-88f850a00001', 'LittleWhite', 'FEMALE'),
('ffff1df4-9fae-4e32-98c1-88f850a00002', 'eeee1df4-9fae-4e32-98c1-88f850a00002', 'Mikimiki', 'FEMALE');



