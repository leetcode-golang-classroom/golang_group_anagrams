# golang_group_anagrams

Given an array of strings `strs`, group **the anagrams** together. You can return the answer in **any order**.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Examples

**Example 1:**

```
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

```

**Example 2:**

```
Input: strs = [""]
Output: [[""]]

```

**Example 3:**

```
Input: strs = ["a"]
Output: [["a"]]

```

**Constraints:**

- `1 <= strs.length <= $10^4$`
- `0 <= strs[i].length <= 100`
- `strs[i]` consists of lowercase English letters.

## 解析

給定一個字串陣列 strs,

定義一個字串 s  的 anagram  t 就是把原本的 s 的字元做重新排列產生的字串

字串 t 如果是字串 s 的 anagram 代表

 t 與 s 的每個字元出現頻率相同

![](https://i.imgur.com/E5q2CGV.png)


題目要求把 strs 根據 anagram 做分類把相同 anagram 的放在一起，回傳分類後的結果

根據 anagram 定義

且出現字元皆是小寫英文字母

每個字串可以透過長度 26 的整數陣列freq來代表其字元出現的頻率

因為 ASCII code 編碼剛好是鄰近的26個數字

每個 freq[pos] 代表 pos + ‘a' 這個字元所出現的次數

![](https://i.imgur.com/qERp3pL.png)

所以可以透過 這個 freq 當作是一個 hashKey 來把每個 str 做分類

假設字串長度最長的是 m 而總共有 n 個字串

所以每次計算 hashKey 花 O(m)

一共 n 個 所以整個時間複雜是 O(m*n)

![](https://i.imgur.com/Lv3gu52.png)

## 程式碼
```go
package sol

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}
	hash := make(map[([26]int)]([]string))
	result := [][]string{}
	var countCharFreq = func(input string) [26]int {
		var charFreq [26]int
		sLen := len(input)
		for pos := 0; pos < sLen; pos++ {
			charFreq[input[pos]-'a']++
		}
		return charFreq
	}
	for _, val := range strs {
		key := countCharFreq(val)
		hash[key] = append(hash[key], val)
	}
	for _, val := range hash {
		result = append(result, val)
	}
	return result
}

```
## 困難點

1. 需要想出 anagram 的特性來做hashKey 分類

## Solve Point

- [x]  建立一個 hashTable hash = map([26]int)([]string)， result = [][]string{}
- [x]  從 pos=0..len(strs)-1 做以下運算
- [x]  計算  s[pos] 的 charFreq 並且把 charFreq 當作 key 更新 hash[key] = append(hash[key], s[pos])
- [x]  把 map 內的所有value 依序 append 到 result 內
- [x]  回傳 result