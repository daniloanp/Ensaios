## Descrição das funcionalidade da proposta de projeto

O projeto proposto é um CMS que tem como objetivo inicial gerir o conteúdo de um portal de notícias.
Eventualmente o projeto poderá sofrer uma extensão nas funcionalidades (e.g. suportar uma loja virtual embutida).

#### Descrição em alto nível das funcionalidades básicas
---

* Administração de usuários (Funcionalidade básica de um CRUD e funções correlatas)
* Controle de permissões/atribuição do usuário (um usuário pode ser um "editor", "administrador" ou um "visitante")
* Administração da organização do conteúdo (divisão em seções, categorizações e metadados associados);
* Interface de produção de conteúdo por um usuário na atribuição de editor, isto é, artigos, textos, ensaios, crônicas;
* Interface de produção de conteúdo por um usuário na atribuição de leitor (comentários, opiniões, processos de cadastro e afins);
* Interface com dados estatísticos sobre o acesso e popularidade dos conteúdos disponibilizados pelos editores.

#### Descrição da proposta de implementação
-----

* Modelar a persistência utilizando o Modelo de Entidade e relacionamento. e implementar em um SGBD SQL;
* Escrever um servidor HTTP que vai oferecer uma API RESTful com as funcionalidades descritas;
* Escrever uma aplicação cliente Web que ofereça a camada de visualização do sistema;

#### Tecnologias
-------------

Por afinidade, experiência e conhecimento as seguintes tecnologias foram escolhidas:
- O SGBD escolhido é o PostgreSQL;
- A linguagem do Servidor de backend será Go(golang.org);
- E a linguagem de script pra aplicação web será Dart(dartlang.org), que compila pra JavaScript, mas tem tipos e abstrações melhores pra OOP.





	
