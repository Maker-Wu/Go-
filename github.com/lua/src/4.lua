function square(iteratorMaxCount, currentNum)
    if currentNum < iteratorMaxCount then
        currentNum = currentNum + 1
        return currentNum, currentNum*currentNum
    end
end
for i,n in square, 3, 0 do
    print(i, n)
end