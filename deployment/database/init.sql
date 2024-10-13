CREATE TABLE IF NOT EXISTS public.products (
	id bigserial NOT NULL,
	isbn text NULL,
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
	created_at timestamptz NULL,
	created_by text NULL,
	updated_at timestamptz NULL,
	updated_by text NULL,
	CONSTRAINT books_pkey PRIMARY KEY (id),
	CONSTRAINT uni_books_isbn UNIQUE (isbn)
);

CREATE TABLE IF NOT EXISTS public.cart_items (
	cart_item_id bigserial NOT NULL,
	cart_id text NULL,
	product_id int8 NULL,
	quantity int8 NULL,
	created_at timestamptz NULL,
	created_by text NULL,
	updated_at timestamptz NULL,
	updated_by text NULL,
	CONSTRAINT cart_items_pkey PRIMARY KEY (cart_item_id)
  CONSTRAINT cart_cart_items_fkey FOREIGN KEY(cart_id) REFERENCES public.carts(cart_id)
  CONSTRAINT product_cart_item_fkey FOREIGN KEY(product_id) REFERENCES public.products(id)
);


CREATE TABLE IF NOT EXISTS public.carts (
	cart_id text NOT NULL,
	user_id text NULL,
	status text NULL,
	created_at timestamptz NULL,
	created_by text NULL,
	updated_at timestamptz NULL,
	updated_by text NULL,
	CONSTRAINT carts_pkey PRIMARY KEY (cart_id)
  CONSTRAINT user_cart_fkey FOREIGN KEY(user_id) REFERENCES public.users(user_id)
);


CREATE TABLE IF NOT EXISTS public.orders (
	id text NOT NULL,
	user_id text NULL,
	status text NULL,
	total_price numeric NULL,
	delivery_address text NULL,
	created_at timestamptz NULL,
	created_by text NULL,
	updated_at timestamptz NULL,
	updated_by text NULL,
	CONSTRAINT orders_pkey PRIMARY KEY (id)
  CONSTRAINT user_orders_fkey FOREIGN KEY(user_id) REFERENCES public.users(user_id)
);

CREATE TABLE IF NOT EXISTS public.order_items (
	id bigserial NOT NULL,
	order_id text NULL,
	product_id int8 NULL,
	subtotal_price numeric NULL,
	quantity int8 NULL,
	product_snapshot jsonb NULL,
	created_at timestamptz NULL,
	created_by text NULL,
	updated_at timestamptz NULL,
	updated_by text NULL,
	CONSTRAINT order_items_pkey PRIMARY KEY (id)
  CONSTRAINT order_order_items_fkey FOREIGN KEY(order_id) REFERENCES public.orders(order_id)
  CONSTRAINT product_order_items_fkey FOREIGN KEY(product_id) REFERENCES public.products(product_id)
);

CREATE TABLE IF NOT EXISTS public.users (
	user_id text NOT NULL,
	email text NULL,
	"password" text NULL,
	created_at timestamptz NULL,
	created_by text NULL,
	updated_at timestamptz NULL,
	updated_by text NULL,
	CONSTRAINT uni_users_email UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (user_id)
);

CREATE INDEX IF NOT EXISTS order_user_id_idx ON public.orders(user_id);
CREATE INDEX IF NOT EXISTS order_items_order_id_idx ON public.order_items(order_id);

CREATE INDEX IF NOT EXISTS cart_user_id_status_idx ON public.carts(user_id, status);
CREATE INDEX IF NOT EXISTS cart_items_cart_id_idx ON public.cart_items(cart_id);

CREATE INDEX IF NOT EXISTS users_email_idx ON public.users(email);

