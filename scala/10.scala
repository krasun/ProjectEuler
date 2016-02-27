object Application {
  def isPrime(number: Int): Boolean = {
    number == 2 || (!(2 to math.sqrt(number).toInt).exists(number % _ == 0))
  }

  def primeSum(limit: Int): Long = {
    var sum:Long = 0
    for (number <- 2 to limit) {
      if (isPrime(number)) {
        sum += number
      }
    }

    return sum
  }

  def main (args: Array[String]) {
    val primeSumUntilTwoMillion = primeSum(2000000)
    println(s"The sum of all the primes below two million is $primeSumUntilTwoMillion.")
  }
}
