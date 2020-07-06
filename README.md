![测试结果](https://github.com/aaronzjc/leetcode/workflows/tests/badge.svg)
[![codecov](https://codecov.io/gh/aaronzjc/leetcode/branch/master/graph/badge.svg?token=U97K2S543N)](https://codecov.io/gh/aaronzjc/leetcode)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/1e452326595b49c2a3f1ff769fae33e8)](https://www.codacy.com/manual/415397228/leetcode?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=aaronzjc/leetcode&amp;utm_campaign=Badge_Grade)

## 介绍

😫啥也别说了，刷吧。。。

🙂自己实现的，如果遇到写不出来的，会参考其他人的思路。每道题不一定是最优解。

😏代码测试覆盖率100%。因为测试用例的关系，不保证100%正确，如果你发现解法有问题的地方，欢迎告诉我！

😎为了方便测试，以及针对重复性的内容，编写了一些辅助方法。例如，根据数组构造链表，比较两个数组是不是相等，比较两个二维数组是不是相等，等。目前用起来还不错。项目还包括，一些自己觉得有意思的或者基础性的算法，例如折半查找，求最大公因数等。

👏希望这个项目成为自己学习算法的沉淀。

## 测试

```shell script
make test
make test_profile // 测试覆盖率
make test_profile_html // 测试覆盖率html查看
```

## 完成问题列表

### 简单

+ 两个数的和
+ 整数反转
+ 数是否是回文数
+ 罗马数转整数
+ 最长公共前缀
+ 括号是否匹配
+ 合并两个链表
+ 删除排序数组重复项
+ 移除数组元素
+ strstr
+ 寻找有序数组插入位置
+ 外观数列

### 中等

+ 链表两个数相加
+ 容器能盛的最多水
+ 最长回文子串
+ 最长无重复子串
+ 字符串转整数
+ Z字形打印字符串
+ 整数转罗马数
+ 3个数之和等于0
+ 3个数之和最接近给定数
+ 电话号码的字母组合
+ 删除单链表的倒数第N个节点
+ 括号生成
+ 两两交换链表中的节点
+ 两数相除
+ 下一个较大排列
+ 搜索旋转排序数组
+ 排序数组查找开始和结束位置
+ 有效的数独
+ 组合总和
+ 组合总和2
+ 字符串相乘
+ 全排列
+ 全排列2
+ 旋转图像
+ 字母异位词分组
+ power(x, n)

### 难 

+ 两个整形数组中位数
+ 合并K个有序链表
+ K个一组翻转链表
+ 串联所有单词的子串
+ 最大有效括号长度
+ 解数独
+ 缺失的第一个正数
+ 接雨水
+ 通配符匹配
+ 跳跃游戏2
+ N皇后
+ N皇后2

### 其他

+ 求最大公约数
+ 折半查找大于给定数

## 工具包

+ 函数
    + 两个一维数组是否相等(Int和String，有序和无序)
    + 两个二维数组是否相等(Int和Byte和String, 有序和无序)
    + 获取字符串数组的全排列
+ 链表
    + 根据数组构造链表
    + 根据链表输出数组
    + 反转链表
    + 输出链表元素个数和首尾节点
+ 树
    + 根据数组构造一棵树
    + DFS，DFS递归版
    + BFS
