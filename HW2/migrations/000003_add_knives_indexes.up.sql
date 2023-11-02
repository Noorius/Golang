CREATE INDEX IF NOT EXISTS knives_title_idx ON knives USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS knives_material_idx ON knives USING GIN (to_tsvector('simple', material));
CREATE INDEX IF NOT EXISTS knives_color_idx ON knives USING GIN (to_tsvector('simple', color));
CREATE INDEX IF NOT EXISTS knives_country_idx ON knives USING GIN (to_tsvector('simple', country));