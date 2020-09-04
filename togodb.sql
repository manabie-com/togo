--
-- PostgreSQL database dump
--

-- Dumped from database version 12.4 (Debian 12.4-1.pgdg100+1)
-- Dumped by pg_dump version 12.4 (Ubuntu 12.4-0ubuntu0.20.04.1)

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

