SET check_function_bodies = false;
CREATE EXTENSION pgcrypto;
CREATE EXTENSION pg_trgm;
CREATE TABLE public.events (
    title text NOT NULL,
    description text,
    images jsonb DEFAULT jsonb_build_array() NOT NULL,
    tags jsonb DEFAULT jsonb_build_array() NOT NULL,
    date timestamp with time zone NOT NULL,
    price numeric DEFAULT 0 NOT NULL,
    location jsonb,
    city_id uuid,
    specific_address text,
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    user_id uuid NOT NULL,
    CONSTRAINT no_negative_price CHECK ((price >= (0)::numeric)),
    CONSTRAINT no_past_date CHECK ((date > now()))
);
CREATE FUNCTION public.event_bookmarked_by_user(event_row public.events, hasura_session json) RETURNS uuid
    LANGUAGE sql STABLE
    AS $$
    SELECT id
    FROM bookmarks B
    WHERE B.user_id::text = hasura_session ->> 'x-hasura-user-id' AND B.event_id = event_row.id
$$;
CREATE FUNCTION public.event_bookmarks_count(event_row public.events) RETURNS bigint
    LANGUAGE sql STABLE
    AS $$
SELECT COUNT(id) FROM bookmarks WHERE event_id = event_row.id
$$;
CREATE FUNCTION public.event_distance(events_row public.events, lat numeric, long numeric) RETURNS numeric
    LANGUAGE sql STABLE
    AS $$
    SELECT SQRT(POWER((events_row.location ->> 0)::NUMERIC - lat, 2) + POWER((events_row.location ->> 1)::NUMERIC - long,2))
$$;
CREATE FUNCTION public.events_by_location(lat numeric, long numeric) RETURNS SETOF public.events
    LANGUAGE sql STABLE
    AS $$
    SELECT *
    from events
    ORDER BY
    SQRT(POWER((location ->> 0)::NUMERIC - lat, 2) + POWER((location ->> 1)::NUMERIC - long,2))
$$;
CREATE TABLE public.users (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    email text NOT NULL,
    name text NOT NULL,
    description text,
    avatar_url text,
    password_hash text NOT NULL,
    member_since timestamp with time zone DEFAULT now() NOT NULL
);
CREATE FUNCTION public.me(hasura_session json) RETURNS public.users
    LANGUAGE sql STABLE
    AS $$
SELECT
  *
FROM
  users
WHERE
  id :: text = hasura_session ->> 'x-hasura-user-id'
LIMIT 1;
$$;
CREATE FUNCTION public.owner_followed_by_user(event_row public.events, hasura_session json) RETURNS uuid
    LANGUAGE sql STABLE
    AS $$
    SELECT id
    FROM follows F
    WHERE F.follower_id::text = hasura_session ->> 'x-hasura-user-id' AND F.followed_id = event_row.user_id
$$;
CREATE FUNCTION public.search_events(search text) RETURNS SETOF public.events
    LANGUAGE sql STABLE
    AS $$
    SELECT *
    FROM events
    WHERE
      search <% (title || ' ' || description || ' ' || tags)
    ORDER BY
      similarity(search, (title || ' ' || description || ' ' || tags)) DESC;
$$;
CREATE FUNCTION public.user_followers_count(users_row public.users) RETURNS bigint
    LANGUAGE sql STABLE
    AS $$
SELECT COUNT(id) FROM follows WHERE followed_id = users_row.id
$$;
CREATE FUNCTION public.user_following_count(users_row public.users) RETURNS bigint
    LANGUAGE sql STABLE
    AS $$
SELECT COUNT(id) FROM follows WHERE follower_id = users_row.id
$$;
CREATE FUNCTION public.users_by_email(useremail text) RETURNS public.users
    LANGUAGE sql STABLE
    AS $$
    SELECT *
    FROM users
    WHERE
      email = userEmail
    LIMIT 1
$$;
CREATE TABLE public.bookmarks (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    event_id uuid NOT NULL,
    user_id uuid NOT NULL
);
CREATE TABLE public.cities (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    name text NOT NULL
);
CREATE TABLE public.follows (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    follower_id uuid NOT NULL,
    followed_id uuid NOT NULL
);
CREATE TABLE public.session_refresh_tokens (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    token text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    user_id uuid NOT NULL
);
CREATE TABLE public.tickets (
    id uuid DEFAULT public.gen_random_uuid() NOT NULL,
    event_id uuid,
    user_id uuid,
    is_valid boolean DEFAULT true NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);
ALTER TABLE ONLY public.bookmarks
    ADD CONSTRAINT bookmarks_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_city_name_key UNIQUE (name);
ALTER TABLE ONLY public.cities
    ADD CONSTRAINT cities_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.follows
    ADD CONSTRAINT follows_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.session_refresh_tokens
    ADD CONSTRAINT session_refresh_tokens_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.tickets
    ADD CONSTRAINT tickets_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.bookmarks
    ADD CONSTRAINT unique_bookmark UNIQUE (event_id, user_id);
ALTER TABLE ONLY public.follows
    ADD CONSTRAINT unique_follows UNIQUE (follower_id, followed_id);
ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);
ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
CREATE INDEX events_gin_idx ON public.events USING gin ((((((title || ' '::text) || description) || ' '::text) || tags)) public.gin_trgm_ops);
CREATE INDEX token_key ON public.session_refresh_tokens USING btree (token);
ALTER TABLE ONLY public.bookmarks
    ADD CONSTRAINT bookmarks_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY public.bookmarks
    ADD CONSTRAINT bookmarks_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_city_id_fkey FOREIGN KEY (city_id) REFERENCES public.cities(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY public.follows
    ADD CONSTRAINT follows_followed_id_fkey FOREIGN KEY (followed_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY public.follows
    ADD CONSTRAINT follows_follower_id_fkey FOREIGN KEY (follower_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY public.session_refresh_tokens
    ADD CONSTRAINT session_refresh_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY public.tickets
    ADD CONSTRAINT tickets_event_id_fkey FOREIGN KEY (event_id) REFERENCES public.events(id) ON UPDATE CASCADE ON DELETE SET NULL;
ALTER TABLE ONLY public.tickets
    ADD CONSTRAINT tickets_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE SET NULL;
