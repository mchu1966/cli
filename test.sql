CREATE TABLE IF NOT EXISTS public.users
(
    id bigint NOT NULL DEFAULT 'nextval('users_id_seq'::regclass)',
    username character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    creation_time timestamp with time zone DEFAULT 'now()',
    modification_time timestamp with time zone DEFAULT 'now()',
    last_login timestamp with time zone,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_username_key UNIQUE (username)
)