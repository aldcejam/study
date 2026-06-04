## Custo
Os custos de manter conexões websocket dependerão explicitamente da linguagem e frameworks adotados, um valor base de gasto de memória por conexão seria entre **5k e 50k.**

| **Tecnologia / Framework**            | **Consumo Médio por Conexão**                                    | **Conexões possíveis com 1 GB de RAM** |
| ------------------------------------- | ---------------------------------------------------------------- | -------------------------------------- |
| **Go (Gorillas/epoll personalizado)** | ~2 KB a 4 KB                                                     | ~250.000 a 500.000                     |
| **Node.js (biblioteca `ws`)**         | ~10 KB a 30 KB                                                   | ~30.000 a 100.000                      |
| **Java (Spring Boot / Tomcat)**       | ~15 KB a 50 KB _(pode ser maior se usar uma Thread por conexão)_ | ~20.000 a 60.000                       |
| **Python (FastAPI / WebSockets)**     | ~15 KB a 40 KB                                                   | ~25.000 a 65.000                       |
!!!! EXTREMO CONSUMO QUANDO:!!!! thread-per-request 