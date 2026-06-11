class Solution {
public:
    int calPoints(vector<string>& operations) {
        vector<int> score;
        for (const string& op : operations) {
            if ("+" == op) {
                int a = score.at(score.size() - 1);
                int b = score.at(score.size() - 2);
                score.push_back(a + b);

                continue;
            }
            if ("D" == op) {
                int a = score.back();
                score.push_back(2 * a);
                continue;
            }
            if ("C" == op) {
                score.pop_back();
                continue;
            }

            score.push_back(stoi(op));
        }

        return accumulate(score.begin(), score.end(), 0);
    }
};