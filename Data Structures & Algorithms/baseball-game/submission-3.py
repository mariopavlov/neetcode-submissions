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
                    try:
                        record.append(int(op))
                    except ValueError:
                        print(f"invalid operation: {op}")
        
        return sum(record)

        