package repository

import (
	"database/sql"

	"github.com/GbSouza15/sistemaEstoque/internal/app/models"
	"github.com/google/uuid"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (p *ProductRepository) CreateProduct(id uuid.UUID, name string, price float64, companyId, segment_id, deposit_id uuid.UUID, quantity int) error {
	_, err := p.Db.Exec("INSERT INTO products (id, name, price, company_id, segment_id) VALUES (?, ?, ?, ?, ?)", id, name, price, companyId, segment_id)
	if err != nil {
		return err
	}

	_, err = p.Db.Exec("INSERT INTO deposit_products (id, product_id, deposit_id, quantity) VALUES (?, ?, ?, ?)", uuid.New(), id, deposit_id, quantity)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) RemoveProduct(id string) error {
	_, err := p.Db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) UpdateProduct(id string, name *string, price *float64, stock *int, segmentId *string) error {
	query := `
		UPDATE
			products
		SET
			name = COALESCE(?, name),
			price = COALESCE(?, price),
			segment_id = COALESCE(?, segment_id)
		WHERE
			id = ?
	`

	_, err := p.Db.Exec(query, name, price, segmentId, id)
	if err != nil {
		return err
	}

	if stock != nil {
		_, err = p.Db.Exec("UPDATE deposit_products SET quantity = ? WHERE product_id = ?", stock, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *ProductRepository) AddProductSegment(id uuid.UUID, name string, company_id uuid.UUID) error {
	_, err := p.Db.Exec("INSERT INTO product_segments (id, name, company_id) VALUES (?, ?, ?)", id, name, company_id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) RemoveProductSegment(id uuid.UUID) error {
	_, err := p.Db.Exec("DELETE FROM product_segments WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductRepository) SearchProduct(name string) ([]models.ProductInfo, error) {
	query := `
        SELECT
            p.id,
            p.name,
            p.price,
            p.company_id,
            p.segment_id,
            ps.name AS segment_name,
            COALESCE(SUM(dp.quantity), 0) AS stock_quantity
        FROM
            products p
        LEFT JOIN
            product_segments ps ON p.segment_id = ps.id
        LEFT JOIN
            deposit_products dp ON p.id = dp.product_id
        WHERE p.name LIKE CONCAT('%', ?, '%')
        GROUP BY
            p.id, p.name, p.price, p.company_id, p.segment_id, ps.name
    `

	rows, err := p.Db.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productInfos []models.ProductInfo
	for rows.Next() {
		var productInfo models.ProductInfo
		if err := rows.Scan(&productInfo.Id, &productInfo.Name, &productInfo.Price, &productInfo.CompanyId, &productInfo.SegmentId, &productInfo.SegmentName, &productInfo.StockQuantity); err != nil {
			return nil, err
		}

		productInfos = append(productInfos, productInfo)
	}

	return productInfos, nil
}
