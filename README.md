# Aplicacao Go Lang - FullCycle
[Aplicações escaláveis com Go Lang - Full Cycle](https://www.youtube.com/watch?v=nTmZlzwTErM)<br>

<br>
Foi desenvolvido uma aplicação para processamento de uma transação em Go Lang, o obejetivo com este serviço foi mais para aprender o que erra o Go e como se programava com ele pois é uma linguagem que está popular, e está no meu portifólio de aprendizagem.<br>
Tinha como regras de negócio que a transação não podia ser maior que R$1000 e menor que R$1 caso nao tendesse essas regras, retornava o status reject e uma mensagem dizendo qual foi a rejeição.<br>
No final implemetado um servidor REST também, para mostrar que ao mesmo tempo que podemos estar processando una requição do RabbitMq e podemos estar recebendo requisições rest, e que o processo vai continuar funcionando normalmente, ou seja independente da minha entrada, eu consigo implementa-la e fazer funcionar.

# O que foi apresentado
Foi apresentado toda a extrutura do Go Lang como é criado Funções e Métodos, como é criado uma struct que é uma estrutura de dados muito utilizado para DTO(Data Transfer Object) e utilizado para tipagem de extruturas de dados, foi apresentado a questão de [Clean Architeture](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg), não é algo novo mas muito utilizado em projetos escalaveis e descentralizados, e nao menos importante a questão de testes, como podemos fazer os testes automatizados, e aplica-los em projetos Go para agilizar o desenvolvimento e garantirmos a integridade da nossa aplicação.

# O que foi utilizado
 - Go Lang - totalemente limpo
 - go-sqlite3 - Utilizado para o banco de dados
 - GoMock - Para conseguirmos simular um banco de dados
 - Rest - Implementado um server Rest para receber requisições
 
# Links Uteis
[Clean Architecture - Implementação do server Rest](https://www.youtube.com/watch?v=YaGVURjB33I)<br>
[Págida do Go Lang](https://go.dev/)<br>
[Págida de packs do GO](https://pkg.go.dev/)<br>




