// Файл: calc.js (Исправленная версия)

function add(a, b) {
  return a + b; // ИСПРАВЛЕНО: выполняет сложение
} // addition

function sub(a, b) {
  return a - b; // Эта функция была верной
} // subtraction

function mul(a, b) {
  return a * b; // ИСПРАВЛЕНО: выполняет умножение
} // multiplication

function div(a, b) {
  if (b === 0) throw new Error('division by zero') // throw on zero
  return a / b
}

module.exports = { add, sub, mul, div }