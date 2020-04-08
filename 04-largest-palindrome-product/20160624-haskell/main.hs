largestProductPalindromeForRange :: [Int] -> Int
largestProductPalindromeForRange n = maximum [ z |
                                               x <- n, y <- n,
                                               let z    = x * y,
                                               let zStr = show z,
                                               zStr == reverse zStr
                                             ]

{-

The above solution works well for [100..999] but for [1000..9999] it takes
a very, very long time. This is because it has to calculate all the palindromes
between the range before maxium can be called. The solution below attempts to
circumvent this by not needing to generate all the palindromes up front. This solution
works well upto [100000..999999]

-}

largestProductPalindromeForRange' :: [Int] -> Int
largestProductPalindromeForRange' n = findLargestPalindrome palindromes
  where
    findLargestPalindrome (x:y:z) = if x > y
                                    then x
                                    else findLargestPalindrome (y:z)
    palindromes = [ z |
                    x <- reverse n, y <- n,
                    let z    = x * y,
                    let zStr = show z,
                    zStr == reverse zStr
                  ]

--

main = do
  print $ largestProductPalindromeForRange [100..999]
