# go-bjf

[![GoDoc](https://godoc.org/github.com/xor-gate/go-bjf?status.svg)](https://godoc.org/github.com/xor-gate/go-bjf)[![Build Status](https://travis-ci.org/xor-gate/go-bjf.svg)](https://travis-ci.org/xor-gate/go-bjf)[![Coverage Status](https://coveralls.io/repos/xor-gate/go-bjf/badge.svg?branch=master&service=github)](https://coveralls.io/github/xor-gate/go-bjf?branch=master)

Golang implementation of bijective function algorithm used in url shorters, see:
http://stackoverflow.com/questions/742013/how-to-code-a-url-shortener

I would continue your "convert number to string" approach. However you will realize that your proposed algorithm fails if your ID is a *prime and greater than 52*.

### Theoretical background

You need a [Bijective Function][1] *f*. This is necessary so that you can find a inverse function *g('abc') = 123* for your *f(123) = 'abc'* function. This means:

 * There must be no *x1, x2 (with x1 ≠ x2)* that will make *f(x1) = f(x2)*,
 * and for every *y* you must be able to find an *x* so that *f(x) = y*.

### How to convert the ID to a shortened URL

 1. Think of an alphabet we want to use. In your case that's `[a-zA-Z0-9]`. It contains *62 letters*.
 1. Take an auto-generated, unique numerical key (the auto-incremented `id` of a MySQL table for example).

    For this example I will use 125<sub>10</sub> (125 with a base of 10).

 1. Now you have to convert 125<sub>10</sub> to X<sub>62</sub> (base 62).
  
    125<sub>10</sub> = 2×62<sup>1</sup> + 1×62<sup>0</sup> = `[2,1]`

    This requires use of integer division and modulo. A pseudo-code example:
        
        digits = []

        while num > 0
          remainder = modulo(num, 62)
          digits.push(remainder)
          num = divide(num, 62)

        digits = digits.reverse
          

    Now map the *indices 2 and 1* to your alphabet. This is how your mapping (with an array for example) could look like:

        0  → a
        1  → b
        ...
        25 → z
        ...
        52 → 0
        61 → 9

    With 2 → c and 1 → b you will receive cb<sub>62</sub> as the shortened URL.

        http://shor.ty/cb

### How to resolve a shortened URL to the initial ID

The reverse is even easier. You just do a reverse lookup in your alphabet.

 1. e9a<sub>62</sub> will be resolved to "4th, 61st, and 0th letter in alphabet".

    e9a<sub>62</sub> = `[4,61,0]` = 4×62<sup>2</sup> + 61×62<sup>1</sup> + 0×62<sup>0</sup> = 19158<sub>10</sub>

 1. Now find your database-record with `WHERE id = 19158` and do the redirect.


### Some implementations (provided by commenters)

 - [Ruby][2]
 - [Python][3]
 - [CoffeeScript][4]
 - [Haskell][5]
 - [Perl][7]
 - [C#][8]


  [1]: http://en.wikipedia.org/wiki/Bijection
  [2]: https://gist.github.com/1073996
  [3]: https://gist.github.com/778542
  [4]: https://gist.github.com/1158171
  [5]: https://gist.github.com/4626401
  [6]: https://gist.github.com/9554733
  [7]: https://metacpan.org/pod/Short::URL
  [8]: https://gist.github.com/9554733
