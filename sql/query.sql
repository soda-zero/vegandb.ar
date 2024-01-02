-- name: ListProducts :many
SELECT * FROM products
ORDER BY name;

-- name: GetCategoryTree :many
WITH RECURSIVE cte_categories (id, name, parent_category_id) AS (
    SELECT c.id, c.name, c.parent_category_id
    FROM categories c
    JOIN products p ON c.id = p.category_id
    WHERE p.id = ?

    UNION ALL

    SELECT c.id, c.name, c.parent_category_id
    FROM categories c
    JOIN cte_categories ct ON c.id = ct.parent_category_id
)

SELECT * FROM cte_categories;


-- name: FilterProducts :many
SELECT * FROM products
WHERE name LIKE '%' || ? || '%' COLLATE NOCASE;
