<h1 align="center">
  <br>
    <img src=".github/ansible-logo.png" style="background-color: #ffffff" alt="Ansible" width="500">
  <br>
  Minicurso de Ansible
  <br>
</h1>

Este repositório contêm todo o material utilizado durante a ministração do minicurso para a ferramenta Ansible.  
A ideia é abordar os principais conceitos com o objetivo de mostar a montagem de playbooks simples e dar início aos estudos sobre automação de processos com Ansible.  

## Conteúdo

1. Introdução: o que é, para que serve e qual a importância de Ansible?
   1. Principais conceitos
   2. Instalação
2. Principais módulos utilizados em atividades
3. Utilização de filtros em processos
4. Montagem de Playbooks
   1. Estrutura
   2. Loops
   3. Blocks
   4. Condicionais
5. Atividades práticas

## Material disponível

- [Apresentação](./Minicurso%20de%20Ansible.pdf)
- Pequenas aplicações simulando o funcionamento básico de uma loja de ecommerce, utilizadas para a parte prática:
  - `notification_service`
  - `order_processing_service`
  - `user_service`
  - `middleware`
- Soluções base para as atividades propostas


## Testar o funcionamento

1. Listar todos os produtos:

```shell
curl -X GET --location "http://localhost:8081/product" \
    -H "Accept: application/json"
```

2. Criar novo usuário, lembre de alterar o nome e email (coloque algum que possa acessar):

```shell
curl -X POST --location "http://localhost:8081/user" \
    -H "Content-Type: application/json" \
    -d '{
          "name": "",
          "email": "",
          "document": "47406432275",
          "address": {
            "city": "Olinda",
            "number": 524,
            "state": "PE",
            "street": "Rua Angelina Guimarães da Silva"
          }
        }'
```

3. Criar um pedido, atualizar o json com o id do usuário criado:

```shell
curl -X POST --location "http://localhost:8081/order" \
    -H "Content-Type: application/json" \
    -d '{
          "userId": 0,
          "companyId": 1,
          "products": [
            {
              "id": 4,
              "quantity": 1
            }
          ]
        }'
```
