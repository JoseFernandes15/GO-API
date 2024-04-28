package BD

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
    // Conecta-se ao banco de dados PostgreSQL
    db, err := sql.Open("postgres", "user=seu_usuario dbname=sua_base_de_dados password=sua_senha host=seu_host port=sua_porta sslmode=disable")
    if err != nil {
        return nil, err
    }
    
    // Verifica se a conexão foi estabelecida com sucesso
    err = db.Ping()
    if err != nil {
        return nil, err
    }
    fmt.Println("Conexão bem-sucedida com o banco de dados PostgreSQL")
    return db, nil
}
