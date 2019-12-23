class Money {
  amount: number

  constructor(amount: number) {
    this.amount = amount
  }

  equals(obj: Money) {
    return this.amount == obj.amount
  }
}

class Dollar extends Money {
  times(multiplier: number) {
    const obj: Dollar = new Dollar(this.amount * multiplier)
    return obj
  }
}

class Franc extends Money {
  times(multiplier: number) {
    const obj: Franc = new Franc(this.amount * multiplier)
    return obj
  }
}

const dollar: Dollar = new Dollar(5)
console.log(dollar.times(2).amount)
console.log(dollar.equals(new Dollar(10)))

const franc: Franc = new Franc(5)
console.log(franc.times(2).amount)
console.log(franc.equals(dollar))

