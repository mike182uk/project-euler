isMultipleOf :: Int -> Int -> Bool
isMultipleOf x y = x `mod` y == 0

isMultipleOf3 :: Int -> Bool
isMultipleOf3 x = isMultipleOf x 3

isMultipleOf5 :: Int -> Bool
isMultipleOf5 x = isMultipleOf x 5

isMultipleOf3Or5 :: Int -> Bool
isMultipleOf3Or5 x = isMultipleOf3 x || isMultipleOf5 x

sumOfAllValuesThatAreMultipesOf3Or5 :: [Int] -> Int
sumOfAllValuesThatAreMultipesOf3Or5 xs = sum $ filter isMultipleOf3Or5 xs

--

main = do
  print $ sumOfAllValuesThatAreMultipesOf3Or5 [0..999]
