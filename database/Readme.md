### Modelo de dados.
Se tratando de um CMS, informações sobre permissão e distribuição de dados, templates HTML e arquivos precisam estar descritas em um banco de dados pois podem ser alteradas dinamicamente.

Este banco é uma proposta de modelo que viabiliza o controle de permissão utilizando role-based access control (RBAC) e um modelo de descrição de dados que agrupa informações em views. Tenta-se também, adquirir um modelo que salve e mantenha um registro de estados anteriores, para ter controle de versão de conteúdo e de dados sensíveis.

Foram dividos em quatro schemas SQL (namespaces): Controller, Session Management, Content e Users.

O Controller controla a permissão para determinadas operações.
O Session Management interliga o módulo usuário ao Controller
O Usuário persiste dados dos usuários e de login.
O Content armazena e agrupa dados em vizualizações.

Existe certo acoplamento entre esses "modulos", uma vez que existem referêcias entre eles.

A ideia é que se implementa uma camada de abstração OO para esse modelo conveniente, não necessáriamente um mapeamento direto (isto é, nem toda tabela vai produzir uma classe);
