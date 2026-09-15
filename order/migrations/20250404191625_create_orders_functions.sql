-- +goose Up

-- Функция добавления заказа
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION add_order(
    p_order_uuid UUID,
    p_user_uuid UUID,
    p_parts_uuid UUID[],
    p_total_price REAL,
    p_status TEXT,
    p_transaction_uuid UUID DEFAULT NULL,
    p_payment_method TEXT DEFAULT NULL
) RETURNS VOID AS
$$
BEGIN
INSERT INTO orders (
    order_uuid, user_uuid, parts_uuid, total_price, transaction_uuid, payment_method, status
) VALUES (
             p_order_uuid, p_user_uuid, p_parts_uuid, p_total_price, p_transaction_uuid, p_payment_method, p_status
         );
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- Функция обновления заказа
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION update_order(
    p_order_uuid UUID,
    p_user_uuid UUID DEFAULT NULL,
    p_parts_uuid UUID[] DEFAULT NULL,
    p_total_price REAL DEFAULT NULL,
    p_status TEXT DEFAULT NULL,
    p_transaction_uuid UUID DEFAULT NULL,
    p_payment_method TEXT DEFAULT NULL
) RETURNS VOID AS
$$
DECLARE
    sql_query TEXT := 'UPDATE orders SET ';
    set_clauses TEXT := '';
    first BOOLEAN := TRUE;
BEGIN
    -- Проверка на существование заказа
    IF NOT EXISTS (SELECT 1 FROM orders WHERE order_uuid = p_order_uuid) THEN
        RAISE EXCEPTION 'Order with UUID % does not exist', p_order_uuid;
    END IF;

    -- Формируем части запроса динамически
    IF p_user_uuid IS NOT NULL THEN
        set_clauses := set_clauses || 'user_uuid = ' || quote_literal(p_user_uuid);
        first := FALSE;
    END IF;

    IF p_parts_uuid IS NOT NULL THEN
        IF NOT first THEN set_clauses := set_clauses || ', '; END IF;
        set_clauses := set_clauses || 'parts_uuid = ' || quote_literal(p_parts_uuid);
        first := FALSE;
    END IF;

    IF p_total_price IS NOT NULL THEN
        IF NOT first THEN set_clauses := set_clauses || ', '; END IF;
        set_clauses := set_clauses || 'total_price = ' || quote_literal(p_total_price);
        first := FALSE;
    END IF;

    IF p_status IS NOT NULL THEN
        IF NOT first THEN set_clauses := set_clauses || ', '; END IF;
        set_clauses := set_clauses || 'status = ' || quote_literal(p_status);
        first := FALSE;
    END IF;

    IF p_transaction_uuid IS NOT NULL THEN
        IF NOT first THEN set_clauses := set_clauses || ', '; END IF;
        set_clauses := set_clauses || 'transaction_uuid = ' || quote_literal(p_transaction_uuid);
        first := FALSE;
    END IF;

    IF p_payment_method IS NOT NULL THEN
        IF NOT first THEN set_clauses := set_clauses || ', '; END IF;
        set_clauses := set_clauses || 'payment_method = ' || quote_literal(p_payment_method);
    END IF;

    -- Если нет ни одного значения для обновления
    IF set_clauses = '' THEN
        RAISE NOTICE 'No fields to update for order UUID %', p_order_uuid;
        RETURN;
    END IF;

    -- Собираем и выполняем финальный SQL-запрос
    sql_query := sql_query || set_clauses || ' WHERE order_uuid = ' || quote_literal(p_order_uuid);
    EXECUTE sql_query;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- Функция удаления заказа
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION delete_order(
    p_order_uuid UUID
) RETURNS VOID AS
$$
BEGIN
DELETE FROM orders WHERE order_uuid = p_order_uuid;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

-- Функция получения заказа
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION get_order(
    p_order_uuid UUID
) RETURNS TABLE (
    order_uuid UUID,
    user_uuid UUID,
    parts_uuid UUID[],
    total_price REAL,
    transaction_uuid UUID,
    payment_method TEXT,
    status TEXT
    ) AS
$$
BEGIN
    RETURN QUERY
        SELECT o.order_uuid, o.user_uuid, o.parts_uuid, o.total_price, o.transaction_uuid, o.payment_method, o.status
        FROM orders o
        WHERE o.order_uuid = p_order_uuid;
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down

-- Удаляем функции в обратном порядке

-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_order(UUID);
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS delete_order(UUID);
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS update_order(UUID, UUID, UUID[], REAL, TEXT, UUID, UUID);
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS add_order(UUID, UUID, UUID[], REAL, TEXT, UUID, UUID);
-- +goose StatementEnd
