const express = require('express');
const app = express();
const PORT = 5000;

const tempconv = require('./util/tempconv.js')

app.get('/celsius/:temp', (req, res) => {
  const temp = req.params.temp
  res.json(tempconv.celsius(temp))
})

app.get('/fahrenheit/:temp', (req, res) => {
  const temp = req.params.temp
  res.json(tempconv.fahrenheit(temp))
})

app.get('/ctof/:temp', (req, res) => {
  const temp = req.params.temp
  res.json(tempconv.CToF(temp))
})

app.get('/ftoc/:temp', (req, res) => {
  const temp = req.params.temp
  res.json(tempconv.FToC(temp))
})

app.listen(PORT, function(){
  console.log('Servidor rodando na URL http://localhost:5000')
});