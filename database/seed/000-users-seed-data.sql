INSERT INTO auth."user" (first_name, last_name, email, password)
VALUES ('test', 'user', 'test@cmichael.dev', digest('testpass', 'sha256'));
