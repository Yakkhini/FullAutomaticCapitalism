# Full-Automatic Capitalism
全自动资本主义机器

## 策略的基本思路

本策略采用网格理论来做量化。即采用市场的震荡，在较短的震荡范围内做空做多。这样能盈利的原理是对于市场震荡来说，其在震荡范围内的价格来回浮动，在做多、做空后，很容易达到小幅度地买入卖出的平仓触发价格。

### 网格划定

对于网格，需要确定的有：均线，上线，下线。均线需要计算近一段时间的币市均价。对于上线、下线，统计币价分布频率，选取边界。

### 建仓

对于建仓来说，在尽可能的临近上线处建空仓，视为防御空仓。在临近下线处建多仓，视为防御多仓。建仓触发价格选取，在满足临近边线的同时也要尽量提高频率，从而满足建仓效率。目前初步选定的频率参数为 0.3。

在防御仓建成后，即为建仓成功。对于防御仓位，不轻易卖出。

#### 防御仓位调整

对于防御仓位在建仓成功后不会卖出，但会作调整。当币价位于防御空仓之上，同时更接近上线时，在此币价建立新的防御多仓，原防御多仓变为交易多仓。

### 交易仓位

交易仓位是主要买卖的仓位。

#### 交易开仓

在建仓成功后，就可以做交易仓的开仓工作。交易多仓可以在均线以下、防御多仓以上开仓，交易空仓可以在均线以上、防御多仓以下开仓。

#### 交易平仓

在交易仓盈利达到止盈点后，将交易仓平仓。止盈参数初步设置为 1%。

### 风险控制

风险控制就是要考虑策略亏损的情况，并通过仓位组合来降低风险。

#### 亏损来源

亏损来自高买低卖的多仓和低买高卖的空仓。

1. 当标记价格低于交易多仓，高于防御多仓
2. 当标记价格低于防御多仓，高于下线
3. 当标记价格低于下线
4. 当标记价格高于交易空仓，低于防御空仓
5. 当标记价格高于防御空仓，低于上线
6. 当标记价格高于上线

对于 1，4 种情况，继续持仓，等待盈利触发价格平仓。

对于 2，5 种情况，新增防御仓位，原防御仓位转为交易仓位，变为 1，4 种情况。

对于 3，6 种情况，在低于下线时交易多仓止损，留防御多仓，所有空仓合并为防御空仓；在高于上线时交易空仓止损，留防御空仓，所有多仓合并为防御多仓。
