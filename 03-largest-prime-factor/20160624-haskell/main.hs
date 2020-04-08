findLargestPrimeFactor :: Int -> Int -> Int
findLargestPrimeFactor value divisor
  | value == divisor         = divisor
  | value `mod` divisor == 0 = findLargestPrimeFactor (value `div` divisor) divisor
  | otherwise                = findLargestPrimeFactor value nextDivisor
  where nextDivisor = if even (divisor + 1) then divisor + 2 else divisor + 1

largestPrimeFactorOf :: Int -> Int
largestPrimeFactorOf x = findLargestPrimeFactor x 2

--

main = do
  print $ largestPrimeFactorOf 600851475143
