# tempconv-api
API para utilizar a biblioteca tempconv

**Parâmetro /:temp deve receber um valor inteiro para retornar resposta correta.**

# Rotas para utilizar


URL-Principal <br>
https://tempconv-api.vercel.app/

Rotas GET <br>

<hr>


**/celsius/:temp** <br>

String imprime uma temperatura 'n' em Celsius no formato n°C <br>

[Exemplo](https://tempconv-api.vercel.app/celsius/30, "https://tempconv-api.vercel.app/celsius/30")

``Exemplo: https://tempconv-api.vercel.app/celsius/30 ``

``Resposta: "30°C" ``

<hr>


**/fahrenheit/:temp** <br>

String imprime uma temperatura 'n' em Fahrenheit no formato n°F <br>

[Exemplo](https://tempconv-api.vercel.app/fahrenheit/30, "https://tempconv-api.vercel.app/fahrenheit/30")

``Exemplo: https://tempconv-api.vercel.app/fahrenheit/30 ``

``Resposta: "30°F" ``

<hr>


**/ctof/:temp** <br>

CToF converte uma temperatura em Celsius para Fahrenheit

[Exemplo](https://tempconv-api.vercel.app/ctof/30, "https://tempconv-api.vercel.app/ctof/30")

``Exemplo: https://tempconv-api.vercel.app/ctof/30 ``

``Resposta: "86°F" ``

<hr>

**/ftoc/:temp** <br>

FToC converte uma temperatura em Fahrenheit para Celsius

[Exemplo](https://tempconv-api.vercel.app/ftoc/30, "https://tempconv-api.vercel.app/ftoc/30")

``Exemplo: https://tempconv-api.vercel.app/ftoc/30 ``

``Resposta: "-1.11°C" ``