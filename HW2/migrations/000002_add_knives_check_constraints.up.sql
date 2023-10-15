ALTER TABLE knives ADD CONSTRAINT knives_duration_check CHECK (duration >= 0);
