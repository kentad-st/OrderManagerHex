package postgresDao

import (
	model "OrderManagerHex/domain/model"

	repository "OrderManagerHex/domain/repository"
	"database/sql"
	"fmt"

    _ "github.com/lib/pq"
)

type orderPersistence struct{}

func NewOrderPersistence() repository.OrderRepository {
	return &orderPersistence{}
}

func (cp orderPersistence) GetOrders(user string, orderDate string) ([]model.Order, error) {

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
    defer db.Close()

    if err != nil {
        fmt.Println(err)
	}
	script := " SELECT"
	script = script + "      order_id"
	script = script + "      , user_nm"
	script = script + "      , item"
	script = script + "      , price"
	script = script + "      , quantity "
	script = script + "      , order_date "
	script = script + "  FROM"
	script = script + "      public.order "
	script = script + "  WHERE 1=1 "
	script = script + "    AND ($1 = '' OR user_nm = $1)  "
	script = script + "    AND ($2 = '' OR order_date = to_date($2, 'YYYY-MM-DD'))  "
	script = script + "  ORDER BY"
	script = script + "      user_nm, order_date desc "
	rows, err := db.Query(script, user, orderDate)

    if err != nil {
		fmt.Println(err)
		return nil, nil
    }
	var res []model.Order
	
	for rows.Next() {
		var e model.Order
		rows.Scan(&e.Id, &e.User, &e.Item, &e.Price, &e.Quantity, &e.OrderDate)
		res = append(res, e)
	}
	
	return res, nil
}

func (cp orderPersistence) Post(order *model.Order) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
    defer db.Close()

    if err != nil {
        fmt.Println(err)
    }

	// INSERT
	script := "  INSERT  INTO public.order(user_nm, item, price, quantity, order_date) "
	script = script + "  VALUES ($1, $2, $3, $4, $5)"
	db.QueryRow(script, order.User, order.Item, order.Price, order.Quantity, order.OrderDate)
}