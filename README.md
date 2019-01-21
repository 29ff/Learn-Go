# Go overview
Để bắt đầu với **Go** (hay còn gọi là **Golang**), hãy cũng tìm hiểu qua một chút về ngôn ngữ này. **Go** được phát triển bởi các kỹ sư của Google. Nó được phát triển từ năm 2007 và được công bố năm 2009. Phiên bản mã nguồn mở đầu tiên của **Go** được ra mắt vào tháng 3 năm 2012. Theo như những nhà phát triển của **Go** thì nó là một ngôn ngữ lập trình mã nguồn mở giúp dễ dàng xây dựng những phần mềm đơn giản, đáng tin cậy và hiệu quả.

Trong nhiều ngôn ngữ, có rất nhiều cách để giải quyết một vấn đề cho trước.  Những lập trình viên có thể phải dành ra rất nhiều thời gian để tìm kiếm cách giải quyết tốt nhất để giải quyết một vấn đề

Với **Go**, những nhà phát triển tin rằng với ít tính năng hơn, họ cũng sẽ chỉ cần nghĩ tới một cách giải quyết duy nhất cho vấn đề cần giải quyết. Điều này tiết kiệm được rất nhiều thời gian của lập trình viên và khiến cho việc tối ưu kể cả những **codebase** lớn cũng trở lên dễ dàng. Cũng sẽ không có những tính năng khó hiểu như **maps** hoặc **filters** trong **Go**

> When you have features that add expensiveness it typically adds expense - Rob Pike

# Getting Started
**Go** được tạo ra bởi các **package**. Package **main** cho trình biên dịch Go biết rằng chương trình được biên dịch dưới dạng thực thi, thay vì dùng thư viện. Đây là điểm bắt đầu biên dịch cho một ứng dụng **Go**. Package **main** được định nghĩa như sau:
```go
package main
```

Tiếp theo hay cùng tạo một ứng dụng "Hello world!" đơn giản với **Go** bằng việc tạo ra file *main.go*.

## Workspace
Một **workspace** trong **Go** được định nghĩa bởi biến môi trường **GOPATH**.

**Go** sẽ tìm kiếm mọi package bên trong thư mục được định nghĩa trong **GOPATH**, hoặc được định nghĩa trong **GOROOT**, được đặt mặc định khi cài đặt **Go**. **GOROOT** là đường dẫn đến vị trí **Go** được cài đặt.

Cài đặt **GOPATH** đến thư mục mà bạn mong muốn. Hãy thêm nó vào trong một folder *~/workspace*

```
# export env
export GOPATH=~/workspace

# go inside the workspace directory
cd ~/workspace
```

Tạo một file *main.go* trong folder mà chúng ta vừa tạo với nội dung bên dưới

## Hello World!
```golang
package main

import "fmt"

func main() {
  fmt.Println("Hello World!")
}
```

Trong ví dụ trên, **fmt** là một package tích hợp bên trong **Go**. Nó sẽ thực hiện chức năng định dạng I/O (Input/Output)

Chúng ta import một package trong **Go** bằng từ khóa **import**. **func main** là điểm chính mà mã được thực thi. **Println** là một function bên trong package **fmt** và nó sẽ thực hiện in ra "Hello World!" cho chúng ta

Hãy cùng xem kết quả sau khi chạy file này. Có hai cách để chạy một file trong **Go**. Như chúng ta đã biết, **Go** là một ngôn ngữ biên dịch, vì vậy trước khi chạy nó, chúng ta cần biên dịch trước:

```
go build main.go
```

Câu lệnh này sẽ tạo ra một file thực thi nhị phân để chúng ta có thể chạy:

```
go run main.go
# Hello World!
```

 # Variables
 Các biến trong **Go** luôn phải được định nghĩa một cách rõ ràng. Kiểu dữ liệu trong **Go** sẽ được kiểm tra tại thời điểm biến được khai báo. Một biến trong **Go** có thể được khai báo như sau:

```golang
var a int
```

Trong trường hợp này, giá trị của biến được đặt bằng 0. Sử dụng cú pháp dưới đây để định nghĩa và khởi tạo một biến với những giá trị bất kỳ:

```golang
var a = 1
```

Chúng ta cũng có thể định nghĩa nhiều biến trong cùng một dòng

```golang
var b, c int = 2, 3
```

# Data Types
Giống như những ngôn ngữ khác, **Go** hỗ trợ nhiều kiểu cấu trúc dữ liệu khác nhau. Hãy cùng thử khám phá chúng:

## Number, String, and Boolean

Một số loại kiểu dữ liệu dạng số là int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, ...

Kiểu chuỗi trong **Go** được lưu thành một chuỗi các bytes. Nó được khai báo với từ khóa **string**

Một giá trị boolean được lưu trữ và khởi tạo bằng từ khóa **bool**

**Go** cũng hỗ trợ những dữ liệu kiểu số phức tạp, chúng có thể được khởi tạo với từ khóa **complex64** và **complex128**

```golang
var a bool = true
var b int = 1
var c string = 'hello world'
var d float 32 = 1.222
var x complex128 = cmplx.Sqrt(-5 * 12i)
```

## Array, Slices, and Maps

Một mảng là một tập hợp trình tự các phần tử có cùng kiểu dữ liệu. Mảng có một chiều dài cố định được định nghĩa lúc khởi tạo, vì thế, nó không thể mở rộng hơn chiều dài đó. Một mảng có thể được định nghĩa như sau:

```golang
var a [5]int
```

Chúng ta cũng có thể khai báo mảng đa chiều như sau:

```golang
var multiD [2][3]int
```

Mảng trong **Go** không cung cấp khả năng để lấy ra một mảng con. Để làm điều đó, **Go** có một kiểu dữ liệu gọi là **slice**

Slice lưu trữ một chuỗi các phần tử có thể được mở rộng bất cứ lúc nào. Khai báo slice cũng giống như khai báo array trong **Go**, chỉ khác là khi khai báo slice chúng ta không cần phải khai báo độ dài của slice hay số phần tử, ví dụ:

```golang
var b []int
```

Dòng khai báo trên sẽ tạo ra một slice với độ dài bằng 0 và số phần tử bằng 0. Tuy nhiên chúng ta vẫn có thể khai báo slice với độ dài hay số phần tử nhất định bằng câu khai báo dưới đây:

```golang
numbers := make([]int, 5, 10)
```

Đoạn code trên sẽ khởi tạo một slice có độ dài là 5 và sức chứa là 10

Chúng ta có thể nghĩ slice như là một lớp của array. Slice sử dụng array như là một cấu trúc cơ bản. Một slice thì sẽ có 3 thành phần: "sức chứa, độ dài, và một con trỏ tới mảng" như hình bên dưới:

![Slice Array](/images/slice_array.png)

Chúng ta có thể tăng sức chứa của slice bằng cách sử dụng chức năng *append* hoặc *copy*. *Append* sẽ thêm giá trị vào cuối slice và sẽ tăng sức chứa của slice nếu cần

```golang
numbers = append(numbers, 1, 2, 3, 4)
```

Một cách khác nữa để gia tăng sức chứa của slice là sử dụng chức năng *copy*. Sử dụng *copy* sẽ tạo ra một slice mới với sức chứa lớn hơn slice cũ và sao chép toàn bộ phần tử của slice gốc

```golang
// create a new slice
number2 := make([]int, 15)
// copy the original slice to new slice
copy(number2, number)
```

Chúng ta cũng có thể tạo một slice con của một slice. Ví dụ:

```golang
// initialize a slice with 4 len and values
var number2 = []int{1, 2, 3, 4}
fmt.Println(number2) // => [1, 2, 3, 4]

// create sub slices
slice1 := number2[2:]
fmt.Println(slice1) // => [3, 4]

slice2 := number2[:3]
fmt.Println(slice2) // => [1, 2, 3]

slice3 := number2[1:4]
fmt.Println(slice3) // => [2, 3, 4]

slice4 := number2[2:4]
fmt.Println(slice4) // => [3, 4]

slice5 := number2[number2[0]:number2[len(number2)-1]]
fmt.Println(slice5) // => [2, 3, 4]     
```

Map là một kiểu dữ liệu trong **Go**, nó bao gồm các key và value được map với nhau. Chúng ta có thể khai báo map trong **Go** bằng cách sử dụng cú pháp sau:

```golang
var m map[string]int
```

Câu lệnh trên khai báo một biến m là một Map, có các key là kiểu String và value là kiểu Integer. Chúng ta có thể thêm các cặp key - value cho Map bằng câu lệnh sau:

```golang
// adding key/value
m ["clearity"] = 2
m ["simplicity"] = 3

// printing the values
fmt.Println(m["clearity"]) // => 2
fmt.Println(m["simplicity"]) // => 3
```

# Typecasting

Một kiểu dữ liệu có thể được convert sang một kiểu khác bằng cách sử dụng **type casting**. Hãy xem một ví dụ bên dưới:

```golang
a := 1.1
b := int(a)
fmt.Println(b)
// => 1
```

Không phải mọi kiểu dữ liệu đều có thể convert được sang kiểu khác. Đảm bảo rằng kiểu dữ có thể chuyển đổi được trước khi thực hiện chuyển đổi

# Conditional Statements

## if else

Đối với câu lệnh điều kiện, chúng ta có thể sử dụng câu lệnh if - else như ví dụ bên dưới:

```golang
if num := 9; num < 0 {
  fmt.Println(num, "is negative")
} else if num < 10 {
  fmt.Println(num, "has 1 digit")
} else {
  fmt.Println(num, "has multiple digits")
}
```

## switch case
Dưới đây là cách sử dụng câu lệnh switch case:

```golang
i := 2
switch i {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  default:
    fmt.Println("none")
}
```

# Looping

Go có từ khóa **for** được dùng cho vòng lặp:

```golang
i := 0
sum := 0
for i < 10 {
  sum += 1
  i++
}

fmt.Println(sum)
```

Vòng lặp vô hạn trong Go:

```golang
for {

}
```

# Pointer

Khi chúng ta gọi một hàm và truyền tham số vào đó, giá trị của tham số đó sẽ được sao chép vào trong hàm đó, ví dụ:

```golang
package main

import "fmt"

func zero(x int) {
  x = 0
}

func main() {
  x := 5
  zero(x)
  fmt.Println(x) // => 5
}
```

Trong đoạn code trên, hàm *zero()* có tác dụng gán cho biến x giá trị là 0. Trong hàm main() chúng ta khai báo một biến x có giá trị là 5 rồi truyền vào hàm *zero()*, sau đó chúng ta in giá trị của biến x ra màn hình, kết quả vẫn bằng 5. Lý do là vì giá trị của biến x trong hàm *main()* được sao chép vào tham số x của riêng hàm *zero()* chứ hàm *zero()* không nhận một biến x nào cả

Tuy nhiên nếu chúng ta muốn hàm *zero()* thao tác trực tiếp luôn với biến x của hàm *main()* thì chúng ta phải dùng đến con trỏ. Ví dụ:

```golang
package main

import "fmt"

func zero(xPrt *int) {
  *xPtr = 0
}

func main() {
  x := 5
  zero(&x)
  fmt.Println(x) // => 0
}
```

Con trỏ trong Go là một loại biến đặc biệt, được dùng để lưu trữ địa chỉ của biến khác trong bộ nhớ RAM chứ không lưu trữ một giá trị cụ thể nào. Để khai báo một biến con trỏ thì chúng ta thêm dấu **\*** vào trước tên kiểu dữ liệu. Khi chúng ta in giá trị của một biến con trỏ ra màn hình thì giá trị đó sẽ là một số ở hệ hexa (hệ 16), đó là dịa chỉ bộ nhớ mà con trỏ này đang trỏ tới

Khi gán giá trị cho biến con trỏ, chúng ta cũng phải đưa vào đó một địa chỉ bộ nhớ nào đó chứ không đưa một giá trị vào, để lấy địa chỉ bộ nhớ của một biến thì chúng ta dùng dấu **&** trước tên biến

Ngoài chức năng khai báo biến con trỏ, dấu **\*** còn có tác dụng lấy giá trị của địa chỉ bộ nhớ mà con trỏ đang tham chiếu tới, ngược lại chúng ta cũng có thể gán giá trị cho địa chỉ đó thông qua dấu **\***. Ví dụ:

```golang
package main

import "fmt"

func main() {
  var x *int
  var y int

  y = 0
  x = &y

  fmt.Println(x)
  fmt.Println(&x)

  *x = 1

  fmt.Println(*x)
  fmt.Println(y)
}
```

Trong đoạn code trên, chúng ta:
 - Khai báo x là biến con trỏ kiểu int, y là một biến int bình thường
 - Gán giá trị cho y là 0
 - Cho x trỏ tới địa chỉ bộ nhớ của y
 - Lúc này x sẽ mang giá trị là địa chỉ bộ nhớ của y, chúng ta có thể dùng dấu & để lấy địa chỉ của bộ nhớ y, hoặc dùng dấu **\*** để lấy giá trị tại địa chỉ của y
 - Gán giá trị cho y (hay giá trị tại địa chỉ bộ nhớ mà x đang tham chiếu tới) là 1 bằng cách dùng dấu **\***

# Functions

Function **main** định nghĩa trong package **main** là nơi bắt đầu của một chương trình Go.
Ngoài ra chúng ta cũng có thể định nghĩ được những Function khác:

```golang
func add(a int, b int) int {
  c := a + b
  return c
}

func main() {
  fmt.Println(add(2, 1)) // => 3
}
```
Như chúng ta có thể thấy ở ví dụ trên, một function trong Go được định nghĩa bằng từ khóa **func** và tiếp theo là tên của function. Những tham số của một function nhận vào được đi kèm với kiểu của tham số đó, và cuối cùng là kiểu của giá trị trả về. Một biến được trả về trong function cũng có thể được xác định trước như sau:

```golang
func add(a int, b int) (c int) {
  c = a + b
  return
}

func main() {
  fmt.Println(add(2, 1)) // => 3
}
```

Biến c được định nghĩa như là một biến trả về. Vì vậy biến c sẽ tự động trả về mà không cần phải ghi rõ trong câu lệnh return

Chúng ta cũng có thể trả về nhiều giá trị từ một function bằng cách chia cách các giá trị bằng dấu "**,**"

```golang
func add(a int, b int) (int, string) {
  c := a + b
  return c, "successfully added"
}

func main() {
  sum, message := add(2, 1)
  fmt.Println(message)
  fmt.Println(sum)
}
```

# Method, Structs and Interfaces

Go không phải là một ngôn ngữ hướng đối tượng, nhưng với **Structs**, **Interfaces** và **Method** thì nó khiến chúng ta cảm thấy như đang sử dụng một ngôn ngữ hướng đối tượng

## Struct

Một Struct là một kiểu giá trị, một tập hợp các trường khác nhau. Một Struct thường được dùng để nhóm các giá trị lại với nhau. Ví dụ, nếu chúng ta một nhóm các giá trị của một kiểu "Person", chúng ta định nghĩa ra những thuộc tính của "Person" như "name", "age", "gender". Một Struct có thể được định nghĩa như sau:

```golang
type Person struct {
  name string
  age int
  gender string
}
```

Với một struct "Person" đã được định nghĩa như trên, chúng ta có thể khởi tạo một "Person" như sau:

```golang
// cách 1: khởi tạo một Person bằng các khai báo cả các trường và giá trị của các trường đó
p = person{name: "Bob", age: 42, gender: "Male"}
// cách 2: chỉ khai báo các giá trị
person("Bob", 42, "Male")
```

Chúng ta có thể dễ dàng lấy giá trị của các trường dựa vào dấu "**.**"

```golang
p.name // => Bob
p.age // => 42
p.gender // => Male
```

Chúng ta cũng có thể lấy giá trị của các thuộc tính trong Struct bằng Pointer như sau:

```golang
pp = &person{name: "Bob", age: 42, gender: "Male"}
pp.name // => Bob
```

## Methods

Method là một kiểu dữ liệu đặc biệt của function với một **receiver**. Một receiver có thể là một giá trị hoặc cũng có thể là một Pointer. Chúng ta hãy cùng thử tạo một Method với tên là "describe" và có một receiver có type là Person:

```golang
package main
import "fmt"

// khởi tạo struct
type person struct {
  name string
  age int
  gender string
}

// khởi tạo method
func (p *person) describe() {
  fmt.Printf("%v is %v years old", p.name, p.age)
}

func (p *person) setAge(age int) {
  p.age = age
}

func (p person) setName(name string) {
  p.name = name
}

func main() {
  pp := &person{name: "Bob", age: 42, gender: "Male"}
  pp.describe()
  // => Bob is 42 years old
  pp.setAge(45)
  fmt.Println(pp.age)
  // => 45
  pp.getName("Hari")
  fmt.Println(pp.name)
  // => Bob
}
```

Như chúng ta có thể thấy ở ví dụ trên, method có thể được gọi đến qua dấu "**.**" như "pp.describe". Chúng ta cần chú ý rằng receiver là một Pointer. Với receiver là một Pointer, chúng ta sẽ chuyển một tham chiếu đến giá trị, vì vậy nếu chúng ta thực hiện bất cứ thay đổi nào, nó sẽ thay đổi trực tiếp đến Pointer. Nó cũng sẽ không tạo thêm bất cứ một bản sao mới nào của đối tượng, giúp tiết kiệm bộ nhớ

Chúng ta có thể để ý rằng trong ví dụ trên, giá trị của "age" đã được thay đổi trong khi giá trị của "name" không thay đổi là do phương thức "setName" là một **receiver type** trong khi đó "setAge" là một **pointer type**

## Interfaces

Interface trong Go là một tập hợp các Method. Interfaces giúp chúng ta nhóm những phương thức của một type lại với nhau. Hãy thử tạo một Interfaces có tên là "animal" như sau:

```golang
type animal interface {
  description() string
}
```

Trong đoạn code trên, "animal" là một kiểu Interface. Bây giờ chúng ta hãy thử tạo 2 đối tượng của "animal" mà có thể thực thi được interface "animal":

```golang
package main
import "fmt"

type animal interface {
  description() string
}

type cat struct {
  Type string
  Sound string
}

type snake struct {
  Type string
  Poisonous bool
}

func (s snake) description() string {
  return fmt.Printf("Poisonous: %v", s.Poisonous)
}

func (c cat) description() string {
  return fmt.Printf("Sound: %v", c.Sound)
}

func main() {
  var a animal
  a = snake{Poisonous: true}
  fmt.Println(a.description()) // => Poisonous: true
  a = cat{Sound: "Meow!!"}
  fmt.Println(a.description()) // => Sound: Meow!!
}
```

Trong function main, chúng ta tạo một biến a có thuộc kiểu animal. Chúng ta khởi tạo 2 struct là "cat" và "snake" thuộc type animal và sử dụng Println để in ra màn hình a.description. Chúng ta thực hiện phương thức description trên "cat" và "snake" khác nhau, vì vậy chúng ta có 2 cách mô tả khác nhau cho 2 con vật
