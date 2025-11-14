def add(a, b)
  a + b
end

def sub(a, b)
  a - b
end

def mul(a, b)
  a * b
end

def div(a, b)
  raise ZeroDivisionError, "division by zero" if b == 0
  a.to_f / b
end