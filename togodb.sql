--
-- PostgreSQL database dump
--

-- Dumped from database version 12.4 (Debian 12.4-1.pgdg100+1)
-- Dumped by pg_dump version 12.4 (Debian 12.4-1.pgdg100+1)

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
-- Name: tasks; Type: TABLE; Schema: public; Owner: togoapp
--

CREATE TABLE public.tasks (
    id text NOT NULL,
    content text NOT NULL,
    user_id text NOT NULL,
    created_date timestamp without time zone NOT NULL
);


ALTER TABLE public.tasks OWNER TO togoapp;

--
-- Name: users; Type: TABLE; Schema: public; Owner: togoapp
--

CREATE TABLE public.users (
    id text NOT NULL,
    password text NOT NULL,
    max_todo integer DEFAULT 5 NOT NULL
);


ALTER TABLE public.users OWNER TO togoapp;

--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: togoapp
--

COPY public.tasks (id, content, user_id, created_date) FROM stdin;
e1da0b9b-7ecc-44f9-82ff-4623cc50446a	first content	firstUser	2020-06-29 00:00:00
055261ab-8ba8-49e1-a9e8-e9f725ba9104	second content	firstUser	2020-06-29 00:00:00
2bf3d510-c0fb-41e9-ad12-4b9a60b37e7a	another content	firstUser	2020-06-29 00:00:00
e35e13f8-35f3-409f-8e2f-f3e0173fcca3	sadsa	firstUser	2020-08-10 00:00:00
2a73a4d5-dd05-4c77-bcbd-f5e51a6d6809	sadsad	firstUser	2020-08-11 00:00:00
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: togoapp
--

COPY public.users (id, password, max_todo) FROM stdin;
firstUser	example	5
\.


--
-- Name: tasks tasks_pk; Type: CONSTRAINT; Schema: public; Owner: togoapp
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pk PRIMARY KEY (id);


--
-- Name: users user_pk; Type: CONSTRAINT; Schema: public; Owner: togoapp
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT user_pk PRIMARY KEY (id);


--
-- Name: tasks tasks_fk; Type: FK CONSTRAINT; Schema: public; Owner: togoapp
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_fk FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

