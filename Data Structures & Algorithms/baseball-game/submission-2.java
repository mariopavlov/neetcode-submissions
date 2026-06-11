class Solution {
    public int calPoints(String[] operations) {
        List<Integer> score = new ArrayList<Integer>();

        for (String op : operations) {
            switch (op) {
            case "+":
                var  a = score.get(score.size() - 1);
                var  b = score.get(score.size() - 2);
                score.add(a + b);
                break;
            case "D":
                var prev = score.get(score.size() - 1);
                score.add(2 * prev);
                break;
            case "C":
                score.remove(score.size() - 1);
                break;
            default:
                int val = Integer.parseInt(op);
                score.add(val);
                break;
            }
            
        }

        int total = score.stream().mapToInt(Integer::intValue).sum();
        
        return total;
    }
}