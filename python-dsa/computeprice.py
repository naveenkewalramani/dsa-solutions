def computePrice(currentCost, reductionFactor, minCost):
    if currentCost <= minCost:
        return minCost
    if currentCost-reductionFactor <=minCost:
        return minCost
    return currentCost-reductionFactor