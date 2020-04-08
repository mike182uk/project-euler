smallestMultipleOfRange :: [Int] -> Int
smallestMultipleOfRange xs = foldl lcm 2 xs

{-

The below solution was my first attempt, and this brute forces by working out
if a number is evenly divisible by all the numbers in the range. It works ok
for [1..10] but for [1..20] it is very very slow.

-}

smallestMultipleOfRange' :: [Int] -> Int
smallestMultipleOfRange' xs = evenlyDivisibleByAll 2
  where evenlyDivisibleByAll divisor = if all (\a -> divisor `mod` a == 0) xs
                                       then divisor
                                       else evenlyDivisibleByAll (divisor + 1)

--

main = do
  print $ smallestMultipleOfRange [1..20]
