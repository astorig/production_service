BEGIN;

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = public, extensions;
SET default_tablespace = '';
SET default_with_oids = FALSE;

-- EXTENSIONS --
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- TABLES --

CREATE TABLE public.currency
(
    id SERIAL PRIMARY KEY,
    name TEXT,
    symbol TEXT
);

CREATE TABLE public.category
(
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE public.product
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    image_id UUID,
    price BIGINT,
    currency_id INT,
    rating INT,
    category_id INT NOT NULL,
    specification JSONB,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
);

--  DATA  --

INSERT INTO public.currency (name, symbol)
VALUES ('RUB', '₽');

    INSERT INTO public.currency (name, symbol)
VALUES ('USD', '$');

COMMIT;