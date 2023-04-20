<!-- README.md -->
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/fishjerky/Go-Bowling/go.yml) 
[![cov](https://fishjerky.github.io/Go-Bowling/badges/coverage.svg)](https://github.com/fishjerky/Go-Bowling/actions)
# go-bowling" 

# Description

- 提供Game Struct
- 具兩個func
 - void 滾球(int 瓶數)
 - int 總分()

# 規則
 - 共10局
 - Bonus
  - Spare Bonus
    - 加下一局的第一次出手瓶數
  - Strike Bonus
    - 加下兩個出手瓶數
  - 注意事項
    - 第10局有三次機會

# Test Case
 - testNewGame()
   - 初始化Game
   - expected: 0
- test擊倒1瓶()
   - 測試擊倒1瓶
   - rest: 0
   - expected: 1
 - test測試全擊倒1瓶
   - 全都擊倒1瓶
   - expected: 20 
 - test 1個Spare
   - 1st: 2,8 (Spare)
   - 2nd: 1,2
   - rest:0
   - expected: 14
 - test SpareBonus
   - case 1
    - 1st: 1,9 (Spare)
    - 2nd: 1,2
    - rest:0
    - expected: 14
   - case 2
    - 1st: 1,1 
    - 2st: 1,9 (Spare)
    - 3nd: 1,2
    - rest:0
    - expected: 16
 - test 1個Strike
   - 1st: X (Strike)
   - 2nd: 1,2
   - rest: 0
   - expected: 16
 - test Perfect Game
   - 全都strike
   - 第十局有3次strike，共12個strike
   - expected 300
