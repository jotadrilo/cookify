--
-- Name: products; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE IF NOT EXISTS products
(
  id             SERIAL PRIMARY KEY,
  uuid           character varying NOT NULL,
  name           character varying NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_products_uuid ON products (uuid);

--
-- Name: units; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE IF NOT EXISTS units
(
    id            SERIAL PRIMARY KEY,
    uuid          character varying NOT NULL,
    name          character varying NOT NULL,
    unit_10em3    character varying,
    unit_1        character varying NOT NULL,
    unit_10e3     character varying
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_units_uuid ON units (uuid);
