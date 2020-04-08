sumSquareDiff :: [Int] -> Int
sumSquareDiff xs = squareOfSum - sumOfSquares
  where sumOfSquares = sum $ map (^2) xs
        squareOfSum  = foldl (+) 0 xs ^ 2

--

main = do
  print $ sumSquareDiff [1..100]
