# fitness-release-2



 kallax migrate --input ./model/ --out ./migrations --name init_db


kallax migrate up --dir ./migrations --dsn 'postgres:password@localhost:5433/test?sslmode=disable' --steps 1
kallax migrate up --dir ./my-migrations --dsn 'user:pass@localhost:5432/dbname?sslmode=disable' --version 1493991142
# fitness-release-2
