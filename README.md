# bank

## Database setup

Requires MySQL 5.1+. The command below sets up a new database called 'bank' and a grant for bank@localhost.

```
mysql -uroot -p < ./sql/create.sql
```

## Routes

GET /transactions
GET /transaction/{id}
POST /transaction
DELETE /transaction/{id}
