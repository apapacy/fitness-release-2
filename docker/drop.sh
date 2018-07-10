docker exec -it fitness_postgres_master psql -U postgres -c 'drop database test';
docker exec -it fitness_postgres_master psql -U postgres -c 'create database test';
