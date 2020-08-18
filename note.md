## 转义字符

| verb      | 描述 |
| :---      | :---  |
| %d        | 十进制整数 |
| %x %0 %b  | 十六进制 八进制 二进制 |
| %f %g %e  | 浮点数 3.141593 3.141592653589793 3.141593e+00 |
| %t        | 布尔型 true false |
| %c        | 字符 (Unicode 码点) |
| %s        | 字符串 |
| %q        | 带引号字符串 ("abc") ('c') |
| %V        | 内置格式的任何值 |
| %T        | 任何值的类型 |
| %%        | 百分号本身 |
| \a        | 警告或响铃 |
| \b        | 退格符 |
| \f        | 换页符 |
| \n        | 换行符 (指直接跳到下一行的同一位置) |
| \r        | 回车符 (指返回到行首) |
| \t        | 制表符 |
| \v        | 垂直制表符 |

## 关键字 (25个)
|               |               |               |               |               |
| :-----------: | :-----------: | :-----------: | :-----------: | :-----------: |
| break         | default       | func          | interface     |  select       |
| case          | defer         | go            | map           |  struct       |
| chan          | else          | goto          | package       |  switch       |
| const         | fallthrough   | if            | range         |  type         |
| continue      | for           | import        | return        |  var          |


## 内置预声明的常量 类型 函数

- 常量:true  false  iota  nil
- 类型: (u)int  (u)int8  (u)int16  (u)int32  (u)int64 uintptr  float32  float64  complex128  complex64  bool  byte  rune  string  error
- 函数: make  len  cap  new  append  copy  close  delete  complex  real imag  panic  recover