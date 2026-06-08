class Solution:
    box_size = 3
    
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        sudoku_row = set()  # Single dictionary that I will reset
        sudoku_columns = [set() for _ in range (9)]
        sudoku_boxes = [set() for _ in range(9)]
        
        for row_index, row in enumerate(board):
            # Before each iteration clear the set, 
            # so that we have single set in the memory
            sudoku_row.clear()
            
            for column_index, elem in enumerate(row):                
                if self.isEmpty(elem):
                    continue
                
                # Check Row
                if elem in sudoku_row:
                    return False
                sudoku_row.add(elem)
                
                # Check Column
                if elem in sudoku_columns[column_index]:
                    return False
                sudoku_columns[column_index].add(elem)
                
                # Check Boxes
                box_index = self.getBoxNumber(row_index, column_index)
                
                if elem in sudoku_boxes[box_index]:
                    print("Box Number: ", box_index)
                    print("Duplicate Eelement. Row: ", row_index, ", Column: ", column_index)
                    return False
                sudoku_boxes[box_index].add(elem)
                
        return True
    
    def getBoxNumber(self, row, column):
        return (row // 3) * 3 + (column // 3)
    
    def isEmpty(self, elem):
        return elem == "."