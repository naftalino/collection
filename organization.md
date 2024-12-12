# Gacha System (não tenho um nome melhor)

# Banco de dados
- Usuários
    - ID do Telegram
    - Fichas para spins
    - Moedas
    - Banido
    - Qnt. de cartas 
    - Spins totais
    - Data de criação da conta
    - É admin?

- Obras
    - ID
    - Foto da obra
    - Nome
    - Qnt. de cartas
    - As cartas em si
    - Pertence à alguma subobra (estúdio, empresa, etc)
    - Descrição

- Cartas
    - ID
    - Foto da carta
    - Nome
    - Qual o ID da obra da qual a carta pertence?
    - Raridade (1: comum, 2:raro, 3:lendário)

## Suporte de features
- Comandos:
    - start
        - inicia o bot e faz diversas verificações (Se o usuário existe, se ele está banido, etc)
    - help
        - Apenas envia uma mensagem com link dos comandos e suas devidas explicações.
    - spin
        - O giro. Algorítimo em desenvolvimento.
    - perfil
        - Mostra o perfil geral e possibilidade de setar fotos personalizadas de graça, porém, as animadas (gifs) serão coisas pagas.
    - clc
        - mostra a coleção e completa, paginando tudo com 15 cartas por página. Os filtros de coleção serão: Por obra, por raridade, por foto (paginar as cartas por foto)
    - eventos
        - Os eventos são mini coisas que o sistema vai fazer para, por exemplo, desconto em tudo, criação de premios, etc. 
    - ping
        - Apenas para testar se o bot ainda está vivo.
    - rank
        - Ranking de pessoas por categorias. Por exemplo, quantidade total de cartas, quantidade de moedas, quantidade de fichas, etc. Será definido por um top 5.
    - reward
        - Presente diário resgatável. Tudo o que poderá ser resgatado será definido ainda.
    - troca
        - Apenas algumas observações:
            1. Poderá ser feita apenas dentro de grupos oficiais.
            2. Poderá ser feita apenas por usuários aptos, onde não estão banidos e os 2 precisam ter as cartas em questão.
            3. A quantidade de trocas diárias são 25.
    - gift
        - Os usuários podem dar presentes uns aos outros. Não há limite.
    - resgate
        - Um gift resgatável, onde você pode ganhar quaquer ocisa em qualquer quantidade. Cartas, fichas, moedas, etc. Não há limites.
    - obra (por ID e nome)
        - Autoexplicativo. Adicionará o ID, nome e obra.
    - carta (por ID e nome)
        - Mesma coisa com a obra.
