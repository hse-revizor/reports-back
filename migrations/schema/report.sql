CREATE TABLE IF NOT EXISTS public.reports (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	title text NULL,
	description text NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);