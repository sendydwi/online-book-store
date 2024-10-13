CREATE TABLE IF NOT EXISTS public.products (
	product_id bigserial NOT NULL,
	isbn text NOT NULL,
	title text NULL,
	subtitle text NULL,
	author text NULL,
	publish_time timestamptz NULL,
	publisher text NULL,
	total_page int8 NULL,
	description text NULL,
	total_stock int8 NULL,
	available_stock int8 NULL,
	is_active bool NULL,
	on_hold_stock int8 NULL,
	price numeric NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	created_by text DEFAULT 'system',
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	updated_by text DEFAULT 'system',
	CONSTRAINT books_pkey PRIMARY KEY (product_id),
	CONSTRAINT uni_books_isbn UNIQUE (isbn)
);


CREATE TABLE IF NOT EXISTS public.users (
	user_id text NOT NULL,
	email text NOT NULL,
	"password" text NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	created_by text DEFAULT 'system',
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	updated_by text DEFAULT 'system',
	CONSTRAINT uni_users_email UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (user_id)
);


CREATE TABLE IF NOT EXISTS public.carts (
	cart_id text NOT NULL,
	user_id text NOT NULL,
	status text NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	created_by text DEFAULT 'system',
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	updated_by text DEFAULT 'system',
	CONSTRAINT carts_pkey PRIMARY KEY (cart_id),
  CONSTRAINT user_cart_fkey FOREIGN KEY(user_id) REFERENCES public.users(user_id)
);

CREATE TABLE IF NOT EXISTS public.cart_items (
	cart_item_id bigserial NOT NULL,
	cart_id text NOT NULL,
	product_id int8 NOT NULL,
	quantity int8 NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	created_by text DEFAULT 'system',
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	updated_by text DEFAULT 'system',
	CONSTRAINT cart_items_pkey PRIMARY KEY (cart_item_id),
  CONSTRAINT cart_cart_items_fkey FOREIGN KEY(cart_id) REFERENCES public.carts(cart_id),
  CONSTRAINT product_cart_item_fkey FOREIGN KEY(product_id) REFERENCES public.products(product_id)
);

CREATE TABLE IF NOT EXISTS public.orders (
	order_id text NOT NULL,
	user_id text NOT NULL,
	status text NOT NULL,
	total_price numeric NOT NULL,
	delivery_address text NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	created_by text DEFAULT 'system',
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	updated_by text DEFAULT 'system',
	CONSTRAINT orders_pkey PRIMARY KEY (order_id),
  CONSTRAINT user_orders_fkey FOREIGN KEY(user_id) REFERENCES public.users(user_id)
);

CREATE TABLE IF NOT EXISTS public.order_items (
	order_item_id bigserial NOT NULL,
	order_id text NOT NULL,
	product_id int8 NOT NULL,
	subtotal_price numeric NOT NULL,
	quantity int8 NOT NULL,
	product_snapshot jsonb NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	created_by text DEFAULT 'system',
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP,
	updated_by text DEFAULT 'system',
	CONSTRAINT order_items_pkey PRIMARY KEY (order_item_id),
  CONSTRAINT order_order_items_fkey FOREIGN KEY(order_id) REFERENCES public.orders(order_id),
  CONSTRAINT product_order_items_fkey FOREIGN KEY(product_id) REFERENCES public.products(product_id)
);

CREATE INDEX IF NOT EXISTS order_user_id_idx ON public.orders(user_id);
CREATE INDEX IF NOT EXISTS order_items_order_id_idx ON public.order_items(order_id);

CREATE INDEX IF NOT EXISTS cart_user_id_status_idx ON public.carts(user_id, status);
CREATE INDEX IF NOT EXISTS cart_items_cart_id_idx ON public.cart_items(cart_id);

CREATE INDEX IF NOT EXISTS users_email_idx ON public.users(email);

COPY public.products
FROM '/docker-entrypoint-initdb.d/products.csv'
DELIMITER ','
CSV HEADER;