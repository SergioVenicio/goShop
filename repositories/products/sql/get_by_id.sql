SELECT
    id,
    name,
    description,
    price,
    amount
FROM product
WHERE id = %d LIMIT 1