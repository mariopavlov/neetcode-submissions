class Solution {
public:
    int calPoints(vector<string>& operations) {
    vector<int> score;
    score.reserve(operations.size());

    for (const auto& op : operations) {
        switch (op[0]) {
            case '+': score.push_back(score.back() + score[score.size()-2]); break;
            case 'D': score.push_back(2 * score.back()); break;
            case 'C': score.pop_back(); break;
            default: {
                try {
                    score.push_back(stoi(op));
                } catch (const exception& e) {
                    printf("Fail.  Function: %s, Value:  %s\n", e.what(),  op.c_str());
                }

                break;
            } 
        }
    }

    return accumulate(score.begin(), score.end(), 0);
}
};