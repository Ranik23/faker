-- Таблица почтовых отделений
CREATE TABLE post_office (
  id                SERIAL PRIMARY KEY NOT NULL,
  name              VARCHAR(255)      NOT NULL,
  phone             VARCHAR(20)
);

-- Таблица клиентов
CREATE TABLE customer (
  id                SERIAL PRIMARY KEY NOT NULL,
  name              VARCHAR(255)      NOT NULL,
  contact_info      VARCHAR(255),
  reg_date          DATE              NOT NULL DEFAULT CURRENT_DATE
);

-- Таблица адресов
CREATE TABLE address (
  id                SERIAL PRIMARY KEY NOT NULL,
  country           VARCHAR(100)      NOT NULL,
  city              VARCHAR(100)      NOT NULL,
  street            VARCHAR(255)      NOT NULL,
  client_id         INTEGER           NOT NULL,
  postal_code       VARCHAR(20)       NOT NULL,
  CONSTRAINT fk_address_customer FOREIGN KEY (client_id) REFERENCES customer(id) ON DELETE CASCADE
);

-- Таблица почтовых отправлений
CREATE TABLE postal_item (
  id                SERIAL PRIMARY KEY NOT NULL,
  track_num         VARCHAR(50)       NOT NULL UNIQUE,
  type              VARCHAR(100)      NOT NULL,
  weight            NUMERIC(10, 2)    NOT NULL CHECK (weight > 0),
  status            VARCHAR(50),
  disp_date         DATE              NOT NULL,
  arrival_date      DATE,
  sender_id         INTEGER,
  recipient_id      INTEGER,
  post_office_id    INTEGER           NOT NULL,
  CONSTRAINT fk_postal_item_post_office FOREIGN KEY (post_office_id) REFERENCES post_office(id) ON DELETE CASCADE,
  CONSTRAINT fk_postal_item_sender FOREIGN KEY (sender_id) REFERENCES customer(id) ON DELETE CASCADE,
  CONSTRAINT fk_postal_item_recipient FOREIGN KEY (recipient_id) REFERENCES customer(id) ON DELETE CASCADE
);

-- Таблица платежей
CREATE TABLE payment (
  id                SERIAL PRIMARY KEY NOT NULL,
  amount            NUMERIC(10, 2)    NOT NULL CHECK (amount > 0),
  payment_method    VARCHAR(100),
  payment_date      DATE              NOT NULL DEFAULT CURRENT_DATE,
  postal_item_id    INTEGER           NOT NULL,
  customer_id       INTEGER           NOT NULL,
  CONSTRAINT fk_payment_postal_item FOREIGN KEY (postal_item_id) REFERENCES postal_item(id) ON DELETE CASCADE,
  CONSTRAINT fk_payment_customer FOREIGN KEY (customer_id) REFERENCES customer(id) ON DELETE CASCADE
);

-- Таблица статусов отправлений
CREATE TABLE postal_status (
  id                SERIAL PRIMARY KEY NOT NULL,
  status_name       VARCHAR(255)      NOT NULL,
  description       TEXT
);

-- Таблица сотрудников
CREATE TABLE employee (
  id                SERIAL PRIMARY KEY NOT NULL,
  name              VARCHAR(255)      NOT NULL,
  position          VARCHAR(100)      NOT NULL,
  post_office_id    INTEGER           NOT NULL,
  hire_date         DATE              NOT NULL,
  CONSTRAINT fk_employee_post_office FOREIGN KEY (post_office_id) REFERENCES post_office(id) ON DELETE CASCADE
);

-- Таблица истории изменения статусов
CREATE TABLE status_transaction (
  id                SERIAL PRIMARY KEY NOT NULL,
  postal_item_id    INTEGER           NOT NULL,
  status_id         INTEGER           NOT NULL,
  status_date       DATE              NOT NULL DEFAULT CURRENT_DATE,
  description       TEXT,
  employee_id       INTEGER,
  CONSTRAINT fk_status_transaction_postal_item FOREIGN KEY (postal_item_id) REFERENCES postal_item(id) ON DELETE CASCADE,
  CONSTRAINT fk_status_transaction_status FOREIGN KEY (status_id) REFERENCES postal_status(id) ON DELETE CASCADE,
  CONSTRAINT fk_status_transaction_employee FOREIGN KEY (employee_id) REFERENCES employee(id) ON DELETE SET NULL
);
