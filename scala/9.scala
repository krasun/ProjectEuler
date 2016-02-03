object Application {
  def productOfSpecificPythagoreanTriplet(): Option[Int] = {
    for (a <- 1 to 1000; b <- 1 to 1000; c <- 1 to 1000) {
      if ((a < b) && (b < c) && ((a + b + c) == 1000) && ((a * a + b * b) == c * c)) {
        return new Some(a * b * c)
      }
    }

    None
  }

  def main (args: Array[String]) {
    val product = productOfSpecificPythagoreanTriplet()
    println(s"The product of abc is $product.")
  }
}
