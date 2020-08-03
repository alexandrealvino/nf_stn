**NF_STN**

API REST desenvolvida em Go com as seguintes funcionalidades:

- Listar invoices (notas fiscais) com paginação

- Buscar invoice por documento

- Cadastrar invoice

- Deleção lógica de invoice

- Listar invoices com filtros por mês, ano e documento

- Listar invoices com ordenação por mês, ano, documento ou combinações desses

Persistência desenvolvida utilizando o banco de dados MySQL.

Autenticação de rotas por token de aplicação. 

**Como rodar o código**

Para rodar a aplicação basta executar o seguinte comando no seu terminal:

docker-compose up -d

Executando os testes:

go test ./...

**Funcionamento da API**

Para acessar os endpoints é necessário realizar o login do usuário, o banco 
de dados já é iniciado com uma tabela de usuários que possui uma linha 
contendo um "username" e uma "hash" gerada a partir da senha "password".
A aplicação então compara as hashs e autoriza ou não o usuário. Caso 
o usuário seja autenticado, a aplicação retorna um token do tipo "Bearer Token" para ser utilizado 
no header de autenticação para as demais rotas. Tal token tem validade de 15min
podendo ser renovado na endpoint api/refresh, e fica armazenado no servidor do Redis para 
futuras verificações.

**Parâmetros de busca**

Para ordenação utilizar o parâmetro "orderBy" no request, com as seguintes opções de ordenação: 
"document", "year, "month" ou uma composição delas separadas por vírgulas.

Para busca por mês, ano ou documento utilizar os parâmetros "month", "year" 
ou "document" respectivamente.

Segue uma postman collection com exemplos de requests (utilizar 
um token novo).


