DO $$
DECLARE r RECORD;
BEGIN -- if the schema you operate on is not "current", you will want to
-- replace current_schema() in query with 'schematodeletetablesfrom'
-- *and* update the generate 'DROP...' accordingly.
FOR r IN (
  SELECT tablename
  FROM pg_tables
  WHERE schemaname = current_schema()
) LOOP EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
END LOOP;
END $$;
-------------------------------------------------------
-- Table: public.loadbalancers
-------------------------------------------------------
CREATE TABLE public.loadbalancers (
  id varchar,
  data jsonb,
  load_balancer_ip varchar,
  load_balancer jsonb,
  last_modified timestamptz,
  source varchar,
  md5hash text,
  last_error varchar,
  last_modified_by varchar,
  CONSTRAINT loadbalancers_pkey PRIMARY KEY (id)
) WITH (OIDS = FALSE) TABLESPACE pg_default;
ALTER TABLE public.loadbalancers OWNER to postgres;
-------------------------------------------------------
-- Table: public.virtualservers
-------------------------------------------------------
CREATE TABLE public.virtualservers (
  id varchar,
  data jsonb,
  load_balancer_ip varchar,
  load_balancer jsonb,
  last_modified timestamptz,
  source varchar,
  md5hash text,
  last_error varchar,
  last_modified_by varchar,
  CONSTRAINT virtualservers_pkey PRIMARY KEY (id)
) WITH (OIDS = FALSE) TABLESPACE pg_default;
ALTER TABLE public.virtualservers OWNER to postgres;
-------------------------------------------------------
-- Table: public.migrate
-------------------------------------------------------
CREATE TABLE public.migrate (
  id varchar,
  data jsonb,
  load_balancer_ip varchar,
  load_balancer jsonb,
  last_modified timestamptz,
  source varchar,
  md5hash text,
  last_error varchar,
  last_modified_by varchar,
  CONSTRAINT migrate_pkey PRIMARY KEY (id)
) WITH (OIDS = FALSE) TABLESPACE pg_default;
ALTER TABLE public.migrate OWNER to postgres;
-------------------------------------------------------
-- Table: public.recycle
-------------------------------------------------------
CREATE TABLE public.recycle (
  id varchar,
  data jsonb,
  load_balancer_ip varchar,
  load_balancer jsonb,
  last_modified timestamptz,
  source varchar,
  md5hash text,
  last_error varchar,
  last_modified_by varchar,
  CONSTRAINT recycle_pkey PRIMARY KEY (id)
) WITH (OIDS = FALSE) TABLESPACE pg_default;
ALTER TABLE public.recycle OWNER to postgres;
-------------------------------------------------------
-- Table: public.status
-------------------------------------------------------
CREATE TABLE public.status (
  id varchar,
  status_id integer,
  data jsonb,
  load_balancer_ip varchar,
  load_balancer jsonb,
  last_modified timestamptz,
  source varchar,
  md5hash text,
  last_error varchar,
  last_modified_by varchar,
  CONSTRAINT status_pkey PRIMARY KEY (id)
) WITH (OIDS = FALSE) TABLESPACE pg_default;
ALTER TABLE public.status OWNER to postgres;
-------------------------------------------------------
-- Table: public.status_description
-------------------------------------------------------
CREATE TABLE public.statusdescription (
  id integer PRIMARY KEY,
  short varchar,
  long text
) WITH (OIDS = FALSE) TABLESPACE pg_default;
ALTER TABLE public.statusdescription OWNER to postgres;
-------------------------------------------------------
-- Populate: public.status_description
-------------------------------------------------------
INSERT INTO public.statusdescription (id,short) VALUES (0,'deployed'),(1,'fail'),(2,'partial'),(3,'migrating'),(4,'migrated'),(5,'creating'),(6,'updating'),(7,'deleting');