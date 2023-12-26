TRUNCATE TABLE users CASCADE;
TRUNCATE TABLE data_types CASCADE;
TRUNCATE TABLE forecast_applications CASCADE;
TRUNCATE TABLE connector_apps_types CASCADE;

INSERT INTO users(user_id, login, password, role)
VALUES ('796c70e1-5f27-4433-a415-95e7272effa5' , 'moderator', 'pass_moderator', 2),
       ('5f58c307-a3f2-4b13-b888-c80ad08d5ed3' , 'client', 'pass_client', 1);

INSERT INTO data_types(data_type_id, image_path, data_type_name, precision, description, unit, data_type_status)
VALUES ('a20163ce-7be5-46ec-a50f-a313476b2bd1' , 'localhost:9000/images/a20163ce-7be5-46ec-a50f-a313476b2bd1.jpg', 'температуры', 5.0, 'Наши термометры самые точные в мире!!!! Купи прогноз, не пожалеешь!', '℃', 'доступен'),
       ('0706419e-b024-469d-a354-9480cd79c6a5' , 'localhost:9000/images/0706419e-b024-469d-a354-9480cd79c6a5.jpg', 'давления', 4.2, 'Наши манометры самые точные в мире!!!! Купи прогноз, не пожалеешь!', 'мм рт. ст.', 'доступен'),
       ('8f157a95-dad1-43e0-9372-93b51de06163' , 'localhost:9000/images/8f157a95-dad1-43e0-9372-93b51de06163.jpg', 'влажности', 6.6, 'Наши гигрометры самые точные в мире!!!! Купи прогноз, не пожалеешь!', '%', 'доступен');

INSERT INTO forecast_applications(application_id, application_status, application_creation_date, application_formation_date, application_completion_date, creator_id, moderator_id, input_start_date)
VALUES ('b0247ccd-28ab-45be-9680-f24213cf7aab', 'удалён', 'CalculateFailed', '2023-10-25', '2023-10-25', '2023-10-25', '5f58c307-a3f2-4b13-b888-c80ad08d5ed3', '796c70e1-5f27-4433-a415-95e7272effa5', '2023-10-22');

INSERT INTO connector_apps_types(data_type_id, application_id, input_first, input_second, input_third, output)
VALUES ('0706419e-b024-469d-a354-9480cd79c6a5', 'b0247ccd-28ab-45be-9680-f24213cf7aab', 750.0, 740.0, 760.0, NULL);
