## Resumo

> [!quote]
> O pattern criacional singleton nos permite criar uma instancia de uma classe para todo o projeto. 

> [!Question] **O que isso significa?**
>  Que podemos, em quaisquer partes da aplicação, acessar a mesma instância de uma classe, ou seja, o que fizermos com está instancia será repercutida para todas as partes que a usam.

> [!question] **Por que não usar uma variável ou função para todo o sistema?**
> - 1. **Conflitos de Nomenclatura**: Como variáveis globais podem ser acessadas de qualquer lugar, existe o risco de conflitos de nomenclatura, especialmente em projetos grandes.
> - 2. **Acoplamento Forte**: Variáveis globais aumentam o acoplamento entre diferentes partes do código, tornando mais difícil a modificação e o teste dessas partes de forma isolada.
> - 3. **Dificuldade de Rastreamento**: É mais difícil rastrear a origem e o uso de uma variável global, o que pode dificultar a depuração e a manutenção do código.
> - 4. **Problemas de Escopo**: Variáveis globais podem ser modificadas em diferentes partes do código, o que pode levar a bugs difíceis de identificar e corrigir.
    
5. **Segurança**: Variáveis

