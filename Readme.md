
### Taller JS 
Make a static crud with JS and make


### Go migrate

#### Create a new migration 
```sql
migrate create -ext sql -dir cmd/migrate/migrations -seq migration_name
```
#### Up the migration 
```sql
migrate -path cmd/migrate/migrations/ -database "mysql://root:admin@tcp(localhost:3306)/db-name" up
```

#### Down the migration 
```sql
migrate -path cmd/migrate/migrations/ -database "mysql://root:admin@tcp(localhost:3306)/db-name" down
```

#### Fix the migration

```sql
migrate -path cmd/migrate/migrations/ -database "mysql://root:admin@tcp(localhost:3306)/db-name" force 1 
```