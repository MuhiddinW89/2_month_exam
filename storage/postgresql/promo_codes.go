package postgresql

import (
	"app/api/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type promoRepo struct {
	db *pgxpool.Pool
}

func NewPromoRepo(db *pgxpool.Pool) *promoRepo {
	return &promoRepo{
		db: db,
	}
}

func (r *promoRepo) Create(ctx context.Context, req *models.CreatePromo) (int, error) {
	var (
		query string
		id int
	)

	query = `
		INSERT INTO promo_codes(
			promo_name, 
			promo_discount,
			promo_discount_type,
			promo_order_limit_price 
		)
		VALUES (
			$1, $2, $3, $4
		)
		RETURNING promo_id
	`
	err := r.db.QueryRow(ctx, query,
		req.Promo_name,
		req.Promo_discount,
		req.Promo_discount_type,
		req.Promo_order_limit_price,
	).Scan(&id)
	
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *promoRepo) GetByID(ctx context.Context, req *models.PromoPrimaryKey) (*models.Promo_code, error) {
	
	var(
		query string
		promoCode models.Promo_code
	)
	
	query = `
		SELECT 
		
		promo_id,
		promo_name,
		promo_discount,
		promo_discount_type,
		promo_order_limit_price

		FROM promo_codes
		WHERE promo_id = $1
	`
	err := r.db.QueryRow(ctx, query, req.Promo_id).Scan(
		&promoCode.Promo_id,
		&promoCode.Promo_name,
		&promoCode.Promo_discount,
		&promoCode.Promo_discount_type,
		&promoCode.Promo_order_limit_price,
	)

	if err != nil {
		return nil, err
	}
	
	return &promoCode, nil
}

func(r *promoRepo) GetList(ctx context.Context, req *models.GetListPromoRequest) (res *models.GetListPromoResponse, err error) {

	res = &models.GetListPromoResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT 
		
		COUNT(*) OVER(),
		promo_id,
		promo_name,
		promo_discount,
		promo_discount_type,
		promo_order_limit_price

		FROM promo_codes
	`


	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var promoCodes models.Promo_code
		err = rows.Scan(
			&res.Count,
			&promoCodes.Promo_id,
			&promoCodes.Promo_name,
			&promoCodes.Promo_discount,
			&promoCodes.Promo_discount_type,
			&promoCodes.Promo_order_limit_price,	
		)

		if err != nil {
			return nil, err
		}

		res.Promo_codes = append(res.Promo_codes, &promoCodes)
	}
	return res, nil	
}

func(r *promoRepo) Delete(ctx context.Context, req *models.PromoPrimaryKey) (int64, error){
	query := `
		DELETE 
		FROM promo_codes
		WHERE promo_id = $1
	`

	result, err := r.db.Exec(ctx, query, req.Promo_id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}