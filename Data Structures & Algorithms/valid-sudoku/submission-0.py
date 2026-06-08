class Solution:    
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        sudoku_row = set()  # Single dictionary that I will reset
        sudoku_columns = [set() for _ in range (9)] # 9 columns itterating over
        sudoku_boxes = [set() for _ in range(9)] # 3 boxes needed 
        
        for row_index, row in enumerate(board):
            # Before each iteration clear the set, so that we have single set in the memory
            sudoku_row.clear()
            
            for column_index, elem in enumerate(row):
                # print(elem)
                
                if self.isDot(elem):
                    continue
                
                # Check Row
                if elem in sudoku_row:
                    #print("Duplicate Row Element: ", elem)
                    return False
                sudoku_row.add(elem)
                
                # Check Column
                if elem in sudoku_columns[column_index]:
                    #print("Duplicate Column Element", elem)
                    return False
                sudoku_columns[column_index].add(elem)
                
                # Check Boxes
                box_index = self.getBoxNumber(row_index, column_index)
                
                if elem in sudoku_boxes[box_index]:
                    return False
                sudoku_boxes[box_index].add(elem)
                         
        return True
    
    def getBoxNumber(self, row, column):
        return math.floor(row / 3) * 3 + math.floor(column / 3)
    
    def isDot(self, elem):
        if elem == ".":
            return True
