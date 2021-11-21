class Solution {
public:
    int majorityElement(vector<int>& nums) {
        int max = 0;
        int count = 0;
        // C++ 11标准中加入了unordered系列的容器。
        // unordered_map记录元素的hash值，根据hash值判断元素是否相同。 
        // map相当于java中的TreeMap，unordered_map相当于HashMap
        unordered_map<int,int> mymap;
        for(auto num:nums){
            mymap[num]++;
            if(mymap[num]>count){
                count = mymap[num];
                max = num;
            }
        }
        return max;
    }
};