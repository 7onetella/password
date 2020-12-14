--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5 (Raspbian 11.5-1+deb10u1)
-- Dumped by pg_dump version 11.5 (Raspbian 11.5-1+deb10u1)

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
SET default_with_oids = false;

--
-- Name: passwords; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.passwords (
    id character varying(36) NOT NULL,
    title character varying(256) NOT NULL,
    url character varying(2048),
    username character varying(64) NOT NULL,
    password character varying(1024) NOT NULL,
    notes character varying(2048),
    tags character varying(256),
    admin_id character varying(36) NOT NULL
);


--
-- Data for Name: passwords; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.passwords (id, title, url, username, password, notes, tags, admin_id) FROM stdin;
23b85391-1b55-41ac-95ab-cc51ad1c5c18	MYBANK	https://www.example.com	foo	bar114$			user1
\.


--
-- Name: passwords pk_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.passwords
    ADD CONSTRAINT pk_id PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

