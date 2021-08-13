UPDATE
    product
SET
    name = $2
    , description = $3
    , price = $4
    , amount = $5
WHERE id = $1