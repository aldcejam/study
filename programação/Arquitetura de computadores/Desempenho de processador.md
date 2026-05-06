## Clock
- O que é
> Determina a frequência em que o processador executa tarefas, ou seja, se o `Clock` do processador for 3 GHz (unidade padrão), este poderá executar 3 bilhões de ciclos de clock por segundo. Logo, em teoria, quanto maior a frequência de clock maior é a capacidade do processador de executar tarefas.

- Detalhes
> Quanto maior a frequência Clock maior o "esforço" do processador, logo, maior gasto de energia, por conseguinte, mais a sua temperatura pode se elevar.
> Não necessariamente uma maior frequência de clock  trás maior performance, pois algo como a memória ram poderia armazenar poucas informações se comparado a capacidade que o processador tem em processar essas informações, ou seja, se a memória RAM não for compatível com o processador ele pode ter que esperar até que mais dados cheguem para serem processados.


---
## Cache
- O que é
> Memória cache é uma memória de alta velocidade que serve para que o processador não
> precise ficar acessando a memória principal (RAM). Em sumo, é uma memória para dados que a pouco foram acessados, assim o processador as utilize o mais rápido possível para executar o que é que seja.

- Como ela é atualizada: 
> A memória cache é atualizada quando a memória principal é acessada pelo processador. Quando o processador acessa uma memória que não está disponível no cache o controlador do cache busca esses dados da memória principal e atualiza o cache.

**Há duas formas principais da memória cache ser atualizada, são elas:
 - _Localidade Espacial:_ quando uma memória é acessada, possivelmente os endereços de memória próximos serão acessados futuramente, então o controlador do cache salva tás dados a fim de otimizar os próximos processos
 - _Localidade Temporal:_ de comum saber, a memória cache salva dados do que foi acessado muito recentemente.

- Por que é uma memória mais rápida
	1. Proximidade Física
		Na arquitetura de um computador a memória cache é dentro do próprio processador ou muito próxima a ele. Dessa forma o relacionamento processador e cache fica muito mais performático.
	2. Tamanho
		Dado a sua menor capacidade, poucos MB, o acesso ao conteúdo da memória é extremamente rápido.
	3. Estrutura de acesso
		Está memória é projetada com diversas técnicas que melhoram seu desempenho

- Como funciona
	Dividade em 3 partes, L1, L2 e L3, sendo L1 mais rápida que L2 e assim por conseguinte, é uma memória de alta rotatividade em que o processador busca ela a fim de executar processos mais rápido sem depender da memória principal (RAM). 

## Memória pricipal
- Descrição comum: _i7: 64 GB, Xeon: 128 GB_
> Isto significa que, no exemplo, o processador I7 tem capacidade máxima de lidar com memória RAM = 64GB e o Xeon tem capacidade máxima de 128GB. Logo, para lidar com maior quantidade de processor simultâneos o Xeon é melhor.

## Barramento
- **Para que serve:**
> Caminhos físico para conectar componentes de um computador.
> É por onde são transferidos os dados, basicamente por onde dados passam para ir de um componente a outro de um computador. Um exemplo disto é o caminho percorrido para o acesso de memória: memória principal -> memória cache -> processador OU memória principal -> processador.

- **Características:**
> - 64 bits vs 32 bits: dado a maior necessidade de transferência de dados entre componente de um sistema 64 bits os barramentos são mais complexos, com mais caminhos
> - 



- PCI Express 7: Peripheral Component Interconnect Express
> Barramento de alta velocidade para integrar componentes de expansão como placas de rede e dispositivos de armazenamento.

---
## Questões gerais
1. Qual é a função principal do clock em um processador? 
	a) Determinar a quantidade de memória RAM utilizada.  
	b) Controlar a velocidade de acesso aos dispositivos de armazenamento.  
```
	c) Estabelecer a frequência de execução de tarefas pelo processador.  
```
	d) Regular a temperatura interna do computador.  
	e) Gerenciar a transferência de dados entre a CPU e a GPU.
2. O que é a memória cache em um sistema de computador? 
	a) Uma memória de longo prazo para armazenamento de dados.  
    b) Uma área de armazenamento permanente no disco rígido.  
```
    c) Uma memória de alta velocidade usada para armazenar dados frequentemente acessados pelo processador.  
```
    d) Uma memória volátil usada exclusivamente para executar processos de GPU.  
    e) Um tipo de memória que armazena apenas instruções do sistema operacional.
    
3. Como a memória cache é atualizada? 
	a) Manualmente pelo usuário.  
    b) Automaticamente durante a inicialização do sistema operacional.  
```
    c) Quando o processador acessa a memória principal.  
```
    d) Através de um processo de cópia de segurança diária.  
    e) Somente quando o computador é reiniciado.
    
4. Por que a memória cache é mais rápida do que a memória principal (RAM)? 
	a) Porque tem uma capacidade de armazenamento muito maior.  
```
    b) Devido à sua localização física próxima ao processador.  
```
    c) Porque armazena apenas dados de longo prazo.  
    d) Porque utiliza uma tecnologia de armazenamento mais avançada.  
    e) Devido à sua conexão direta com a internet.
    
5. Qual é uma das características que torna a memória cache mais rápida? 
	a) Tamanho físico maior.  
```
    b) Distância física em relação ao processador.  
```
    c) Capacidade de armazenamento ilimitada.  
    d) Estrutura de acesso simplificada.  
    e) Utilização de uma tecnologia de memória mais antiga.
    
6. Qual é a função principal da memória principal (RAM) em um computador? 
	a) Armazenar permanentemente todos os arquivos do sistema operacional.  
    b) Fornecer energia ao sistema quando o computador está desligado.  
```
    c) Armazenar temporariamente dados e instruções usados pelo processador.  
```
    d) Realizar cálculos complexos para a CPU.  
    e) Gerenciar a conexão com dispositivos de entrada e saída.
    
7. O que é o barramento em um computador? 
	a) Um componente responsável pela refrigeração do sistema.  
    b) Um tipo de memória volátil usada para armazenar dados temporários.  
```
    c) Caminhos físicos para conectar componentes internos do computador.  
```
    d) Um dispositivo de armazenamento de longo prazo para o sistema operacional.  
    e) Uma parte do sistema operacional responsável pela gestão de recursos.
    
8. Qual a diferença entre barramentos de 32 bits e 64 bits? 
```
	a) A velocidade de transferência de dados.  
```
    b) O número de componentes internos suportados.  
    c) A quantidade de memória RAM que pode ser endereçada.  
    d) A temperatura máxima suportada pelo barramento.  
    e) A capacidade de processamento do processador.
    
9. Qual é a função do barramento PCI Express? 
```
	a) Conectar dispositivos de entrada, como teclados e mouses.  
```
    b) Gerenciar a energia fornecida aos componentes do computador.  
    c) Transferir dados de alta velocidade entre componentes de expansão.  
    d) Controlar a frequência de operação da CPU.  
    e) Fornecer conectividade de rede sem fio.
    
10. Por que o barramento é crucial para o desempenho do computador?
	a) Porque determina a quantidade de memória disponível.  
    b) Porque controla a velocidade da ventoinha do sistema de resfriamento.  
```
    c) Porque influencia diretamente na velocidade de transferência de dados entre componentes.  
```
    d) Porque gerencia a potência da fonte de alimentação.  
    e) Porque determina a resolução máxima suportada pelo monitor.
    

**Perguntas Dissertativas:**

1. Explique como o clock do processador influencia na execução de tarefas.
> O *clock* do processador determina a quantidade de tarefas possíveis de serem executas em 1s. Caso o processador seja de 3Ghz ele terá a capacidade de fazer 3 bilhões de processos por segundo.
2. Discorra sobre a importância da memória cache no desempenho do processador.
> A memória cache para o processador é um grande aliado na performance devido sua rápida velocidade e retenção de dados comumente/pouco antes acessados, assim, o cache oferece dados de forma extremamente rápida ao processador, não necessitando em alguns casos que o processador acesse a memória ṕr
3. Compare a memória cache com a memória principal (RAM) em termos de velocidade e capacidade.
> A memória principal é de capacidade de ordens de grandeza superior a memória cache, sendo comum memória cache de 6mb e RAM de 4gb - MUITOS gigas (4 gb sendo limite para sistema 32 bits). Principais diferenças entre RAM e Cache no quesito performance  é seu tamanho, que pelo cache possuir pouquíssima memória ela é rapidamente lida, além disso o cache conta com estratégias como sua posição muito próxima do processador ou até mesmo no processador no sistema, como também estratégias de leitura.
4. Descreva como a memória cache é atualizada e por que isso é importante para o desempenho do sistema.
> A memória cache é atualizada seguindo alguns critérios, são eles o relativo a tempo e a proximidade de endereço de memória. O relativo a tempo tem como base o tempo o qual aquele dado foi acessado pelo processador, quanto mais antigo, maior a chance desses dados serem substituídos por outros dados similares. O relativo a endereço de memória segue a probabilidade de que se um endereço de memória foi acessado a pouco tempo, possivelmente os outros endereços próximos a ele também serão acessados, assim, ao acessar salvar endereços de memória o cache pode também salvar endereços próximos para possíveis usos do processador.
5. Analise a relação entre a localidade espacial e a localidade temporal na eficiência da memória cache.
> explicadas na questão anterior
6. Explique como a proximidade física e o tamanho contribuem para a rapidez da memória cache.
> A alta velocidade ofertada por maior proximidade física se dá pois os dados que precisam ir do cache para processador percorrem um caminho menor, portanto possibilitando maior velocidade na transferência. No quesito tamanho é uma premissa simples de seguir, quanto menor a memória, mais fácil é para o processador acessa-lá.
3. Compare e contraste os diferentes níveis de memória cache (L1, L2 e L3) em um processador.
4. Explique o significado das especificações de memória principal em processadores, como "i7: 64 GB, Xeon: 128 GB".
5. Discuta as diferenças entre barramentos de 32 bits e 64 bits, e como isso afeta o desempenho do sistema.
6. Descreva a importância do barramento PCI Express em um computador moderno e cite exemplos de dispositivos que utilizam esse barramento.