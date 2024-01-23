--
-- PostgreSQL database dump
--

-- Dumped from database version 15.5 (Ubuntu 15.5-1.pgdg22.04+1)
-- Dumped by pg_dump version 15.5 (Ubuntu 15.5-1.pgdg22.04+1)

-- Started on 2024-01-23 09:17:17 MSK

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 217 (class 1259 OID 43747)
-- Name: connector_apps_types; Type: TABLE; Schema: public; Owner: nop
--

CREATE TABLE public.connector_apps_types (
    data_type_id uuid DEFAULT gen_random_uuid() NOT NULL,
    application_id uuid DEFAULT gen_random_uuid() NOT NULL,
    input_first numeric,
    input_second numeric,
    input_third numeric,
    output numeric
);


ALTER TABLE public.connector_apps_types OWNER TO nop;

--
-- TOC entry 215 (class 1259 OID 43723)
-- Name: data_types; Type: TABLE; Schema: public; Owner: nop
--

CREATE TABLE public.data_types (
    data_type_id uuid DEFAULT gen_random_uuid() NOT NULL,
    image_path character varying(256),
    data_type_name character varying(128) NOT NULL,
    "precision" numeric NOT NULL,
    description character varying(1024) NOT NULL,
    unit character varying(32) NOT NULL,
    data_type_status character varying(50) NOT NULL
);


ALTER TABLE public.data_types OWNER TO nop;

--
-- TOC entry 216 (class 1259 OID 43731)
-- Name: forecast_applications; Type: TABLE; Schema: public; Owner: nop
--

CREATE TABLE public.forecast_applications (
    application_id uuid DEFAULT gen_random_uuid() NOT NULL,
    application_status character varying(50) NOT NULL,
    calculate_status character varying(50),
    application_creation_date timestamp without time zone NOT NULL,
    application_formation_date timestamp without time zone,
    application_completion_date timestamp without time zone,
    creator_id uuid NOT NULL,
    moderator_id uuid,
    input_start_date date
);


ALTER TABLE public.forecast_applications OWNER TO nop;

--
-- TOC entry 214 (class 1259 OID 43715)
-- Name: users; Type: TABLE; Schema: public; Owner: nop
--

CREATE TABLE public.users (
    user_id uuid DEFAULT gen_random_uuid() NOT NULL,
    login character varying(256) NOT NULL,
    password character varying(256) NOT NULL,
    role bigint
);


ALTER TABLE public.users OWNER TO nop;

--
-- TOC entry 3389 (class 0 OID 43747)
-- Dependencies: 217
-- Data for Name: connector_apps_types; Type: TABLE DATA; Schema: public; Owner: nop
--

COPY public.connector_apps_types (data_type_id, application_id, input_first, input_second, input_third, output) FROM stdin;
8f157a95-dad1-43e0-9372-93b51de06163	bfc51b71-fc57-4125-85b6-75d3ec1f5097	12	14	16	18
a20163ce-7be5-46ec-a50f-a313476b2bd1	bfc51b71-fc57-4125-85b6-75d3ec1f5097	14	12	16	16
0706419e-b024-469d-a354-9480cd79c6a5	508ba876-f11c-4971-aa26-9363ee794be9	700	750	800	850
8f157a95-dad1-43e0-9372-93b51de06163	508ba876-f11c-4971-aa26-9363ee794be9	66	44	55	44.000000000000014
8f157a95-dad1-43e0-9372-93b51de06163	634d0ac9-8177-43b7-9d25-37a8080083ed	55	44	77	80.66666666666666
8f157a95-dad1-43e0-9372-93b51de06163	21045f7a-bbbe-40ad-9b57-efd73c505616	12	32	34	47.99999999999999
8f157a95-dad1-43e0-9372-93b51de06163	3f1f218e-a849-423f-9e6d-55748d7279da	23	45	68	90.33333333333331
8f157a95-dad1-43e0-9372-93b51de06163	4fe22a66-9f30-41cf-8b02-80ecff31b869	\N	\N	\N	\N
a20163ce-7be5-46ec-a50f-a313476b2bd1	656295a5-347b-46a5-9b4c-ad2cac9af2a9	\N	\N	\N	\N
a20163ce-7be5-46ec-a50f-a313476b2bd1	76f2dcb8-0af4-4046-ada4-98eb0d6217c4	\N	\N	\N	\N
a20163ce-7be5-46ec-a50f-a313476b2bd1	d168f9d6-167a-4f23-b2eb-88db0e3e0010	\N	\N	\N	\N
a20163ce-7be5-46ec-a50f-a313476b2bd1	db789d42-c443-4c76-9bcb-61782d2d6e00	\N	\N	\N	\N
a20163ce-7be5-46ec-a50f-a313476b2bd1	65ca2cfa-977a-40e9-82ab-c5fe770c330a	\N	\N	\N	\N
a20163ce-7be5-46ec-a50f-a313476b2bd1	4d5250f7-945e-49b6-ae85-ea83a9a9a1a6	\N	\N	\N	\N
a20163ce-7be5-46ec-a50f-a313476b2bd1	8bfcd218-3658-42e5-962b-6545b51454c6	\N	\N	\N	\N
a20163ce-7be5-46ec-a50f-a313476b2bd1	b938f662-27dd-4a8f-a32c-4c5c01e8ef75	\N	\N	\N	\N
\.


--
-- TOC entry 3387 (class 0 OID 43723)
-- Dependencies: 215
-- Data for Name: data_types; Type: TABLE DATA; Schema: public; Owner: nop
--

COPY public.data_types (data_type_id, image_path, data_type_name, "precision", description, unit, data_type_status) FROM stdin;
0706419e-b024-469d-a354-9480cd79c6a5	http://localhost:9000/images/0706419e-b024-469d-a354-9480cd79c6a5.png	давления	4.2	Наши манометры самые точные в мире!!!! Купи прогноз, не пожалеешь!	мм рт. ст.	доступен
8f157a95-dad1-43e0-9372-93b51de06163	http://localhost:9000/images/8f157a95-dad1-43e0-9372-93b51de06163.png	влажности	6.6	Наши гигрометры самые точные в мире!!!! Купи прогноз, не пожалеешь!	%	доступен
a20163ce-7be5-46ec-a50f-a313476b2bd1	http://localhost:9000/images/a20163ce-7be5-46ec-a50f-a313476b2bd1.png	температуры	5	Наши термометры самые точные в мире!!!! Купи прогноз, не пожалеешь!	℃	доступен
\.


--
-- TOC entry 3388 (class 0 OID 43731)
-- Dependencies: 216
-- Data for Name: forecast_applications; Type: TABLE DATA; Schema: public; Owner: nop
--

COPY public.forecast_applications (application_id, application_status, calculate_status, application_creation_date, application_formation_date, application_completion_date, creator_id, moderator_id, input_start_date) FROM stdin;
970c3af2-df75-4db8-875b-7142ae6c7e00	удалён	\N	2024-01-09 14:29:03.804505	\N	\N	796c70e1-5f27-4433-a415-95e7272effa5	\N	\N
94cf97e4-fe32-455f-94e0-58d37fce98ec	удалён	\N	2024-01-09 16:30:05.611797	\N	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	\N
db789d42-c443-4c76-9bcb-61782d2d6e00	удалён	\N	2024-01-19 16:47:40.505359	\N	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-16
65ca2cfa-977a-40e9-82ab-c5fe770c330a	удалён	\N	2024-01-19 16:49:00.104784	\N	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-22
79309c3c-3879-4fe8-ae91-a423db1a895b	завершён	прогноз рассчитан	2024-01-09 16:36:23.077597	2024-01-09 16:37:02.903461	2024-01-09 16:54:03.661717	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	796c70e1-5f27-4433-a415-95e7272effa5	2024-01-05
f2c1afde-a813-4b60-9288-c4c00fd13da0	отклонён	прогноз рассчитан	2024-01-09 16:37:22.708708	2024-01-09 16:37:46.677523	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	796c70e1-5f27-4433-a415-95e7272effa5	2024-01-01
4d5250f7-945e-49b6-ae85-ea83a9a9a1a6	удалён	\N	2024-01-19 16:50:09.96871	\N	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-18
8bfcd218-3658-42e5-962b-6545b51454c6	удалён	\N	2024-01-19 16:58:18.225955	\N	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	0001-01-01
7bd1f13d-3ab3-4ace-ba9d-02035d009294	сформирован	прогноз рассчитан	2024-01-09 16:55:52.262884	2024-01-09 16:56:15.362491	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-02
b938f662-27dd-4a8f-a32c-4c5c01e8ef75	черновик	\N	2024-01-19 16:58:41.34518	\N	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-17
bfc51b71-fc57-4125-85b6-75d3ec1f5097	сформирован	прогноз рассчитан	2024-01-09 17:28:00.137337	2024-01-09 17:31:14.277934	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-09
22198c3d-8156-43ee-83da-ddfd0ef5ff5d	завершён	прогноз рассчитан	2024-01-09 16:54:57.264267	2024-01-09 16:55:45.251545	2024-01-09 17:50:46.289254	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	796c70e1-5f27-4433-a415-95e7272effa5	2023-12-31
508ba876-f11c-4971-aa26-9363ee794be9	сформирован	прогноз рассчитан	2024-01-09 17:55:44.248577	2024-01-09 17:56:44.079208	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-23
634d0ac9-8177-43b7-9d25-37a8080083ed	сформирован	прогноз рассчитан	2024-01-09 17:56:52.511561	2024-01-09 18:14:05.297455	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-05
21045f7a-bbbe-40ad-9b57-efd73c505616	сформирован	прогноз рассчитан	2024-01-09 18:16:11.832597	2024-01-09 18:19:24.987791	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-06
3f1f218e-a849-423f-9e6d-55748d7279da	сформирован	прогноз рассчитан	2024-01-09 18:23:30.821998	2024-01-09 18:28:03.120107	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-09
4fe22a66-9f30-41cf-8b02-80ecff31b869	удалён	\N	2024-01-11 21:24:28.795373	\N	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2023-12-30
656295a5-347b-46a5-9b4c-ad2cac9af2a9	удалён	\N	2024-01-19 16:19:17.713366	\N	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-20
76f2dcb8-0af4-4046-ada4-98eb0d6217c4	удалён	\N	2024-01-19 16:40:24.152133	\N	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-17
d168f9d6-167a-4f23-b2eb-88db0e3e0010	удалён	\N	2024-01-19 16:41:15.909165	\N	\N	5f58c307-a3f2-4b13-b888-c80ad08d5ed3	\N	2024-01-20
\.


--
-- TOC entry 3386 (class 0 OID 43715)
-- Dependencies: 214
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: nop
--

COPY public.users (user_id, login, password, role) FROM stdin;
796c70e1-5f27-4433-a415-95e7272effa5	moderator	40bd001563085fc35165329ea1ff5c5ecbdbbeef	2
5f58c307-a3f2-4b13-b888-c80ad08d5ed3	nop	40bd001563085fc35165329ea1ff5c5ecbdbbeef	1
73bba461-bbf0-4238-8263-b60a8828e6cf	test	40bd001563085fc35165329ea1ff5c5ecbdbbeef	1
\.


--
-- TOC entry 3239 (class 2606 OID 43755)
-- Name: connector_apps_types connector_apps_types_pkey; Type: CONSTRAINT; Schema: public; Owner: nop
--

ALTER TABLE ONLY public.connector_apps_types
    ADD CONSTRAINT connector_apps_types_pkey PRIMARY KEY (data_type_id, application_id);


--
-- TOC entry 3235 (class 2606 OID 43730)
-- Name: data_types data_types_pkey; Type: CONSTRAINT; Schema: public; Owner: nop
--

ALTER TABLE ONLY public.data_types
    ADD CONSTRAINT data_types_pkey PRIMARY KEY (data_type_id);


--
-- TOC entry 3237 (class 2606 OID 43736)
-- Name: forecast_applications forecast_applications_pkey; Type: CONSTRAINT; Schema: public; Owner: nop
--

ALTER TABLE ONLY public.forecast_applications
    ADD CONSTRAINT forecast_applications_pkey PRIMARY KEY (application_id);


--
-- TOC entry 3233 (class 2606 OID 43722)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: nop
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- TOC entry 3242 (class 2606 OID 43761)
-- Name: connector_apps_types fk_connector_apps_types_application; Type: FK CONSTRAINT; Schema: public; Owner: nop
--

ALTER TABLE ONLY public.connector_apps_types
    ADD CONSTRAINT fk_connector_apps_types_application FOREIGN KEY (application_id) REFERENCES public.forecast_applications(application_id);


--
-- TOC entry 3243 (class 2606 OID 43756)
-- Name: connector_apps_types fk_connector_apps_types_data_type; Type: FK CONSTRAINT; Schema: public; Owner: nop
--

ALTER TABLE ONLY public.connector_apps_types
    ADD CONSTRAINT fk_connector_apps_types_data_type FOREIGN KEY (data_type_id) REFERENCES public.data_types(data_type_id);


--
-- TOC entry 3240 (class 2606 OID 43742)
-- Name: forecast_applications fk_forecast_applications_creator; Type: FK CONSTRAINT; Schema: public; Owner: nop
--

ALTER TABLE ONLY public.forecast_applications
    ADD CONSTRAINT fk_forecast_applications_creator FOREIGN KEY (creator_id) REFERENCES public.users(user_id);


--
-- TOC entry 3241 (class 2606 OID 43737)
-- Name: forecast_applications fk_forecast_applications_moderator; Type: FK CONSTRAINT; Schema: public; Owner: nop
--

ALTER TABLE ONLY public.forecast_applications
    ADD CONSTRAINT fk_forecast_applications_moderator FOREIGN KEY (moderator_id) REFERENCES public.users(user_id);


-- Completed on 2024-01-23 09:17:17 MSK

--
-- PostgreSQL database dump complete
--

