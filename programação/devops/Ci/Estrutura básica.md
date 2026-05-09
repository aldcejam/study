``` json
// (.github/workflows/ci.yml) //

name: ci-js-workflow

on: 
	pull_request:
		branches: 
			- dev // Definir branch onde job será executado

jobs:
	check-application: // Nome do Job
		runs-on: ubuntu-latest // Ambiente
		steps:
			- uses: actions/checkout@v4 // acesso ao código
			- uses: actions/setup-node@v4 // define ambiente node
			with:
				- run: npm install 
				- run: npm test
				- run: node src/index.js
```
