<!-- # Go - Overview
Go là một ngôn ngữ đa mục đích, được dùng để thiết kế hệ thống. Nó được phát triển bởi Google vào 2007. Nó mạnh mẽ, biên dịch tĩnh và còn hỗ trợ lập trình đồng thời.
Các chương trình được xây dựng bằng cách sử dụng các packages, để quản lý hiệu quả các dependencies. Việc thực thi Go sử dụng một trình biên dịch truyền thống và liên kiết các model để tạo ra các đoạn mã
thực thi. Ngôn ngữ Go được giới thiệu vào tháng 11 năm 2009 và được sử dụng trong một số hệ thống của Google.

Features of Go Programming
=======
Những tính năng quan trọng nhất của Go được liệt kê bên dưới:
- Hỗ trợ -->

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
 Các biến trong **Go** luôn phải được định nghĩa một cách rõ ràng. **Go** là một 
