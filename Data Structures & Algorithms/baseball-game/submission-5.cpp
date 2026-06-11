class Solution {
public:
    int calPoints(vector<string>& operations) {
    vector<int> score;
    score.reserve(operations.size());

    for (const string& op : operations) {
        switch (op[0]) {
            case '+': score.push_back(score.back() + score[score.size()-2]); break;
            case 'D': score.push_back(2 * score.back()); break;
            case 'C': score.pop_back(); break;
            default:  score.push_back(stoi(op)); break;
        }
    }

    return accumulate(score.begin(), score.end(), 0);
}
};