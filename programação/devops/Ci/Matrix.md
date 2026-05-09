Resumo: o matrix é uma forma simples e eficiente de testar nossas aplicações em múltiplos ambientes e versões. Por exemplo, podemos testar em ambientes Windows e Linux, bem como em diferentes versões de uma linguagem de programação, como Node.js 18 e Node.js 22.

Exemplo de código:
``` json
// (.github/workflows/ci.yml) //

name: ci-js-workflow

on: 
	pull_request: // Situação
		branches: 
			- dev // Definir branch onde job será executado

jobs:
	check-application:
		runs-on: ubuntu-latest
		strategy:
			matrix:
				node-version: [14.x, 15.x, 16.x, 20.x] // app será testa nessas V
		steps:
			- uses: actions/checkout@v4
			- uses: actions/setup-node@v4
			with:
				node-version: ${{ matrix.node-version }}
				- run: npm install
				- run: npm test
				- run: node src/index.js
```