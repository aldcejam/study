## Resumo
O design Pattern Strategy é um padrão comportamental que nos permite definir uma família de algoritmos, assim podemos escolher um membro dela sem conhecermos como esse membro executa o que queremos, mas sabendo o que ele precisa e o que ele retorna. Isso significa que com o este pattern nós podemos esconder da classe de execução como algo foi feito, e dar a ela somente a responsabilidade de escolher como e em qual momento aquele membro da família será executado.

## Contexto
### Strategy interface
> Esta interface será responsável por definir funções que **deverão** ser implementadas pelas Strategy.
``` TS
interface StrategyInterface{
	execute(a: number, b: number): void;
}
```

### Strategy
> Esta classe será responsável por implementar o código para as funções definidas na Strategy Interface
``` TS
class AddStrategy implements StrategyInterface{
	execute(a: number, b: number): void {
		console.log(a+b);
	}
}

class SubstractStrategy implements StrategyInterface{
	execute(a: number, b: number): void {
		console.log(a-b);
	}
}
```
### Context
>É nesta classe que vamos possibilitar a seleção de uma Strategy.
``` TS
class Context{
	private strategy?: StrategyInterface;
	setStrategy(strategy: StrategyInterface){
		this.strategy = strategy;
	}
	executeStrategy(a: number, b: number){
		return this.strategy?.execute(a, b);
	}
}
```

### Application
>Aqui você poderá usar uma Strategy, agora você não mais terá diversos códigos de mesmo contexto mas com objetivos diferentes na mesma classe.
``` TS
class Application{
	static main(){
		const context = new Context();
		const first = 10;
		const second = 10;
		const operation = 'add';
		
		if (operation === 'add'){
			context.setStrategy(new AddStrategy());
		}
		
		if(operation === 'substract'){
			context.setStrategy(new SubstractStrategy());
		}
		
		context.executeStrategy(first, second);
	}
}
```
