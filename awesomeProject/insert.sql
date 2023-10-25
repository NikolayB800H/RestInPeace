INSERT INTO users(user_id, login, password, is_moderator)
VALUES (1 , 'moderator', 'pass_moderator', TRUE),
       (2 , 'client', 'pass_client', FALSE);

--SELECT SETVAL('public."users_id_seq"', COALESCE(MAX(user_id), 1)) FROM public."users";

INSERT INTO data_types(data_type_id, image_path, data_type_name, precision, description, unit, data_type_status)
VALUES (1 , 'term.svg', 'температуры', 5.0, 'Наши термометры самые точные в мире!!!! Купи прогноз, не пожалеешь!', '℃', 'valid'),
       (2 , 'gau.svg', 'давления', 4.2, 'Наши манометры самые точные в мире!!!! Купи прогноз, не пожалеешь!', 'мм рт. ст.', 'valid'),
       (3 , 'rain.svg', 'влажности', 6.6, 'Наши гигрометры самые точные в мире!!!! Купи прогноз, не пожалеешь!', '%', 'valid');

--SELECT SETVAL('public."data_types_id_seq"', COALESCE(MAX(data_type_id), 1)) FROM public."data_types";

--INSERT INTO forecast_applications(application_id, application_status, application_creation_date, application_formation_date, application_completion_date, creator_id, moderator_id)
--VALUES ();

--SELECT SETVAL('public."forecast_applications_id_seq"', COALESCE(MAX(application_id), 1)) FROM public."forecast_applications";

--INSERT INTO connector_apps_types(data_type_id, application_id, input_first, input_second, input_third, output)
--VALUES ();

--SELECT SETVAL('public."connector_apps_types_id_seq"', COALESCE(MAX(id), 1)) FROM public."connector_apps_types";
