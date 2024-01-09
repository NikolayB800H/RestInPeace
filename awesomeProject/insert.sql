DROP TABLE users CASCADE;
DROP TABLE data_types CASCADE;
DROP TABLE forecast_applications CASCADE;
DROP TABLE connector_apps_types CASCADE;

INSERT INTO users(user_id, login, password, role)
VALUES ('796c70e1-5f27-4433-a415-95e7272effa5' , 'moderator', '40bd001563085fc35165329ea1ff5c5ecbdbbeef', 2),
       ('5f58c307-a3f2-4b13-b888-c80ad08d5ed3' , 'nop', '40bd001563085fc35165329ea1ff5c5ecbdbbeef', 1);

INSERT INTO data_types(data_type_id, image_path, data_type_name, precision, description, unit, data_type_status)
VALUES ('a20163ce-7be5-46ec-a50f-a313476b2bd1' , 'http://localhost:9000/images/a20163ce-7be5-46ec-a50f-a313476b2bd1.jpg', 'температуры', 5.0, 'Наши термометры самые точные в мире!!!! Купи прогноз, не пожалеешь!', '℃', 'доступен'),
       ('0706419e-b024-469d-a354-9480cd79c6a5' , 'http://localhost:9000/images/0706419e-b024-469d-a354-9480cd79c6a5.jpg', 'давления', 4.2, 'Наши манометры самые точные в мире!!!! Купи прогноз, не пожалеешь!', 'мм рт. ст.', 'доступен'),
       ('8f157a95-dad1-43e0-9372-93b51de06163' , 'http://localhost:9000/images/8f157a95-dad1-43e0-9372-93b51de06163.jpg', 'влажности', 6.6, 'Наши гигрометры самые точные в мире!!!! Купи прогноз, не пожалеешь!', '%', 'доступен');
