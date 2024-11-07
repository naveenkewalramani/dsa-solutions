
import math

def divisiorsOfNaturalNumber(number)->list:
    if number == 0:
        return []
    divisors = []
    for i in range(1, int(math.sqrt(number))+1):
        if number % i == 0:
            if number // i == i:
                divisors.append(i)
            else:
                divisors.append(i)
                divisors.append(number // i)
    return divisors