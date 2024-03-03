# Atividade 3

## Perguntas do roteiro

- Pergunta: O que acontece se você utilizar o mesmo ClientID em outra máquina ou sessão do browser? Algum pilar do CIA Triad é violado com isso?

**Resposta**:

Se o mesmo "ClientID" for utilizado em outra máquina ou sessão do navegador, isso pode afetar a confidencialidade e a disponibilidade dos dados.
Se o "ClientID" é usado como parte de um sistema de autenticação para acessar recursos protegidos, o compartilhamento dele entre diferentes máquinas ou sessões do navegador pode comprometer a confidencialidade dos dados, pois alguém que tenha acesso ao "ClientID" poderá acessar os recursos protegidos. Além disso, acessando em diferentes sessões o client é desconectado sem nenhum aviso.

- Pergunta: Com os parâmetros de resources, algum pilar do CIA Triad pode ser facilmente violado?
```
version: "3.7"
services:
    deploy:
      resources:
        limits:
          cpus: '0.01'
          memory: 100M
```

**Resposta**:

Se os limites de recursos forem definidos muito baixos, isso pode comprometer a integridade das operações do serviço. Como por exemplo, se o serviço estiver processando dados críticos e não tiver recursos suficientes para concluir as operações de forma adequada, isso poderá resultar em falhas ou corrupção de dados.
Os limites de recursos têm o potencial de impactar significativamente a disponibilidade do serviço. Se os recursos forem limitados de forma muito restrita, o serviço pode ficar indisponível ou responder lentamente, especialmente durante períodos de pico de uso ou sob carga intensa.

- Pergunta: Sem autenticação (repare que a variável allow_anonymous está como true), como a parte de confidencialidade pode ser violada?

  **Resposta**:

  Se a variável "allow_anonymous" estiver como true o pilar de confidencialidade pode ser violado de algumas maneiras:
-  Qualquer pessoa pode se conectar ao serviço MQTT e receber ou enviar mensagens, mesmo que não tenha permissão pra isso.
-  O usuário pode interceptar mensagens de qualquer outra pessoa que esteja na mesma rede.

## Tente simular uma violação do pilar de Confidencialidade
Para comprometer a confidencialidade, poderíamos nos inscrever em um tópico sem autorização adequada, permitindo a interceptação das mensagens enviadas. Essa violação pode ser realizada por meio de técnicas de monitoramento de rede, tais como a captura de pacotes ou ataques de "man-in-the-middle".

```
mosquitto_sub -h endereço_do_broker -t tópico
```

## Tente simular uma violação do pilar de Integridade
Para comprometer o pilar de integridade, poderíamos realizar uma série de ações, tais como a injeção, exclusão ou envio de dados falsificados, resultando na adulteração dos dados transmitidos. Essas ações podem ser executadas de diversas maneiras, incluindo a exploração de vulnerabilidades no sistema de comunicação MQTT.

Ao injetar dados falsos, seria possível introduzir informações fictícias ou imprecisas na corrente de dados, potencialmente levando a decisões erradas com base em informações incorretas. Da mesma forma, a exclusão de dados legítimos pode distorcer a visão geral e impedir a análise precisa dos eventos ou condições monitoradas.

## Tente simular uma violação do pilar de Disponibilidade
Para simular uma violação na disponibilidade poderíamos sobrecarregar o sistema com uma grande quantidade de conexões simultâneas para as portas disponíveis. Essas solicitações de conexão em massa sobrecarregariam os recursos do servidor MQTT, impedindo que ele respondesse adequadamente às solicitações legítimas.  



