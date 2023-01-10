# Valid Pass

## Como usar?

### Passo a passo

1 - Inicie o programa com o comando `docker compose up`.Após iniciado o programa estará escutando em `http://localhost/PORT_SERVICE`.
OBS: `PORT_SERVICE` deverá conter a porta do serviço que você definiu no docker-compose.yml. Em desenvolvimento o PORT_SERVICE tem valor 8080.

2 - Método `POST`.

3 - Rota `verify`.

4 - Monte a requisição. Nela deve conter um header e um body:

#### Header

| Atributo      | Valor              | Obrigatório |
| :------------ | :----------------- | :---------: |
| Authorization | Bearer token       |     \*      |
| Content-type  | `application/json` |     \*      |

Para fins didáticos o Bearer token será o seguinte:  
`Bearer 3406b8d48918e96a912111467f070e4be22ea2402a1633e14d3ae4febf47598b`

#### Body

| Atributo | Tipo        | Valor                    | Obrigatório |
| :------- | :---------- | :----------------------- | :---------: |
| password | string      | Senha à ser autenticada. |     \*      |
| rules    | ObjectRules | Array de regras.         |     \*      |

##### ObjectRules

| Atributo | Tipo   | Valor                                              | Obrigatório |
| :------- | :----- | :------------------------------------------------- | :---------: |
| rule     | string | `Tipo de regra` a ser aplicada.                    |     \*      |
| value    | int    | Quantidade de letras em que será aplicada a regra. |     \*      |

##### Tipos de Regras

| Nome            |
| :-------------- |
| minSize         |
| minUppercase    |
| minLowercase    |
| minDigit        |
| minSpecialChars |
| noRepeted       |

##### Exemplo em json

OBSERVAÇÃO: se desejar passar uma barra invertida, \, lembre-se de passar duas barras invertidas para que o json leia corretamente sua requisição, caso contrário resultará em falha.

```json
{
  "password": "!TesteSenháForteÀ!123&",
  "rules": [
    {
      "rule": "minSize",
      "value": 8
    },
    {
      "rule": "minUppercase",
      "value": 8
    },
    {
      "rule": "minLowercase",
      "value": 4
    },
    {
      "rule": "minDigit",
      "value": 2
    },
    {
      "rule": "minSpecialChars",
      "value": 4
    },
    {
      "rule": "noRepeted",
      "value": 2
    }
  ]
}
```

### Response

#### 200

| Atributo | Tipo         | Valor                                                                                |
| :------- | :----------- | :----------------------------------------------------------------------------------- |
| verify   | boolean      | Resposta indicando se a validação passou nas regras.                                 |
| noMatch  | Array string | Regras que não passaram na validação. Caso verify esta true esse campo estará vazio. |

#### Exemplo

```json
{
  "verify": false,
  "noMatch": ["minUppercase"]
}
```

### 401

Caso receba um erro 405, certifque que colocou o campo Authorization no header e se seu valor está correto.

### 405

Caso receba um erro 405, certifique que o método que está chamando é do tipo correto (POST, GET, PUT, ...).
