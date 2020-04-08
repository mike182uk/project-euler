-- basic sieve implementation
primes :: [Int]
primes = findPrimes 2 [3,5..]
  where findPrimes p xs = p : findPrimes (head onlyPrimes) onlyPrimes
          where onlyPrimes = filter (\x -> x `mod` p /= 0) xs

getPrimeAtIndex :: Int -> Int
getPrimeAtIndex n = primes !! (n - 1)

--

main = do
  print $ getPrimeAtIndex 10001
