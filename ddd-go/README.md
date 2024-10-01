[How To Implement Domain-Driven Design (DDD) in Golang](https://programmingpercy.tech/blog/how-to-domain-driven-design-ddd-golang/)
[How to Structure DDD in Golang](https://programmingpercy.tech/blog/how-to-structure-ddd-in-go/)

**Dúvidas**

- Na função `CreateOrder` o sistema devia atualizar os dados do `Customer` para adicionar os produtos pedidos
  - Pra isso eu precisaria atualizar o `domain` do customer com uma função de settar os produtos
  - A função teria que verificar se o produto já foi adicionado e aumentar a quantidade dele
    - Essa verificação deve ficar no domain e não no service
- `CreateOrder` não deveria retornar o preço total, só adicionar produtos no cliente
- Retornar o preço total e realizar a transação deveria ficar no `BillingService` que não foi criado ainda
  - A transação, assim como os produtos, deve ser atualizado no `Customer`
  - Como a taverna não tem um UserID, n faz sentido a transaction ter um `from` `to`