-- CREATE TABLE IF NOT EXISTS public.inventory
-- (
--   id SERIAL PRIMARY KEY,
--   code TEXT,
--   value TEXT
-- );

-- insert into public.inventory (code, value) values ('X001', 'pepsi'), ('X002', 'coke'), ('X003', 'ace');

CREATE TABLE IF NOT EXISTS public.data
(
  id SERIAL PRIMARY KEY,
  value INT NOT NULL,
  key INT NOT NULL UNIQUE
);
CREATE INDEX key_idx ON public.data (key);
ALTER TABLE public.data REPLICA IDENTITY FULL;