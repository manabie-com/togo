--
-- PostgreSQL database dump
--

-- Dumped from database version 14.0 (Debian 14.0-1.pgdg110+1)
-- Dumped by pg_dump version 14.0 (Debian 14.0-1.pgdg110+1)

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
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: togo_user
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO togo_user;

--
-- Name: tasks; Type: TABLE; Schema: public; Owner: togo_user
--

CREATE TABLE public.tasks (
    id character varying(50) NOT NULL,
    content text NOT NULL,
    user_id bigint NOT NULL,
    created_date date NOT NULL
);


ALTER TABLE public.tasks OWNER TO togo_user;

--
-- Name: users; Type: TABLE; Schema: public; Owner: togo_user
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(256) NOT NULL,
    max_todo integer DEFAULT 5 NOT NULL
);


ALTER TABLE public.users OWNER TO togo_user;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: togo_user
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO togo_user;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: togo_user
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: togo_user
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: togo_user
--

COPY public.schema_migrations (version, dirty) FROM stdin;
1	t
\.


--
-- Data for Name: tasks; Type: TABLE DATA; Schema: public; Owner: togo_user
--

COPY public.tasks (id, content, user_id, created_date) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: togo_user
--

COPY public.users (id, username, password, max_todo) FROM stdin;
1	nohattee	1qaz@WSX	5
\.


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: togo_user
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- PostgreSQL database dump complete
--

