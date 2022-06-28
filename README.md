# API / Server

## Definição

Vamos dividir a API em duas partes: conexão device-server e autenticação usando OAuth.
A conexão device-server (usando Websockets) é feita para que a transferência do token seja feita de forma segura (já que não podemos)
garantir isso por outros motivos.

Graficamente falando, temos isso:

```text
                    ┌──────────┐      1 - Websocket
                    │  Device  ├───────────┐
                    └┬───▲───▲─┘           │
                     │3  │   │             │
                     │   │   │             │
                     │-  │   │             │
                     │   │   │ 2 - Token   │
                     │Q  │   └──────────┬──▼───┐
                     │R  │              │Server│
                     │C  └──────────────┴──▲───┘
                     │o    5 - Command     │
                     │d                    │
                     │e                    │
                     │                     │
                    ┌▼──────┐              │
                    │  App  ├──────────────┘
                    └───────┘   4 - Open
```

A parte da API é introduzida quando vamos responder o app: uma requisição é feita, nós recebemos ela, autenticamos o usuário
e liberamos a porta.

Iremos definir melhor alguns objetivos pra que possamos separar melhor as tarefas durante o trabalho.