class Solution:
    def calPoints(self, operations: List[str]) -> int:
        record = []

        for op in operations:
            match op:
                case "+":
                    a = record[-1]
                    b = record[-2]
                    record.append(a + b)
                case "D":
                    a = record[-1]
                    record.append(2 * a)
                case "C":
                    record.pop()
                case _:
                    score = int(op)
                    record.append(score)
        
        result = 0

        for r in record:
            result = result + r


        return result

        